package main

import (
	"fmt"
	"math"
)

// O(sqrt(n))
// assumes array is sorted
func jumpSearch(array []int, target int) int {

	// initialize jump amount
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(array)))))
	startIndex := 0

	// find the section that contains the target
	// run jump search until the target is lesser than the value at [start + jump amount]
	for target > array[startIndex+jumpAmount] {
		startIndex += jumpAmount

		// check if startIndex + jump amount is out of bounds
		// if so, break, since the section is set to the last section
		if startIndex+jumpAmount >= len(array) {
			break
		}
	}

	// if the section is the last section, set the min to the length of the array
	min := min(startIndex+jumpAmount, len(array))

	// run linear search until target is found
	for i := startIndex; i < min; i++ {
		if array[i] == target {
			return i
		}
	}

	return -1 // sentinel value returned if target is not found
}

func main() {
	fmt.Println(jumpSearch([]int{1, 3, 4, 7, 12, 32}, 32))
}
