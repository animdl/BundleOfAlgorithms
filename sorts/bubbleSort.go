package main

import "fmt"

// O(n^2)
// bubble sort ensures that the largest element is always at the end
func bubbleSort(array []int) []int {

	// if the array is nil, empty, or has only one element, it is already sorted
	// and we can return the array
	if array == nil || len(array) <= 1 {
		return array
	}

	// initialize a boolean flag to check if a swap has been made
	swapped := false

	// every iteration, the largest element is moved to the end
	// assumption: the back length - i elements of the array are sorted
	for i := len(array) - 1; i > 0; i-- {
		swapped = false
		// loop through the array
		for j := 0; j < i; j++ {
			// if the current element is greater than the next element, swap
			if array[j] > array[j+1] {
				// swap using syntactic sugar
				array[j], array[j+1] = array[j+1], array[j]
				// if a swap has been made, set swapped to true
				swapped = true
			}
		}

		// if no swap has been made, then we can assume that the array is sorted
		// and we can break
		if !swapped {
			break
		}
	}
	return array
}

func main() {
	fmt.Println(bubbleSort([]int{4, 32, 12, 7, 1, 3}))
}
