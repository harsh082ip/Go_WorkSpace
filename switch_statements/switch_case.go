package main

import (
	"fmt"
)

func main() {

	// similar as if / else
	val := 1
	switch val {
	case 1:
		fmt.Printf("value is %v \n", val)
	case 2:
		fmt.Printf("Value is %v \n", val)
	default:
		fmt.Println("Invalid")
	}

	// check multiple conditions
	x := 6
	switch x {
	case 1, -1, 6:
		fmt.Printf("Value is %v \n", x)
	default:
		fmt.Println("Invalid")
	}
}
