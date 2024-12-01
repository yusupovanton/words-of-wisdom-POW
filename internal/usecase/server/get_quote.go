package server

import (
	randomChoice "github.com/yusupovanton/words-of-wisdom-POW/pkg/random_choice"
)

type storage interface {
	GetQuoteByID(id int) (string, error)
	QuotesLength() int
}

type UseCase struct {
	storage storage
}

// NewUseCase creates a new instance of UseCase with the given storage.
func NewUseCase(storage storage) *UseCase {
	return &UseCase{
		storage: storage,
	}
}

// GetRandomQuote retrieves a random quote from the storage.
func (uc *UseCase) GetRandomQuote() (string, error) {
	length := uc.storage.QuotesLength()
	id, err := randomChoice.RandomInt(1, length)
	if err != nil {
		return "", err
	}

	quote, err := uc.storage.GetQuoteByID(id)
	if err != nil {
		return "", err
	}

	return quote, nil
}
