package main

import "fmt"

func main() {
	var arr []int = []int{2, 3, 4, 5, 6, 3, 6, 4, 67, 87, 90}

	for index, element := range arr {
		for index2, element2 := range arr[index+1:] {
			if element == element2 {
				fmt.Printf("%v \n", element)
			}
			index2 += index + 1
		}
	}
}
