package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!!"
}

func main() {
	x := 7
	y := &x // it stores address of x
	fmt.Println(x, y)
	*y = 4
	fmt.Println(x, *y, y)

	s := "Hey"
	fmt.Println(s)
	changeValue(&s)
	fmt.Println(s) // value will change coz, func() it is accepting a pointer
}
