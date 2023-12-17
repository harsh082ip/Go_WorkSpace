package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	var s [5]int
	for x := 0; x < 5; x++ {
		scanner.Scan()
		s[x], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(s)
}
