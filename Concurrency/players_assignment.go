package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		c <- "ball"
	}()

	// Wait for the goroutine to finish sending to the channel
	wg.Wait()

	// Receive the value from the channel
	sendToPlayer1("w", c)
}

func sendToPlayer1(ball string, ch chan string) {
	fmt.Println(<-ch)
	fmt.Println("ball received from player 2")
}
