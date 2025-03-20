package main

import "fmt"

// Variadic function is a function that have trailing arguments
// Yes, trailing. It must be on the last arguments
// We can do (name string, nums ...int)
// but not (nums ...int, name string)
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
