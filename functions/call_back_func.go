package main

import (
	"fmt"
	"math"
)

func cube(sqr func(float64) float64) float64 {
	return math.Pow(sqr(2), 3)
}

func square(val float64) float64 {
	return math.Pow(val, 2)
}

func main() {
	fmt.Println(cube(square))
}
