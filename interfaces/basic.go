package main

import (
	"fmt"
	"math"
)

// In Go, an interface is a collection of method signatures.
// An interface defines a set of methods that a type must implement
// in order to be considered as implementing that interface.
// Interfaces provide a way to achieve polymorphism in Go, allowing
// different types to be used interchangeably based on their shared
// behavior rather than their explicit type.

type Shape interface {
	Area() float64
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type rectangle struct {
	length  float64
	breadth float64
}

func (r rectangle) Area() float64 {
	return r.length * r.breadth
}

func printArea(s Shape) {
	fmt.Printf("Area: %v \n", s.Area())
}

func main() {
	c := circle{
		radius: 5.8,
	}

	r := rectangle{
		length:  12,
		breadth: 14,
	}

	printArea(c)
	printArea(r)
}

// In this example:

// We define an interface named Shape with a single method signature Area() that returns a float64.
// We create two structs, Circle and Rectangle, and implement the Area method for each of them.
// The printArea function takes a parameter of type Shape, which means it can accept any type that
// implements the Shape interface. Inside printArea, it calls the Area method on the provided shape.
// Interfaces are implicitly satisfied in Go. If a type implements all the methods declared by
//  an interface, it is said to satisfy that interface, and you can use values of that type
//  wherever the interface is expected.

// Interfaces play a crucial role in achieving flexibility and code reusability in Go, a
// llowing you to write functions and methods that can work with different types as long as
// they adhere to a common set of behaviors defined by the interface.
