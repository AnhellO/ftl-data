package main

import (
	"bookstore"
	"fmt"
)

func main() {
	c := bookstore.Catalog{
		{Title: "Harry Potter and the Philosopher's Scone"},
		{Title: "Chocolat 2: Fifty Shades of Chocolat"},
		{Title: "The Catcher in the Pie"},
	}
	fmt.Println(c.GetAllBooks())
}
