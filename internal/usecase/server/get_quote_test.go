package server_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/yusupovanton/words-of-wisdom-POW/internal/usecase/server"
	"github.com/yusupovanton/words-of-wisdom-POW/internal/usecase/server/mocks"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/clog"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/metrics"
)

type UseCaseTestSuite struct {
	suite.Suite

	ctx     context.Context
	logger  clog.CLog
	storage *mocks.Storage
	useCase *server.UseCase
}

func (s *UseCaseTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.logger = clog.NewCLogStub()

	s.storage = mocks.NewStorage(s.T())

	s.useCase = server.NewUseCase(s.logger, s.storage, metrics.NewRegistryStub())
}

func (s *UseCaseTestSuite) TestGetRandomQuote_Success() {
	s.storage.EXPECT().QuotesLength().Return(10).Once()
	s.storage.EXPECT().GetQuoteByID(mock.AnythingOfType("int")).Return("Random Quote", nil).Once()

	quote, err := s.useCase.GetRandomQuote(s.ctx)
	s.Require().NoError(err)
	s.Equal("Random Quote", quote)
}

func (s *UseCaseTestSuite) TestGetRandomQuote_StorageError() {
	s.storage.EXPECT().QuotesLength().Return(10).Once()
	s.storage.EXPECT().GetQuoteByID(mock.AnythingOfType("int")).Return("", errors.New("repository error")).Once()

	quote, err := s.useCase.GetRandomQuote(s.ctx)
	s.Require().Error(err)
	s.Empty(quote)
	s.ErrorContains(err, "repository error")
}

func (s *UseCaseTestSuite) TestGetRandomQuote_InvalidStorageLength() {
	s.storage.EXPECT().QuotesLength().Return(0).Once()

	quote, err := s.useCase.GetRandomQuote(s.ctx)
	s.Require().Error(err)
	s.Empty(quote)
}

func TestUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UseCaseTestSuite))
}
