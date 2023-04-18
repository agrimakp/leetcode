package main

import "fmt"

func isPerfectSquare(num int) bool {
	low := 1
	high := num / 2

	for low <= high {
		mid := low + (high-low)/2
		mul := mid * mid

		if mul == num {
			return true
		}
		if mul > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func main() {
	num := 16
	fmt.Println(isPerfectSquare(num))
}
