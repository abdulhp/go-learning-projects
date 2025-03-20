package main

import "fmt"

// in OOP imagine a struct is a class
type rect struct {
    width, height int
}

// meanwhile a modified function is called method
// method in go quite different form with regular OOP class method
func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

	// both of methods is auto converted between reference to value
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}