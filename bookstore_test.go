package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		Title:           "Ready Player One",
		Authors:         []string{"Ernest Cline"},
		Copies:          10,
		PriceCents:      100,
		DiscountPercent: 15,
		Edition:         1,
		SeriesName:      "Ready Player One",
		SeriesID:        1,
		Featured:        true,
	}
}
