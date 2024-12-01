package server

import (
	"context"
	"fmt"
	"net"

	"github.com/yusupovanton/words-of-wisdom-POW/internal/handlers"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/clog"
)

type Server struct {
	port   string
	logger clog.CLog

	getQuoteHandler *handlers.GetQuoteHandler
}

// NewServer creates a new server instance with the provided configuration.
func NewServer(port string, logger clog.CLog, getQuoteHandler *handlers.GetQuoteHandler) *Server {
	return &Server{
		port:            port,
		getQuoteHandler: getQuoteHandler,
		logger:          logger,
	}
}

// Run starts the server and listens for client connections.
func (s *Server) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return fmt.Errorf("error starting server on port %s: %w", s.port, err)
	}
	defer func() {
		if err = listener.Close(); err != nil {
			s.logger.ErrorCtx(ctx, err, "could not close listener")
		}
	}()

	s.logger.InfoCtx(ctx, "server is running on port %s", s.port)

	for {
		var conn net.Conn

		conn, err = listener.Accept()
		if err != nil {
			s.logger.ErrorCtx(ctx, err, "error accepting connection")

			continue
		}

		go s.getQuoteHandler.GetQuote(ctx, conn)
	}
}
