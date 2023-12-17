package main

import "fmt"

func main() {

	/* Slices in go works on the top of an array,
	In simple terms it is just a small part
	extracted from an array

	To declare a slice we will use []

	ARRAY :- {1, 2, 3, 4}
	len = 4
	capacity = 4

	Slice form it = {2, 3}
	len = 2
	capacity = 3, (can be extended to one more element)
	*/

	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3]
	fmt.Println(len(s)) // len = 2
	fmt.Println(cap(s)) // cap = 4, extendible to one element
	fmt.Println(s)
	fmt.Println(s[1:])
	fmt.Println(s[1:len(s)])

	var num []int = []int{}
	// var num2 []int = []int{}
	num = append(num, 10)
	fmt.Println(num)
	// fmt.Println(num2)

	// make() function

	a := make([]int, 5)
	fmt.Println(a)
}
