package repository

// Repository provides methods to access quotes.
type Repository struct{}

// New creates a new Repository instance.
func New() *Repository {
	return &Repository{}
}

// GetQuoteByID is a dummy function that simulates some repository action
func (r *Repository) GetQuoteByID(id int) (string, error) {
	if id < 0 || id >= r.QuotesLength() {
		return "", ErrOutOfRange
	}

	return quotes[id], nil
}

func (r *Repository) QuotesLength() int {
	return len(quotes)
}
