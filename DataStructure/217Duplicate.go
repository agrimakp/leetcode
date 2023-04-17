package main

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)

	for _, number := range nums {
		if set[number] {
			return true
		}
		set[number] = true
	}
	return false
}

func main() {

}
