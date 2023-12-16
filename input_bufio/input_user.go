package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// scanner := bufio.NewScanner(os.Stderr)
	fmt.Printf("Type something: ")
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())
	// input, _ = strconv.Atoi(input)
	fmt.Printf("You Typed: %v \n	", input+2)
}
