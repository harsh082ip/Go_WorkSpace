package main

import (
	"fmt"
	"sync"
)

// In Go (or Golang), a race condition occurs when two or more goroutines
// (concurrently executing threads) access shared data concurrently, and at least
// one of them modifies the data. If proper synchronization mechanisms are not
// in place, the outcome of the program becomes unpredictable and might lead to unexpected behavior.

var msg string

var wg sync.WaitGroup

func main() {

	msg = "Hello World"

	wg.Add(2)
	go updateMsg("Hey")
	go updateMsg("Hlo")

	wg.Wait()

	fmt.Println(msg) // output is "hey" (maybe), this is a race condition
}

func updateMsg(s string) {
	defer wg.Done()
	msg = s
}
