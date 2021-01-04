package bookstore

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Book describes a book object
type Book struct {
	ID              string
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

// Catalog holds a list of the existing books
var Catalog = map[string]Book{}

// Authors holds a list of the existing books
var Authors = map[string][]string{}

// ABC is the string used for the random generation of unique IDs
const ABC = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetAllBooks returns a list with the existing books
func GetAllBooks() map[string]Book {
	return Catalog
}

// AddBookToCatalog adds a new book to the catalog
func AddBookToCatalog(b Book) {
	Catalog[b.ID] = b
	for _, author := range b.Authors {
		Authors[author] = append(Authors[author], b.ID)
	}
}

// GetBookDetails will return a formatted string with some minor details of the matching Book.ID
func GetBookDetails(ID string) string {
	if book, ok := Catalog[ID]; ok {
		return fmt.Sprintf("@@@@@ Book ID #%s @@@@@\nTitle: %s\nAuthor(s): %s\nPrice: $%d\nDescription: %s\n", book.ID, book.Title, strings.Join(book.Authors, ", "), book.PriceCents, book.Description)
	}

	return ""
}

// GetAllByAuthor will return a slice of books where the provided string is one of the authors of the book
func GetAllByAuthor(findAuthor string) []Book {
	books := []Book{}
	if ids, ok := Authors[findAuthor]; ok {
		for _, id := range ids {
			books = append(books, Catalog[id])
		}
	}

	return books
}

// GetCatalogDetails returns a formatted string with the existing books in the Catalog
func GetCatalogDetails() string {
	var s string
	for _, book := range GetAllBooks() {
		s += GetBookDetails(book.ID)
	}

	return s
}

// NewID generates a new unique identifier which can be used as the ID for a Book
func NewID(size int) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = ABC[rand.Intn(len(ABC))]
	}

	return string(b)
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
