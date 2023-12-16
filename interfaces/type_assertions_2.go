package main

import (
	"fmt"
	"log"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}
type rectangle struct {
	len     float64
	breadth float64
}

// Implement the area method for the circle type
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func (r rectangle) area() float64 {
	return r.len * r.breadth
}

func main() {
	// Now, s is of the shape interface type
	var s shape = circle{
		radius: 6.2,
	}

	// Type assertion to check if s is a circle
	c, ok := s.(circle)
	if !ok {
		// s wasn't a circle
		log.Fatal("s is not a circle")
	}

	// Now that we know s is a circle, we can access its properties
	radius := c.radius
	fmt.Println("Radius:", radius)

	// You can also call the area method through the interface
	area := s.area()
	fmt.Println("Area:", area)
}
