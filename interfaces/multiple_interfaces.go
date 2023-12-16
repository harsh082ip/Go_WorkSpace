package main

import (
	"fmt"
	"math"
)

// In Go, a type can implement multiple interfaces.
// This is known as "implicit interface satisfaction." When a type provides
// implementations for all the methods declared by two or more interfaces, it is
// considered to implement all of those interfaces.

// ---------------- Interfaces -------------------
type Shape interface {
	Area() float64
}

type PerimeterCalculator interface {
	Perimeter() float64
}

// -------------------------------------------------

// ---------------- CIRCLE-------------------

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// ---------------- CIRCLE-------------------

// ---------------- RECTANGLE -------------------
type rectangle struct {
	length  float64
	breadth float64
}

func (r rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

// ---------------- RECTANGLE -------------------

func printArea(s Shape, name string) {
	fmt.Printf("Area of %v: %v \n", name, s.Area())
}

func PrintPerimeter(pr PerimeterCalculator, name string) {
	fmt.Printf("Perimeter of %v: %v \n", name, pr.Perimeter())
}

func main() {
	c := circle{radius: 4.5}
	r := rectangle{
		length:  14,
		breadth: 16,
	}

	printArea(c, "circle")
	printArea(r, "rectangle")

	PrintPerimeter(c, "circle")
	PrintPerimeter(r, "rectangle")
}

// In this example:

// We define two interfaces, Shape and PerimeterCalculator, each with a method.
// The Circle and Rectangle types implement both interfaces by providing implementations
// for the required methods.
// The printArea function takes a Shape parameter, and the printPerimeter function
// takes a PerimeterCalculator parameter. Both functions can accept instances of types
// that implement the respective interfaces.
// This demonstrates how a type in Go can be versatile and provide different behaviors
//  depending on the interfaces it implements. It allows for clean and modular code by
//  promoting flexibility and reuse of code across different contexts.
