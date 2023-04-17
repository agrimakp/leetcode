package main

func average(salary []int) float64 {
	minNumber := 1000000
	maxNumber := 1000
	total := 0
	count := float64(len(salary) - 2)

	for _, s := range salary {
		total += s
		if s < minNumber {
			minNumber = s
		}
		if s > maxNumber {
			maxNumber = s
		}

	}
	return float64(total-minNumber-maxNumber) / float64(count)
}

func main() {

}
