package pow_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yusupovanton/words-of-wisdom-POW/pkg/pow"
)

type PowTestSuite struct {
	suite.Suite
}

func (s *PowTestSuite) TestGenerateChallenge() {
	challenge, err := pow.GenerateChallenge("test", 4)
	s.Require().NoError(err)
	s.NotNil(challenge)
	s.Equal("test", challenge.Prefix)
	s.Equal(4, challenge.Difficulty)

	challenge, err = pow.GenerateChallenge("test", 0)
	s.Require().Error(err)
	s.Nil(challenge)
}

func (s *PowTestSuite) TestCheckSolution() {
	challenge, err := pow.GenerateChallenge("test", 4)
	s.Require().NoError(err)

	challenge.Prefix = ""
	valid, err := challenge.CheckSolution("1234")
	s.Require().Error(err)
	s.False(valid)

	challenge.Prefix = "test"
	solution, err := challenge.FindSolution()
	s.Require().NoError(err)

	valid, err = challenge.CheckSolution(solution)
	s.Require().NoError(err)
	s.True(valid)

	valid, err = challenge.CheckSolution("wrong_nonce")
	s.Require().NoError(err)
	s.False(valid)
}

func (s *PowTestSuite) TestFindSolution() {
	challenge, err := pow.GenerateChallenge("test", 4)
	s.Require().NoError(err)

	solution, err := challenge.FindSolution()
	s.Require().NoError(err)
	s.NotEmpty(solution)

	valid, err := challenge.CheckSolution(solution)
	s.Require().NoError(err)
	s.True(valid)
}

func TestPowTestSuite(t *testing.T) {
	suite.Run(t, new(PowTestSuite))
}
