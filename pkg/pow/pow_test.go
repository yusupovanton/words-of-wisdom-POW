package pow_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yusupovanton/words-of-wisdom-POW/pkg/pow"
)

type PowTestSuite struct {
	suite.Suite
}

func (suite *PowTestSuite) TestGenerateChallenge() {
	challenge, err := pow.GenerateChallenge("test", 4)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), challenge)
	assert.Equal(suite.T(), "test", challenge.Prefix)
	assert.Equal(suite.T(), 4, challenge.Difficulty)

	challenge, err = pow.GenerateChallenge("test", 0)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), challenge)
}

func (suite *PowTestSuite) TestCheckSolution() {
	challenge, err := pow.GenerateChallenge("test", 4)
	assert.NoError(suite.T(), err)

	challenge.Prefix = ""
	valid, err := challenge.CheckSolution("1234")
	assert.Error(suite.T(), err)
	assert.False(suite.T(), valid)

	challenge.Prefix = "test"
	solution, err := challenge.FindSolution()
	assert.NoError(suite.T(), err)

	valid, err = challenge.CheckSolution(solution)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), valid)

	valid, err = challenge.CheckSolution("wrong_nonce")
	assert.NoError(suite.T(), err)
	assert.False(suite.T(), valid)
}

func (suite *PowTestSuite) TestFindSolution() {
	challenge, err := pow.GenerateChallenge("test", 4)
	assert.NoError(suite.T(), err)

	solution, err := challenge.FindSolution()
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), solution)

	valid, err := challenge.CheckSolution(solution)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), valid)
}

func TestPowTestSuite(t *testing.T) {
	suite.Run(t, new(PowTestSuite))
}
