package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printTable(num int) {

	for i := 1; i <= 10; i++ {
		fmt.Printf("%v * %v = %v \n", num, i, (num * i))
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter no: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	printTable(num)
}
