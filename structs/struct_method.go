package main

import "fmt"

// While Go is not object-oriented, it does support methods
// that can be defined on structs. Methods are just functions
// that have a receiver. A receiver is a special parameter that
// syntactically goes before the name of the function.

type Person struct {
	name string
	age  int
}

// getAge has a receiver of (p Person)
func (p Person) getAge() int {
	return p.age
}

// A receiver is just a special kind of function parameter.
// Receivers are important because they will, as you'll learn
// in the exercises to come, allow us to define interfaces that
// our structs (and other types) can implement.

func (p Person) getName() string {
	return p.name
}

func main() {
	pr := Person{
		name: "Harsh",
		age:  18,
	}

	fmt.Println(pr.getAge())
	fmt.Println(pr.getName())
}
