# Questions

Here are my answers to the questions in the book

## Chapter 1 - Happy Fun Books

### Question 1

> If you run this program, what output you will see? Try to work it out before running the program to check if you were right!

**Response:** You will get the following:
```
The Happines Project
99
```

Verified after executing the program in the playground: <https://play.golang.org/p/ex4TNiPWNF4>

## Chapter 2 - Structural Engineering

### Question 1

> How would you change the Title field of the b variable?

**Response:** `b.Title = "New Title"`

### Question 2

> Suppose you declare a certain field in a struct (for example, Edition int, representing the edition number of the book), but then you don’t supply a value for that field in your struct literal. What happens? Can you refer to that field of the variable? What value does it have? See if you can guess the answer before trying it out in your program. Try omitting the value for various different field types. What value do they get if you don’t explicitly assign one?

**Response:** We can access to the fields of the struct literal even if no value was assigned to them. This is possible since they will hold the "zero value" for their specific type: <https://tour.golang.org/moretypes/5>

## Chapter 3 - Slicing & Dicing

### Question 1

> What would be the type of a slice of int values? How about a slice of strings?

**Response:** For a slice of `int` values we would have something like the following -> `[]int`, while for a slice of strings it would be like -> `[]string`

### Question 2

> What happens if you refer to a slice index which doesn't, in fact, exist? Consider our two-book catalog. Its first element is catalog[0], and its second is catalog[1]. What would happen if you referenced catalog[2] in your program? Do you get nothing at all? An empty Book value? An error? Try it!

**Response:** We will get a `panic: runtime error: index out of range` error

### Question 3

> Given a slice of 3 elements, what is the index of its last element?

**Response:** Index `2` for the last element

> Given a slice x of unknown length, what is the index of its last element, and how does it relate to len(x)?

**Response:** The slice of the last element should be `len(x) - 1`. This is related to the `len` of the slice since these type of structures are "0-based", meaning they start counting the elements from the index `0`, but the `len` function returns the total count of items in the slice in a "1-based" form, so that's the reason we can't simply get the last element of a slice without decreasing the `len` by one first
