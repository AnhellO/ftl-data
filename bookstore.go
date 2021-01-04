package bookstore

import (
	"fmt"
)

// Book describes a book object
type Book struct {
	Title           string
	Authors         []string
	Description     string
	Genre           string
	Copies          int
	PriceCents      int
	DiscountPercent int
	Edition         int
	SeriesName      string
	SeriesID        int
	Featured        bool
}

// SalePriceCents returns the sale price of a book, which is the 50% of the normal price
func (b Book) SalePriceCents() int {
	return b.PriceCents / 2
}

// Customer describes a customer of the bookstore
type Customer struct {
	Title   string
	Name    string
	Address string
}

// MailingLabel returns a string containing the customer’s title, name, and address,
// suitable for printing out to a mailing label.
func (c Customer) MailingLabel() string {
	return fmt.Sprintf("Customer: %s %s\nAddress: %s", c.Title, c.Name, c.Address)
}
