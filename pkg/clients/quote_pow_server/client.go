package quote_pow_server

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/yusupovanton/words-of-wisdom-POW/pkg/clog"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/metrics"
)

type Client struct {
	address  string
	logger   clog.CLog
	registry metrics.Registry
	series   metrics.Series
}

func NewClient(host, port string, logger clog.CLog, registry metrics.Registry) *Client {
	address := fmt.Sprintf("%s:%s", host, port)
	return &Client{
		address:  address,
		logger:   logger,
		registry: registry,
		series:   metrics.NewSeries(metrics.SeriesTypeClient, "get_quote"),
	}
}

func (c *Client) GetQuote(ctx context.Context) (string, error) {
	ctx, series := c.series.WithOperation(ctx, "fetch_quote")

	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		c.registry.Inc(series.Error("connect_failure"))
		c.logger.ErrorCtx(ctx, err, "Failed to connect to server")
		return "", fmt.Errorf("failed to connect to server: %w", err)
	}
	defer func() {
		if err = conn.Close(); err != nil {
			c.logger.ErrorCtx(ctx, err, "could not close connection")
		}
	}()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	challengeLine, err := reader.ReadString('\n')
	if err != nil {
		c.registry.Inc(series.Error("read_challenge_failure"))
		c.logger.ErrorCtx(ctx, err, "Failed to receive challenge line")
		return "", fmt.Errorf("failed to receive challenge line: %w", err)
	}
	challengeLine = strings.TrimSpace(challengeLine)
	c.logger.DebugCtx(ctx, fmt.Sprintf("Received challenge line: %s", challengeLine))

	difficultyLine, err := reader.ReadString('\n')
	if err != nil {
		c.registry.Inc(series.Error("read_difficulty_failure"))
		c.logger.ErrorCtx(ctx, err, "Failed to receive difficulty line")
		return "", fmt.Errorf("failed to receive difficulty line: %w", err)
	}
	difficultyLine = strings.TrimSpace(difficultyLine)
	c.logger.DebugCtx(ctx, fmt.Sprintf("Received difficulty line: %s", difficultyLine))

	if !strings.HasPrefix(challengeLine, "Challenge: ") || !strings.HasPrefix(difficultyLine, "Difficulty: ") {
		c.registry.Inc(series.Error("invalid_challenge_format"))
		return "", fmt.Errorf("invalid challenge format")
	}

	prefix := strings.TrimSpace(strings.TrimPrefix(challengeLine, "Challenge: "))
	difficultyStr := strings.TrimSpace(strings.TrimPrefix(difficultyLine, "Difficulty: "))
	difficulty, err := strconv.Atoi(difficultyStr)
	if err != nil {
		c.registry.Inc(series.Error("invalid_difficulty"))
		c.logger.ErrorCtx(ctx, err, "Invalid difficulty value")
		return "", fmt.Errorf("invalid difficulty: %w", err)
	}

	nonce := c.solveChallenge(ctx, prefix, difficulty)
	c.logger.DebugCtx(ctx, fmt.Sprintf("Solved challenge: nonce=%s", nonce))

	_, err = fmt.Fprintf(writer, "%s\n", nonce)
	if err != nil {
		c.registry.Inc(series.Error("send_nonce_failure"))
		c.logger.ErrorCtx(ctx, err, "Failed to send nonce")
		return "", fmt.Errorf("failed to send nonce: %w", err)
	}
	err = writer.Flush()
	if err != nil {
		c.registry.Inc(series.Error("flush_writer_failure"))
		c.logger.ErrorCtx(ctx, err, "failed to flush writer")
		return "", fmt.Errorf("failed to flush writer: %w", err)
	}
	c.logger.DebugCtx(ctx, fmt.Sprintf("Submitted nonce: %s", nonce))

	quote, err := reader.ReadString('\n')
	if err != nil {
		c.registry.Inc(series.Error("read_quote_failure"))
		c.logger.ErrorCtx(ctx, err, "failed to receive quote")
		return "", fmt.Errorf("failed to receive quote: %w", err)
	}
	quote = strings.TrimSpace(quote)
	c.logger.DebugCtx(ctx, fmt.Sprintf("received quote: %s", quote))

	c.registry.Inc(series.Success())
	return quote, nil
}

func (c *Client) solveChallenge(ctx context.Context, prefix string, difficulty int) string {
	requiredPrefix := strings.Repeat("0", difficulty)
	nonce := 0
	for {
		select {
		case <-ctx.Done():
			return ""
		default:
			data := fmt.Sprintf("%s%d", prefix, nonce)
			hash := sha256.Sum256([]byte(data))
			hashHex := hex.EncodeToString(hash[:])
			if strings.HasPrefix(hashHex, requiredPrefix) {
				return fmt.Sprintf("%d", nonce)
			}
			nonce++
		}
	}
}
