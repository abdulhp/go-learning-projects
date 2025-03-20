package main

import "fmt"

// as mentioned in variables section
// we can declare multiple variable in a single line
// we can return multiple values in a function
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println("a:",a)
	fmt.Println("b:",b)

	// underscore means we ignore the value
	_, c := vals()
	fmt.Println("c:",c)
}