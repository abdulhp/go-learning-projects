package main

import (
	"fmt"
)

// function declared with `func` statement
// func <function name>(<args>) <return type> {}
func plus(a int, b int) int {
	return a + b
}

// we can also save some space
// var a will be defined as int
// as long as the last arg have data type,
// the earlier declaration will follows the data type
func plus1(a, b int, c float32) int {
	// casted floats into ints will floor the value
	return a + b + int(c)
}

func main() {
	a := plus(1, 1)
	fmt.Println(a)

	b := plus1(1, 2, 0.7)
	fmt.Println(b)
}
