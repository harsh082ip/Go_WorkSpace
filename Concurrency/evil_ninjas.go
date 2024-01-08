package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	ninjas := []string{"ninja 1", "ninja 2", "ninja 3", "ninja 4", "ninja 5"}

	for _, ninja := range ninjas {
		go attack(ninja)
	}
}

func attack(target string) {

	fmt.Println("attacking ", target)
	time.Sleep(1 * time.Second)
}
