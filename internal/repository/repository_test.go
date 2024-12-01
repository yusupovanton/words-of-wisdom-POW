package repository_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yusupovanton/words-of-wisdom-POW/internal/repository"
)

type RepositoryTestSuite struct {
	suite.Suite

	repo *repository.Repository
}

func (s *RepositoryTestSuite) SetupTest() {
	s.repo = repository.New()
}

func (s *RepositoryTestSuite) TestGetQuoteByID_Success() {
	expectedQuote := "All we have to decide is what to do with the time that is given us."

	id := 3
	quote, err := s.repo.GetQuoteByID(id)

	s.Require().NoError(err)
	s.Equal(expectedQuote, quote)
}

func (s *RepositoryTestSuite) TestGetQuoteByID_OutOfRange() {
	invalidID := 1000

	quote, err := s.repo.GetQuoteByID(invalidID)

	s.Empty(quote)
	s.ErrorIs(err, repository.ErrOutOfRange)
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
