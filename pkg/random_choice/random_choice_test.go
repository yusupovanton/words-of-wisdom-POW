package random_choice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	randomChoice "github.com/yusupovanton/words-of-wisdom-POW/pkg/random_choice"
)

type RandomChoiceTestSuite struct {
	suite.Suite
}

func (suite *RandomChoiceTestSuite) TestRandomInt() {
	result, err := randomChoice.RandomInt(1, 10)
	assert.NoError(suite.T(), err)
	assert.GreaterOrEqual(suite.T(), result, 1)
	assert.LessOrEqual(suite.T(), result, 10)

	result, err = randomChoice.RandomInt(5, 5)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 5, result)

	result, err = randomChoice.RandomInt(10, 1)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), 0, result)
}

func TestRandomChoiceTestSuite(t *testing.T) {
	suite.Run(t, new(RandomChoiceTestSuite))
}
