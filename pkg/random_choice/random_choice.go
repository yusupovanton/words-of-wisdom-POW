package random_choice

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func RandomInt(s1, s2 int) (int, error) {
	if s1 > s2 {
		return 0, errors.New("min must be less than or equal to max")
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(s2-s1+1)))
	if err != nil {
		return 0, err
	}

	return s1 + int(n.Int64()), nil
}
