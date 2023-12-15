package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	// Guard clause to check for division by zero
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}

	result := a / b
	return result, nil
}

func main() {
	// Example usage
	dividend := 10.0
	divisor := 0.0

	result, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%.2f / %.2f = %.2f\n", dividend, divisor, result)
}
