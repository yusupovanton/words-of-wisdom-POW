package random_choice

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func RandomInt(min, max int) (int, error) {
	if min > max {
		return 0, errors.New("min must be less than or equal to max")
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}

	return min + int(n.Int64()), nil
}
