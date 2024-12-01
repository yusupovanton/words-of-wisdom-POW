package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type Challenge struct {
	Prefix     string
	Difficulty int
}

func GenerateChallenge(prefix string, difficulty int) (*Challenge, error) {
	if difficulty < 1 {
		return nil, errors.New("difficulty must be greater than zero")
	}
	return &Challenge{Prefix: prefix, Difficulty: difficulty}, nil
}

func (c *Challenge) CheckSolution(nonce string) (bool, error) {
	if c.Prefix == "" || c.Difficulty < 1 {
		return false, errors.New("invalid challenge parameters")
	}

	data := fmt.Sprintf("%s%s", c.Prefix, nonce)
	hash := sha256.Sum256([]byte(data))
	hashHex := hex.EncodeToString(hash[:])

	requiredPrefix := strings.Repeat("0", c.Difficulty)
	return strings.HasPrefix(hashHex, requiredPrefix), nil
}

func (c *Challenge) FindSolution() (string, error) {
	if c.Prefix == "" || c.Difficulty < 1 {
		return "", errors.New("invalid challenge parameters")
	}

	nonce := big.NewInt(0)
	for {
		valid, err := c.CheckSolution(nonce.String())
		if err != nil {
			return "", err
		}
		if valid {
			return nonce.String(), nil
		}

		nonce.Add(nonce, big.NewInt(1))
	}
}
