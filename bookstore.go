package bookstore

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
