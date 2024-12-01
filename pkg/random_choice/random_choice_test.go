package random_choice_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	randomChoice "github.com/yusupovanton/words-of-wisdom-POW/pkg/random_choice"
)

type RandomChoiceTestSuite struct {
	suite.Suite
}

func (s *RandomChoiceTestSuite) TestRandomInt() {
	result, err := randomChoice.RandomInt(1, 10)
	s.Require().NoError(err)
	s.GreaterOrEqual(result, 1)
	s.LessOrEqual(result, 10)

	result, err = randomChoice.RandomInt(5, 5)
	s.Require().NoError(err)
	s.Equal(5, result)

	result, err = randomChoice.RandomInt(10, 1)
	s.Require().Error(err)
	s.Equal(0, result)
}

func TestRandomChoiceTestSuite(t *testing.T) {
	suite.Run(t, new(RandomChoiceTestSuite))
}
