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
