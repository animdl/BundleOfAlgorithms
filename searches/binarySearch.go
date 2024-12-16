package main

import "fmt"

// O(log n)
// assumes array is sorted
func binarySearch(array []int, target int) int {
	// initialize low and high indexes
	lowIndex, highIndex := 0, len(array)-1

	// check bounds
	if lowIndex > highIndex || target < array[lowIndex] || target > array[highIndex] {
		return -1 // sentinel value returned if target is out of bounds
	}

	// run binary search as long as low index is less than or equal to high index
	for lowIndex <= highIndex {
		// calculate midpoint between low and high indexes
		midIndex := (lowIndex + highIndex) / 2

		if target == array[midIndex] { // checks if target is at midpoint
			return midIndex
		} else if target > array[midIndex] { // checks if target is greater than midpoint
			lowIndex = midIndex + 1
		} else { // checks if target is less than midpoint
			highIndex = midIndex
		}
	}
	return -1 // sentinel value returned if target is not found within array
}

func main() {
	fmt.Println(binarySearch([]int{1, 3, 4, 7, 12, 32}, 32))
}
