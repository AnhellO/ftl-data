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

// Catalog represents a catalog of books.
type Catalog []Book

// MappedCatalog holds a list of the existing books
var MappedCatalog = map[string]Book{}

// Authors holds a list of the existing books
var Authors = map[string][]string{}

// ABC is the string used for the random generation of unique IDs
const ABC = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetAllMappedBooks returns a list with the existing books
func GetAllMappedBooks() map[string]Book {
	return MappedCatalog
}

// GetAllBooks returns the number of the books on the catalog
func (c Catalog) GetAllBooks() []Book {
	return []Book(c)
}

// Len returns the slice of all books
// in the Catalog.
func (c Catalog) Len() int {
	return len(c.GetAllBooks())
}

// GetAllTitles returns the a slice with the titles
// for all the books in the Catalog
func (c Catalog) GetAllTitles() []string {
	books := make([]string, 0)
	for _, v := range c.GetAllBooks() {
		books = append(books, v.Title)
	}

	return books
}

// GetUniqueAuthors returns a slice containing all
// the unique authors of the books in the Catalog
func (c Catalog) GetUniqueAuthors() []string {
	authors := make([]string, 0)
	for _, v := range c.GetAllBooks() {
		for _, a := range v.Authors {
			if inSlice(a, authors) {
				continue
			}

			authors = append(authors, a)
		}
	}

	return authors
}

func inSlice(author string, authors []string) bool {
	for _, a := range authors {
		if author == a {
			return true
		}
	}

	return false
}

// AddBookToCatalog adds a new book to the catalog
func AddBookToCatalog(b Book) {
	MappedCatalog[b.ID] = b
	for _, author := range b.Authors {
		Authors[author] = append(Authors[author], b.ID)
	}
}

// GetBookDetails will return a formatted string with some minor details of the matching Book.ID
func GetBookDetails(ID string) string {
	if book, ok := MappedCatalog[ID]; ok {
		return fmt.Sprintf("@@@@@ Book ID #%s @@@@@\nTitle: %s\nAuthor(s): %s\nPrice: $%d\nDescription: %s\n", book.ID, book.Title, strings.Join(book.Authors, ", "), book.PriceCents, book.Description)
	}

	return ""
}

// GetAllByAuthor will return a slice of books where the provided string is one of the authors of the book
func GetAllByAuthor(findAuthor string) []Book {
	books := []Book{}
	if ids, ok := Authors[findAuthor]; ok {
		for _, id := range ids {
			books = append(books, MappedCatalog[id])
		}
	}

	return books
}

// GetCatalogDetails returns a formatted string with the existing books in the Catalog
func GetCatalogDetails() string {
	var s string
	for _, book := range GetAllMappedBooks() {
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
