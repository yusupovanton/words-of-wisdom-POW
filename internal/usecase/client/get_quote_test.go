package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	usecase "github.com/yusupovanton/words-of-wisdom-POW/internal/usecase/client"
	"github.com/yusupovanton/words-of-wisdom-POW/internal/usecase/client/mocks"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/clog"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/metrics"
)

type QuoteUseCaseTestSuite struct {
	suite.Suite

	ctx     context.Context
	logger  clog.CLog
	client  *mocks.QuoteGetterClient
	useCase *usecase.QuoteUseCase
}

func (s *QuoteUseCaseTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.logger = clog.NewCLogStub()

	s.client = mocks.NewQuoteGetterClient(s.T())

	s.useCase = usecase.NewQuoteUseCase(s.client, s.logger, metrics.NewRegistryStub())
}

func (s *QuoteUseCaseTestSuite) TestFetchQuote_Success() {
	expectedQuote := "Life is beautiful."

	s.client.EXPECT().GetQuote(mock.Anything).Return(expectedQuote, nil).Once()

	err := s.useCase.FetchQuote(s.ctx)
	s.Require().NoError(err)
}

func (s *QuoteUseCaseTestSuite) TestFetchQuote_ClientError() {
	expectedError := errors.New("client error")

	s.client.EXPECT().GetQuote(mock.Anything).Return("", expectedError).Once()

	err := s.useCase.FetchQuote(s.ctx)
	s.Require().Error(err)
	s.ErrorContains(err, "client error")
}

func TestQuoteUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(QuoteUseCaseTestSuite))
}
