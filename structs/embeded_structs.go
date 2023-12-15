package main

import "fmt"

// Go is not an object-oriented language. However, embedded structs
// provide a kind of data-only inheritance that can be useful at times.
// Keep in mind, Go doesn't support classes or inheritance in the complete sense,
// embedded structs are just a way to elevate and share fields between struct definitions.

type car struct {
	car_name string
}

type person struct {
	name string
	age  int
	car
}

func main() {

	p := person{
		car:  car{car_name: "AUDI"},
		name: "Harsh",
		age:  18,
	}

	// Another way

	// p.name = "Harsh"
	// p.age = 18
	// p.name = "AUDI"

	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.car_name)

}
