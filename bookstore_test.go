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

func TestSalePriceCents(t *testing.T) {
	b := bookstore.Book{
		Title:      "A Clockwork Orange Soda",
		PriceCents: 500,
	}
	want := 250
	got := b.SalePriceCents()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCustomerMailingLabel(t *testing.T) {
	c := bookstore.Customer{
		Title:   "Mr.",
		Name:    "A.J.",
		Address: "Street X #123, Neighbour Hood, MX",
	}
	want := "Customer: Mr. A.J.\nAddress: Street X #123, Neighbour Hood, MX"
	got := c.MailingLabel()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
