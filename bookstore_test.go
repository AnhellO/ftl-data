package bookstore_test

import (
	"bookstore"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		ID:              bookstore.NewID(24),
		Title:           "Ready Player One",
		Authors:         []string{"Ernest Cline"},
		Description:     "Cool 80's Sci-Fi book",
		Copies:          10,
		PriceCents:      100,
		DiscountPercent: 15,
		Edition:         1,
		SeriesName:      "Ready Player One",
		SeriesID:        1,
		Featured:        true,
	}
}

func TestSliceOfStrings(t *testing.T) {
	// Intentionally failing test in order to show the go-cmp package in action
	want := []string{"same", "same", "same"}
	got := []string{"same", "diff", "same"}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestNewID(t *testing.T) {
	const SIZE = 100
	ids := make(map[string]struct{}, SIZE)

	for i := 0; i < SIZE; i++ {
		id := bookstore.NewID(24)
		if len(id) != 24 {
			t.Errorf("invalid ID length, expected '24', got %d", len(id))
		}

		ids[id] = struct{}{}
	}

	if len(ids) != SIZE {
		t.Fatalf("expected '%d' IDs, got %d", SIZE, len(ids))
	}
}

func TestGetAllBooks(t *testing.T) {
	book1 := bookstore.Book{
		ID:              "Book1",
		Title:           "Ready Player One",
		Authors:         []string{"Ernest Cline"},
		Description:     "Cool 80's Sci-Fi book",
		Copies:          10,
		PriceCents:      100,
		DiscountPercent: 15,
		Edition:         1,
		SeriesName:      "Ready Player One",
		SeriesID:        1,
		Featured:        true,
	}
	book2 := bookstore.Book{
		ID:              "Book2",
		Title:           "The Martian",
		Authors:         []string{"Andy Weir"},
		Description:     "A guy gets stuck in Mars",
		Copies:          50,
		PriceCents:      120,
		DiscountPercent: 5,
		Edition:         1,
		Featured:        true,
	}
	bookstore.Catalog = map[string]bookstore.Book{"Book1": book1, "Book2": book2}
	want := map[string]bookstore.Book{"Book1": book1, "Book2": book2}
	got := bookstore.GetAllBooks()

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAddBookToCatalog(t *testing.T) {
	bookstore.Catalog = map[string]bookstore.Book{}
	bookstore.Authors = map[string][]string{}
	book := bookstore.Book{
		ID:              "Book1",
		Title:           "Ready Player One",
		Authors:         []string{"Ernest Cline"},
		Description:     "Cool 80's Sci-Fi book",
		Copies:          10,
		PriceCents:      100,
		DiscountPercent: 15,
		Edition:         1,
		SeriesName:      "Ready Player One",
		SeriesID:        1,
		Featured:        true,
	}
	bookstore.AddBookToCatalog(book)
	want := map[string]bookstore.Book{"Book1": book}
	got := bookstore.GetAllBooks()

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookDetails(t *testing.T) {
	bookstore.Catalog = map[string]bookstore.Book{}
	bookstore.Authors = map[string][]string{}
	id := bookstore.NewID(24)
	book := bookstore.Book{
		ID:              id,
		Title:           "Ready Player One",
		Authors:         []string{"Ernest Cline"},
		Description:     "Cool 80's Sci-Fi book",
		Copies:          10,
		PriceCents:      100,
		DiscountPercent: 15,
		Edition:         1,
		SeriesName:      "Ready Player One",
		SeriesID:        1,
		Featured:        true,
	}
	bookstore.AddBookToCatalog(book)
	want := fmt.Sprintf("@@@@@ Book ID #%s @@@@@\nTitle: Ready Player One\nAuthor(s): Ernest Cline\nPrice: $100\nDescription: Cool 80's Sci-Fi book\n", book.ID)
	got := bookstore.GetBookDetails(id)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	// Add another book
	id = bookstore.NewID(24)
	book = bookstore.Book{
		ID:              id,
		Title:           "The Martian",
		Authors:         []string{"Andy Weir"},
		Description:     "A guy gets stuck in Mars",
		Copies:          50,
		PriceCents:      120,
		DiscountPercent: 5,
		Edition:         1,
		Featured:        false,
	}
	bookstore.AddBookToCatalog(book)
	want = fmt.Sprintf("@@@@@ Book ID #%s @@@@@\nTitle: The Martian\nAuthor(s): Andy Weir\nPrice: $120\nDescription: A guy gets stuck in Mars\n", book.ID)
	got = bookstore.GetBookDetails(id)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetAllByAuthor(t *testing.T) {
	bookstore.Catalog = map[string]bookstore.Book{}
	bookstore.Authors = map[string][]string{}
	book1 := bookstore.Book{
		ID:              bookstore.NewID(24),
		Title:           "Ready Player One",
		Authors:         []string{"Ernest Cline"},
		Description:     "Cool 80's Sci-Fi book",
		Copies:          10,
		PriceCents:      100,
		DiscountPercent: 15,
		Edition:         1,
		SeriesName:      "Ready Player One",
		SeriesID:        1,
		Featured:        false,
	}
	book2 := bookstore.Book{
		ID:              bookstore.NewID(24),
		Title:           "The Martian",
		Authors:         []string{"Andy Weir"},
		Description:     "A guy gets stuck in Mars",
		Copies:          50,
		PriceCents:      120,
		DiscountPercent: 5,
		Edition:         1,
		Featured:        false,
	}
	book3 := bookstore.Book{
		ID:              bookstore.NewID(24),
		Title:           "Armada",
		Authors:         []string{"Ernest Cline"},
		Description:     "Another cool 80's Sci-Fi book",
		Copies:          3,
		PriceCents:      200,
		DiscountPercent: 0,
		Edition:         1,
		Featured:        true,
	}
	book4 := bookstore.Book{
		ID:              bookstore.NewID(24),
		Title:           "Some Random Invented Book",
		Authors:         []string{"Ernest Cline", "Andy Weir"},
		Description:     "Doesn't exist yet, but hopefully soon",
		Copies:          3,
		PriceCents:      200,
		DiscountPercent: 0,
		Edition:         1,
		Featured:        true,
	}
	bookstore.AddBookToCatalog(book1)
	bookstore.AddBookToCatalog(book2)
	bookstore.AddBookToCatalog(book3)
	bookstore.AddBookToCatalog(book4)
	want := []bookstore.Book{book1, book3, book4}
	got := bookstore.GetAllByAuthor("Ernest Cline")

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetCatalogDetails(t *testing.T) {
	bookstore.Catalog = map[string]bookstore.Book{}
	bookstore.Authors = map[string][]string{}
	book1 := bookstore.Book{
		ID:              "Book1",
		Title:           "Ready Player One",
		Authors:         []string{"Ernest Cline"},
		Description:     "Cool 80's Sci-Fi book",
		Copies:          10,
		PriceCents:      100,
		DiscountPercent: 15,
		Edition:         1,
		SeriesName:      "Ready Player One",
		SeriesID:        1,
		Featured:        false,
	}
	book2 := bookstore.Book{
		ID:              "Book2",
		Title:           "The Martian",
		Authors:         []string{"Andy Weir"},
		Description:     "A guy gets stuck in Mars",
		Copies:          50,
		PriceCents:      120,
		DiscountPercent: 5,
		Edition:         1,
		Featured:        false,
	}
	bookstore.AddBookToCatalog(book1)
	bookstore.AddBookToCatalog(book2)
	want := "@@@@@ Book ID #Book1 @@@@@\nTitle: Ready Player One\nAuthor(s): Ernest Cline\nPrice: $100\nDescription: Cool 80's Sci-Fi book\n" +
		"@@@@@ Book ID #Book2 @@@@@\nTitle: The Martian\nAuthor(s): Andy Weir\nPrice: $120\nDescription: A guy gets stuck in Mars\n"
	got := bookstore.GetCatalogDetails()

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
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
