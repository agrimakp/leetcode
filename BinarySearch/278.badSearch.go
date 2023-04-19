package main

// You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

// Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

// You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.
func isBadVersion(version int) bool {
	bad := 4
	if version == bad {
		return true
	}
	return false
}

func firstBadVersion(n int) int {

	low := 1
	high := n
	var mid int
	var result int

	for low <= high {
		// calculate the middle index
		mid = low + (high-low)/2
		// check whether mid value is a bad version
		isBadVersion := isBadVersion(mid)

		if isBadVersion {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return result
}

func main() {

}
