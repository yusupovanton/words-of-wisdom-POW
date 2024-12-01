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
)

type Client struct {
	address string
	logger  clog.CLog
}

func NewClient(port string, logger clog.CLog) *Client {
	address := fmt.Sprintf("localhost:%s", port)
	return &Client{
		address: address,
		logger:  logger,
	}
}

func (c *Client) GetQuote(ctx context.Context) (string, error) {
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return "", fmt.Errorf("failed to connect to server: %w", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	challengeLine, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to receive challenge line: %w", err)
	}
	challengeLine = strings.TrimSpace(challengeLine)
	c.logger.InfoCtx(ctx, "Received challenge line", "line", challengeLine)

	difficultyLine, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to receive difficulty line: %w", err)
	}
	difficultyLine = strings.TrimSpace(difficultyLine)
	c.logger.InfoCtx(ctx, "Received difficulty line", "line", difficultyLine)

	if !strings.HasPrefix(challengeLine, "Challenge: ") || !strings.HasPrefix(difficultyLine, "Difficulty: ") {
		return "", fmt.Errorf("invalid challenge format")
	}

	prefix := strings.TrimSpace(strings.TrimPrefix(challengeLine, "Challenge: "))
	difficultyStr := strings.TrimSpace(strings.TrimPrefix(difficultyLine, "Difficulty: "))
	difficulty, err := strconv.Atoi(difficultyStr)
	if err != nil {
		return "", fmt.Errorf("invalid difficulty: %w", err)
	}

	nonce := c.solveChallenge(prefix, difficulty)
	c.logger.InfoCtx(ctx, "Solved challenge", "nonce", nonce)

	_, err = fmt.Fprintf(writer, "%s\n", nonce)
	if err != nil {
		return "", fmt.Errorf("failed to send nonce: %w", err)
	}
	err = writer.Flush()
	if err != nil {
		return "", fmt.Errorf("failed to flush writer: %w", err)
	}
	c.logger.InfoCtx(ctx, "Submitted nonce", "nonce", nonce)

	quote, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to receive quote: %w", err)
	}
	quote = strings.TrimSpace(quote)
	c.logger.InfoCtx(ctx, "Received quote", "quote", quote)

	return quote, nil
}

func (c *Client) solveChallenge(prefix string, difficulty int) string {
	requiredPrefix := strings.Repeat("0", difficulty)
	nonce := 0
	for {
		select {
		case <-context.Background().Done():
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
