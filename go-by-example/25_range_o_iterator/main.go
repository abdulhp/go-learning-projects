/*
===== RANGE OVER ITERATOR =====
By default, range only support built-in data types
such as arrays, slices, maps, strings, and channels.

If we want to loop custom data type using range
we have to implement iterator which implemented
in this file.
*/
package main

import (
    "fmt"
    "iter"
    "slices"
)

// suppost a type is defined with generic type parameter
type element[T any] struct {
    next *element[T]
    val  T
}

// the list type also defined with generic type parameter
// so element is embedded to List type
type List[T any] struct {
    head, tail *element[T]
}

// a list has a method called Push
// this method modify original instance
// when pushing an integer
func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

// Suppose you want to loop all of nodes
// using for...range methods. Because the
// range is only supports built in datatypes,
// we have to define an iterator

// An iterator is a callback function of iter.Seq[T]
func (lst *List[T]) All() iter.Seq[T] {

	// returns the callback
	// yield value is determined by caller
	// if caller defined a terminating
	// condition/who kills the looping
    return func(yield func(T) bool) {

		// this block is has terminating condition
		// means in parent/caller doesn't have to 
		// define a terminating condition
        for e := lst.head; e != nil; e = e.next {
			// yield returns true only if 
			// the looping is continued implicitly by range
            if !yield(e.val) {
				// return is loop termination
                return
            }
        }
    }
}

func genFib() iter.Seq[int] {
    return func(yield func(int) bool) {
        a, b := 1, 1

        for {
            if !yield(a) {
                return
            }
            a, b = b, a+b
        }
    }
}

func main() {
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)

    for e := range lst.All() {
        fmt.Println(e)
    }

    all := slices.Collect(lst.All())
    fmt.Println("all:", all)

    for n := range genFib() {

        if n >= 10 {
            break
        }
        fmt.Println(n)
    }
}