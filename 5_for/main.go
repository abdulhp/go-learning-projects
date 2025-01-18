package main

import (
	"fmt"
)

func main() {

	fmt.Println("For Loop with only condition")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println()

	fmt.Println("For Loop with full components")
	for j := 1; j < 3; j++ {
		fmt.Println(j)
	}
	fmt.Println()

	fmt.Println("For Range")
	for i := range 3 {
		fmt.Println("range", i)
	}
	fmt.Println()
	
	fmt.Println("for with no components A.K.A unlimited loop")
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println()

	fmt.Println("For range with continue")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	fmt.Println()
}
