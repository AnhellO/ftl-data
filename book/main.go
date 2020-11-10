package main

import (
	"fmt"
)

// Book describes a book object
type Book struct {
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
	Edition         int
	SeriesName      string
	SeriesID        int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title: %s\nAuthor: %s\nCopies: %d\nPrice: %d\nDiscount: %d%%\nEdition: %d\nBook #%d of book series %s",
		b.Title, b.Author, b.Copies, b.PriceCents, b.DiscountPercent, b.Edition, b.SeriesID, b.SeriesName,
	)
}

// NetPrice takes a Book value and returns the net price of the book with the discount applied
func NetPrice(b Book) int {
	return b.PriceCents - ((b.PriceCents * b.DiscountPercent) / 100)
}

func main() {
	b := Book{
		Title:           "Ready Player One",
		Author:          "Ernest Cline",
		Copies:          10,
		PriceCents:      100,
		DiscountPercent: 15,
		Edition:         1,
		SeriesName:      "Ready Player One",
		SeriesID:        1,
	}
	fmt.Printf("%s\n", b)
	fmt.Printf("Net Price: %d\n", NetPrice(b))
}
