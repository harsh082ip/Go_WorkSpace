package main

import (
	"fmt"
)

// 2-D Array, (rows, columns)

func main() {

	arr := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr)
	fmt.Println(len(arr))
}
