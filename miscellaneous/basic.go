package main

import "fmt"

type User struct {
	ID   string
	Pass string
}

func main() {

	u := User{ID: "142511244", Pass: "Pass123"}
	fmt.Println("Before: ", u)
	// userPointer := &u
	// userPointer.changePass("hello123")

	// or directly this
	u.changePass("hello123")
	fmt.Println("After: ", u)

	// Declare a variable of type int
	var x int = 42

	// Declare a pointer that points to the address of x
	var ptr *int = &x

	// Use the * operator to dereference the pointer and retrieve the value
	value := *ptr

	// Print the value
	fmt.Println("Value:", value) // Output: Value: 42
}

func (u *User) changePass(pass string) {

	u.Pass = pass
}
