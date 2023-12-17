package main

import "fmt"

func test() {
	fmt.Println("Hey")
}

func test2(myFunc func(int) int) {
	fmt.Println(myFunc(3))
}

func returnfunc() func() {
	return func() { fmt.Println("Function is returned") }
}

func returnfunc2(str string) func() {
	return func() {
		fmt.Printf("%v is sent \n", str)
	}
}

func main() {
	// we can assign a function to a variable and call it using that variable
	x := test
	x()

	// we can initialize a func() to a var directly
	fun := func() {
		fmt.Printf("Hello world \n")
	}
	fun()

	// Another way of calling
	f := func(x int) int {
		return x * x
	}(5)
	fmt.Println(f)

	// we can send a func() as an argument in another func()
	f1 := func(x int) int {
		return x * x
	}
	test2(f1)

	// simply calling a function in main
	func() {
		fmt.Println("Hello, World")
	}()

	// if i want to call the func() that is returned, we'll use ()()
	returnfunc()()
	y := returnfunc2("ID")
	y()
}
