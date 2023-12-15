package main

import "fmt"

// func add (a int, b int) int {
//     // return a + b;
//     sum := a+b;
//     return sum
// }

func add(a, b int) int {
	// return a + b;
	sum := a + b
	return sum
}

func sum_and_sub(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func main() {
	// --------Function Calling------------

	res := add(5, 10)
	fmt.Println(res)

	sum, sub := sum_and_sub(10, 5)
	fmt.Println(sum, sub)

	// --------Function Calling------------

	// fmt.Println("Hello world")
	// var age int
	// _, err := fmt.Scan(&age)

	// if err != nil {
	//     fmt.Println("-----------")
	//     fmt.Println("Yoyoy")
	//     fmt.Println(err)
	//     fmt.Println("-----------")
	//     return
	// }

	// fmt.Println(age)
	// if age >= 18{
	//     fmt.Println("Can Vote")
	// }

	// if id := 2343222; id > 000011 {
	//     fmt.Println("Valid")
	// }

}
