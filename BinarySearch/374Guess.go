package main

func guessNumber(n int) int {
	low := 1
	high := n

	for low <= high {
		mid := low + (high-low)/2
		if guess(mid) == -1 {
			high = mid - 1
		} else if guess(mid) == 1 {
			low = mid + 1

		} else if guess(mid) == 0 {
			return mid
		}
	}
	return 1

}

func main() {

}
