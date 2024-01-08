package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(9)

	ran := []string{
		"hlo",
		"hii",
		"something1",
		"something2",
		"something3",
		"something4",
		"something5",
		"something6",
		"something7",
	}

	for i, val := range ran {
		go printSomething(fmt.Sprintf("%v : %v", i, val), &wg)
	}

	wg.Wait()

	// if i don't do this, all 9 wait groups are already subtracted
	// and it will cause a negative wait group error
	wg.Add(1)
	printSomething("last msg", &wg)
}

func printSomething(msg string, wg *sync.WaitGroup) {

	fmt.Println(msg)
	wg.Done()
	// here wait groups are getting subtracted one by one
}
