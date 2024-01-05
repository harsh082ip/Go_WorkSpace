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
}

func (u *User) changePass(pass string) {

	u.Pass = pass
}
