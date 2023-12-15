package main

import "fmt"

func demo() (x, y int) {
	return
}

// as same as above one
func demo1() (int, int) {
	var x int
	var y int
	return x, y
}

func print_demo() {
	d1, d2 := demo()
	fmt.Printf("demo() ===> %v & %v \n", d1, d2)
	d1, d2 = demo1()
	fmt.Printf("demo1() ===> %v & %v", d1, d2)
}

func main() {
	print_demo()
}
