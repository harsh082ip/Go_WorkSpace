package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a buffered channel with a capacity of 3
	ch := make(chan int, 5)

	// Start a goroutine to send values to the channel
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("Sent %d to the channel\n", i)
		}
		close(ch) // close the channel when done sending
	}()

	// Allow some time for the sender goroutine to operate
	time.Sleep(2 * time.Second)

	// Receive values from the channel
	for {
		value, ok := <-ch
		if !ok {
			// Channel is closed, and no more values will be received
			break
		}
		fmt.Printf("Received %d from the channel\n", value)
	}

}

// Output:
// Sent 1 to the channel
// Sent 2 to the channel
// Sent 3 to the channel
// Received 1 from the channel
// Received 2 from the channel
// Sent 4 to the channel
// Sent 5 to the channel
// Received 3 from the channel
// Received 4 from the channel
// Received 5 from the channel

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	c := make(chan string, 3)
// 	go sendValues("Hey", c)

// 	getValues := <-c
// 	fmt.Println(getValues)

// }

// func sendValues(val string, ch chan string) {
// 	time.Sleep(2 * time.Second)
// 	ch <- val
// 	ch <- val
// 	ch <- val
// }
