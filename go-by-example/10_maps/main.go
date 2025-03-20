package main

import (
	"fmt"
	"maps"
)

func main() {
	// Only declare a variable x with map isn't initialized.
	// zero value is placed and the x is unusable
	var x map[string]int

	// To use variable x we must initialize
	x = make(map[string]int)
	x["x1"] = 1
	fmt.Println("x:", x)

	// shorthand of creating a map
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v2 := m["k2"]
	fmt.Println("v2:", v2)

	fmt.Println("len:", len(m))

	// deleta a key
	delete(m, "k2")
	fmt.Println("map:", m)

	// reset / delete all key
	clear(m)
	fmt.Println("map:", m)

	// test key presence
	_, prs := m["k2"]
	fmt.Println("is_present:", prs)

	n := map[string]int{
		"foo": 1,
		"bar": 2, //when initializing multiline like this, comma must be available on the last item
	}
	fmt.Println("n:", n)

	n2 := map[string]int{"foo": 1, "bar": 2} //when initializing single line, comma is optional
	fmt.Println("n2:", n2)

	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	} else {
		fmt.Println("n != n2")
	}
}
