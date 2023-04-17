package main

// Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).
// Constraints: 0 <= low <= high <= 10^9

func countOdds(low int, high int) int {
	if low%2 == 0 && high%2 == 0 {
		return (high - low) / 2
	}
	return (high-low)/2 + 1
}

func main() {

}
