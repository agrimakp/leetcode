package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	value := math.Sqrt(float64(x))
	return int(value)
}

func main() {
	x := 4
	fmt.Println(mySqrt(x))
}
