package main

import (
	"fmt"
)

// Book describes a book object
type Book struct {
	Title           string
	Authors         []string
	Copies          int
	PriceCents      int
	DiscountPercent int
	Edition         int
	SeriesName      string
	SeriesID        int
	Featured        bool
}

func (b Book) String() string {
	str := fmt.Sprintf(
		"Title: %s\nAuthor(s): %s\nCopies: %d\nPrice: %d\nDiscount: %d%%\nEdition: %d",
		b.Title, b.Authors, b.Copies, b.PriceCents, b.DiscountPercent, b.Edition,
	)

	if b.SeriesName != "" {
		str = fmt.Sprintf("%s\nBook #%d of book series '%s'", str, b.SeriesID, b.SeriesName)
	}

	return str
}

// NetPrice takes in a Book value and returns the net price of the book with the discount applied
func NetPrice(b Book) int {
	if b.DiscountPercent > 0 {
		return b.PriceCents - ((b.PriceCents * b.DiscountPercent) / 100)
	}

	return b.PriceCents
}

// PrintCatalog takes in a slice of books and prints each of the books with their NetPrice
func PrintCatalog(catalog []Book) {
	for i, book := range catalog {
		fmt.Printf("##### Book #%d #####\n", i+1)
		fmt.Println(book)
		fmt.Printf("Net Price: %d\n", NetPrice(book))
	}
	fmt.Printf("Total books in the provided catalog: %d\n", len(catalog))
}

// AddToCatalog takes in a book and a slice of books as parameters, and adds the new book to the slice
func AddToCatalog(catalog []Book, b Book) []Book {
	return append(catalog, b)
}

// FeaturedBooks takes in a slice of books and returns a new slice with the books marked as 'Featured'
func FeaturedBooks(catalog []Book) []Book {
	featured := make([]Book, 0)
	for _, b := range catalog {
		if b.Featured {
			featured = append(featured, b)
		}
	}

	return featured
}

func main() {
	catalog := []Book{
		{
			Title:           "Ready Player One",
			Authors:         []string{"Ernest Cline"},
			Copies:          10,
			PriceCents:      100,
			DiscountPercent: 15,
			Edition:         1,
			SeriesName:      "Ready Player One",
			SeriesID:        1,
			Featured:        true,
		},
		{
			Title:           "The Martian",
			Authors:         []string{"Andy Weir"},
			Copies:          50,
			PriceCents:      120,
			DiscountPercent: 5,
			Edition:         1,
			Featured:        true,
		},
	}

	// PrintCatalog(catalog)

	// Intentionally trying to access a non-existing index per suggestion on the book
	// fmt.Printf("Net Price: %d\n", NetPrice(catalog[2]))

	// Appending a new book to the catalog
	newBook := Book{
		Title:           "Dune",
		Authors:         []string{"Frank Herbert"},
		Copies:          200,
		PriceCents:      150,
		DiscountPercent: 0,
		Edition:         2,
		SeriesName:      "Dune",
		SeriesID:        1,
	}

	catalog = AddToCatalog(catalog, newBook)
	PrintCatalog(catalog)

	// Filter out non-featured books
	featured := FeaturedBooks(catalog)
	fmt.Println("····· FEATURED BOOKS OF THE MONTH ·····")
	PrintCatalog(featured)
}
