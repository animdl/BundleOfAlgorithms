package main

import "fmt"

func insertionSort(array []int) []int {

	// if the array is nil, empty, or has only one element, it is already sorted
	// and we can return the array
	if array == nil || len(array) <= 1 {
		return array
	}

	// every iteration, the sub array is sorted
	// assumption: the first i elements of the array are sorted
	for i := 1; i < len(array); i++ {
		// loop through the first i elements of the array
		for j := i; j > 0; j-- {
			// if the current element is less than the previous element, swap
			if array[j] < array[j-1] {
				// swap using syntactic sugar
				array[j], array[j-1] = array[j-1], array[j]
			} else {
				// if no swap has been made, then we can assume that the sub array is sorted
				break
			}
		}
	}
	return array
}

func main() {
	fmt.Println(insertionSort([]int{4, 32, 12, 7, 1, 3}))
}
