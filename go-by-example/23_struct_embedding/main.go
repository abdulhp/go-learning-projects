/* 
==== STRUCT EMBEDDING ===
To understand struct embedding,
we must know the definition of embed.
Embed means include or attached an item.

Struct embedding means a process to
embed a behavior/property to inherited object/var.

Resulting an embedded behaviour/property on a var.
*/ 
package main

import "fmt"

// suppose a type named base
type base struct {
	num int
}

// base have describe method
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// now suppose another type that have base type
type container struct {
	base
	str string
}

func main() {

	fmt.Println()
	// define a data for base
	a_base := base{
		num: 1,
	}
	// now you can access the num property and describe method
	fmt.Println("num:", a_base.num)
	fmt.Println("desc:", a_base.describe())

	fmt.Println()

	// Now, we put a_base into container. Allowing container
	// to directly access the a_base property and method.
	// Or we could say, the a_base property and method owned by container
	co := container{
		base: a_base,
		str: "some name",
	}

	// Direct access to base type without using base property name in container
	// this is struct embedded
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// This is the normal way to access the num
	fmt.Println("also num:", co.base.num)

	// Method also embedded into container type
	fmt.Println("describe:", co.describe())

	// Suppose we have an interface of describe method
	type describer interface {
		describe() string
	}

	// now, each time d accessing describe method, its embedding co
	var d describer = co
	fmt.Println("describer:", d.describe())
}
