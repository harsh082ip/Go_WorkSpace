package main

import "fmt"

func main() {

	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, element := range arr {
		fmt.Printf("%v : %v \n", index, element)
	}

}
