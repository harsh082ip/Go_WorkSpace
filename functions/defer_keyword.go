package main

import "fmt"

func display() {

	// it defers the execution of only the below line until return is executed
	defer fmt.Println("Exiting from display()")
	fmt.Println("Inside display()")
	return
}

func main() {
	display()
}
