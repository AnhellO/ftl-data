package main

import (
	"fmt"
)

func main() {
	var author string = "Ernest Cline"
	var edition int = 2
	var hasDiscount bool = true
	var amountDiscount float64 = .10
	fmt.Printf("%s\n%d\n%+v\n%.2f\n", author, edition, hasDiscount, amountDiscount)
}
