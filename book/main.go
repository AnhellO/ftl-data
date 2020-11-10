package main

import (
	"fmt"
)

// Book describes a book object
type Book struct {
	Title      string
	Author     string
	Copies     int
	Cost       float64
	Discount   float64
	Edition    int
	SeriesName string
	SeriesID   int
}

func main() {
	fmt.Printf("%+v\n", Book{
		Title:      "Ready Player One",
		Author:     "Ernest Cline",
		Copies:     10,
		Cost:       10,
		Discount:   0,
		Edition:    1,
		SeriesName: "Ready Player One",
		SeriesID:   1,
	})
}
