package main

import "fmt"

func main() {

	// example of an Anonymous Structs
	// Anonymous Structs - A struct with no name, recommended to use when we
	// want to create only one instance of it
	example := struct {
		name string
		age  int
	}{
		name: "Harsh",
		age:  18,
	}
	fmt.Println(example)
	fmt.Println(example.name)
	fmt.Println(example.age)
}
