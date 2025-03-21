package main

import (
	"fmt"
	"math"
)

// interface is similar to abstract class
// if a struct methods doesn't have this interface methods
// it isn't called as geometry and throw an error
type geometry interface {
	area() float64
	perim() float64
}

// defining class/struct
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// defining a methods
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// accepting g as defined as geometry
// which a var needs to have 2 methods with float64 returns
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// we can assert g is a circle using comma-ok
// or else we can assert g type with switch case
// switch shape := g.(type)
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// calls measure function with rect/circle
	// cause rect/circle has the requirement of interface geometry
	//
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
