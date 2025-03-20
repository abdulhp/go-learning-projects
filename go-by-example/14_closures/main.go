package main

import "fmt"

func intSeq() func() int {
	i := 0

	// closure only capture the variable that is used inside of it.
	return func() int {
		// https://go.dev/doc/faq#inc_dec
		// increment/decrement is a statement in go, not an expression
		// It's expected, increment and decrement is modifying the variable
		i++

		// this returns i value to the caller of the function
		// not returning to i variable
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}