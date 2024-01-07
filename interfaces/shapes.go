package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type rectangle struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

func (r rectangle) getArea() float64 {

	return r.length * r.breadth
}

func (c circle) getArea() float64 {

	return math.Pi * c.radius * c.radius
}

func printAreaofShape(s shape) {

	fmt.Println(s.getArea())
}

func main() {

	r := rectangle{length: 12, breadth: 15}
	c := circle{radius: 19.5}

	printAreaofShape(r)
	printAreaofShape(c)
}
