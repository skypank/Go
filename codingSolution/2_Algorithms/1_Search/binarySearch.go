package main

/*
Binary search algorithm

PROS : The fastest searching algo, runs in O(log n) time complexity

CONS : This Algo works ONLY if the input list is already SORTED.

USAGE : In conjuction with sorting techniques like quicksort

*/

func BinarySearch(array []int, target int) int {
	// initialize left to the starting index and right to the last index of array
	left := 0
	right := len(array) - 1
	// until the left and right and equal, run through the loop
	for left <= right {
		// calcuate the mid index of the array
		mid := (left + right) / 2
		// if the target integer is found at mid position, return mid index
		if target == array[mid] {
			return mid
		} else if target > array[mid] {
			// otherwise check if target is bigger than mid integer, search in second half by increasing left
			left = mid + 1
		} else {
			// if not then target must be smaller than mid integer since array is sorted, so search in first half
			right = mid - 1
		}
	}
	// if after iterating through all the elements a target is not found, return -1 to indicate search item not found
	return -1
}

/*
The driver code is left intentionally blank
TODO : Add driver code for practice
*/
