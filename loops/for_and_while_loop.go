package main

import "fmt"

func main() {

	// For Loops

	// ----- ALWAYS TRUE ------
	// for {
	// 	fmt.Printf("Always True")
	// }
	// or
	// for true {
	// 	fmt.Printf("Always true")
	// }

	// ----- with condition -----

	// In Go for loop is used to create a while loop
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// for loop
	for num := 1; num <= 100; num++ {
		if num%3 == 0 {
			fmt.Println(num)
		}
	}
}
