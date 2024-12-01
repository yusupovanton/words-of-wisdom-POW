package handlers

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/yusupovanton/words-of-wisdom-POW/pkg/clog"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/metrics"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/pow"
)

type quoteGetter interface {
	GetRandomQuote() (string, error)
}

type GetQuoteHandler struct {
	logger   clog.CLog
	registry metrics.Registry
	series   metrics.Series

	quoteGetter quoteGetter

	Difficulty int
	Prefix     string
}

func NewGetQuoteHandler(logger clog.CLog, registry metrics.Registry, quoteGetter quoteGetter, Difficulty int, Prefix string) *GetQuoteHandler {
	return &GetQuoteHandler{
		logger:      logger,
		registry:    registry,
		series:      metrics.NewSeries(metrics.SeriesTypeApiHandler, "get_quote"),
		quoteGetter: quoteGetter,
		Difficulty:  Difficulty,
		Prefix:      Prefix,
	}
}

func (h *GetQuoteHandler) GetQuote(conn net.Conn) {
	ctx, series := h.series.WithOperation(context.TODO(), "get_quote")
	defer func() {
		if err := conn.Close(); err != nil {
			h.logger.ErrorCtx(ctx, err, "could not close connection")
		}
	}()

	challenge, err := pow.GenerateChallenge(h.Prefix, h.Difficulty)
	if err != nil {
		h.logger.ErrorCtx(ctx, err, "could not generate challenge")
		h.registry.Inc(series.Error("generate_challenge"))
		h.writeError(ctx, conn, "could not generate challenge")

		return
	}

	h.logger.InfoCtx(ctx, "Challenge generated", "prefix", challenge.Prefix, "difficulty", challenge.Difficulty)

	if err = h.writeMessage(conn, fmt.Sprintf("Challenge: %s\nDifficulty: %d\n", challenge.Prefix, challenge.Difficulty)); err != nil {
		h.logger.ErrorCtx(ctx, err, "could not send challenge to client")
		h.registry.Inc(series.Error("send_challenge"))

		return
	}

	reader := bufio.NewReader(conn)
	nonce, err := reader.ReadString('\n')
	if err != nil {
		h.logger.ErrorCtx(ctx, err, "error reading nonce")
		h.registry.Inc(series.Error("read_nonce"))
		h.writeError(ctx, conn, "error reading nonce")

		return
	}
	nonce = strings.TrimSpace(nonce)

	h.logger.InfoCtx(ctx, "Received nonce", "nonce", nonce)

	valid, err := challenge.CheckSolution(nonce)
	if err != nil {
		h.logger.ErrorCtx(ctx, err, "error verifying solution")
		h.registry.Inc(series.Error("verify_solution"))

		h.writeError(ctx, conn, "error verifying solution")
		return
	}

	if !valid {
		h.logger.WarnCtx(ctx, "invalid solution", "nonce", nonce)
		h.registry.Inc(series.Info("invalid_solution"))

		err = h.writeMessage(conn, "invalid solution.\n")
		if err != nil {
			h.logger.ErrorCtx(ctx, err, "error writing message to connection")
		}
		return
	}

	h.logger.InfoCtx(ctx, "Solution verified", "nonce", nonce)

	quote, err := h.quoteGetter.GetRandomQuote()
	if err != nil {
		h.logger.ErrorCtx(ctx, err, "error fetching quote")
		h.registry.Inc(series.Error("fetch_quote"))
		h.writeError(ctx, conn, "error fetching quote")

		return
	}

	h.logger.InfoCtx(ctx, "Quote fetched", "quote", quote)

	if err = h.writeMessage(conn, fmt.Sprintf("Quote: %s\n", quote)); err != nil {
		h.logger.ErrorCtx(ctx, err, "could not send quote to client")
		h.registry.Inc(series.Error("send_quote"))

		return
	}

	h.registry.Inc(series.Success())
}

func (h *GetQuoteHandler) writeMessage(conn net.Conn, message string) error {
	_, err := fmt.Fprint(conn, message)
	if err != nil {
		return err
	}
	return nil
}

func (h *GetQuoteHandler) writeError(ctx context.Context, conn net.Conn, errorMessage string) {
	if err := h.writeMessage(conn, fmt.Sprintf("Error: %s\n", errorMessage)); err != nil {
		h.logger.ErrorCtx(ctx, err, "could not write error message to client")
	}
}
