package main

import "fmt"

type person struct {
	name     string
	age      int
	id       string
	password string
}

func initialize(p person) {
	p.name = "Kunal"
	// p.age = 18
	// p.id = "fhafst3532"
	// p.password = "qwerty124"

	fmt.Printf("Name: %v ,\n Age: %v ,\n ID: %v, \n Password %v \n", p.name, p.age, p.id, p.password)
}

func main() {
	initialize(person{name: "Harsh", age: 18, id: "gfdffd23", password: "qwrtrw12"})
}
