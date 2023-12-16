package main

import (
	"fmt"
	"log"
)

// The display function takes an empty interface (interface{}) as a parameter.
// An empty interface can hold values of any type.
// POLYMORPHISM

func display(v interface{}) {
	s, ok := v.(string)
	if !ok {
		log.Fatal(ok)
	}
	fmt.Printf("%v \n", s)
}

// The display function you provided in your example is using an empty interface (interface{}),
//  which is different from an interface type that defines method signatures.

// In Go, when we use an empty interface (interface{}) as a function parameter,
// it means that the function can accept values of any type. It's a way of writing
// functions that are more flexible and can work with a variety of data types.
// The empty interface is often used in scenarios where you need to handle values of
// different types without explicitly specifying those types.

type shape interface{}

func main() {

	display(34)
	// display("harsh")
	shape := 4
	fmt.Println(shape)
}

// In Go, interfaces serve two primary use cases:

// Method Signature:

// An interface can define a set of method signatures that a type must implement.
// This is the more traditional use of interfaces and is often associated with achieving polymorphism.
// Types that implement the methods declared by an interface are said to satisfy that
//  interface, allowing instances of those types to be used wherever the interface is expected.

// type Shape interface {
//     Area() float64
//     Perimeter() float64
// }

// Empty Interface (Any Type):

// An interface with no method signatures, often referred to as an empty interface
// (interface{}), can hold values of any type. This provides a high degree of
// flexibility and is commonly used in scenarios where you need to work with values of
// various types without specifying them in advance.

// var anyValue interface{}
// anyValue = 42
// anyValue = "hello"

// These two use cases highlight the versatility of interfaces in Go.
// The first use case (method signature) promotes code abstraction, polymorphism, and
// adherence to a contract, while the second use case (empty interface) provides a dynamic
//  mechanism for working with values of diverse types, sacrificing some degree of type safety
//  in the process. The choice of which type of interface to use depends on the specific
//  requirements of your code and design.
