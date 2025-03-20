package main

import "fmt"

func main() {
    nums := []int{2, 3, 4}
    sum := 0

	// range of a slice/arrays will returns multiple values
	// we could ignore the indexes by putting underscore
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

	// by default, if only one variable assigned
	// it will refer to index/key only
    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}