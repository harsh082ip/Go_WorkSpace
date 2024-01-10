package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

var message string
var wt sync.WaitGroup

func main() {

	message = "Message"

	wt.Add(2)

	go updateMessage("Hey", &mutex)
	go updateMessage("Hello", &mutex)

	wt.Wait()
	fmt.Println(message)

}

func updateMessage(ms string, m *sync.Mutex) {

	defer wt.Done()
	m.Lock()
	message = ms
	// fmt.Println("msg", message)
	m.Unlock()
}
