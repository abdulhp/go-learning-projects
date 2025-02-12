package main

import "fmt"


func main () {
	fmt.Println(mergeAlternately("abc"))
}

func mergeAlternately(word1 string) string {
	
	var merged string

	for *i := &word1; i != \0; i++ {
		merged += i
	}

	return merged
}
