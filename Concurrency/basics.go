package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	now := time.Now()

	// fmt.Println(fetchUserData(122))
	// fmt.Println(fetchUserRecommendation(244))
	// fmt.Println(fetchUserLikes(212))

	// there is a method by which we can reduce the time, bu using goroutines and channels
	respch := make(chan string, 128)
	userID := 122
	wg := &sync.WaitGroup{}

	go fetchUserData(userID, respch, wg)
	wg.Add(1)
	go fetchUserRecommendation(userID, respch, wg)
	wg.Add(1)
	go fetchUserLikes(userID, respch, wg)
	wg.Add(1)

	wg.Wait()

	close(respch)

	for res := range respch {

		fmt.Println(res)
	}

	fmt.Println(time.Since(now))

}

func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {

	time.Sleep(90 * time.Millisecond)
	respch <- "UserData"

	wg.Done()
}

func fetchUserRecommendation(userID int, respch chan string, wg *sync.WaitGroup) {

	time.Sleep(130 * time.Millisecond)
	respch <- "User Recommendation"

	wg.Done()
}

func fetchUserLikes(userID int, respch chan string, wg *sync.WaitGroup) {

	time.Sleep(110 * time.Millisecond)
	respch <- "User Likes"

	wg.Done()
}
