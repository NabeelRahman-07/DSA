// This is a classic algorithm that used to sort an array containing three distinct values.
//It is commonly used in problems like:
//  - Sorting arrays with 0, 1, 2
//  - Partitioning arrays into three groups
//  - Solving the Sort Colors problem (leetcode)

// Time complexity :- O(n)
// Space complexity :- O(1)

package main

import "fmt"

func DutchSort(arr []int) []int {
	low := 0
	mid := 0
	high := len(arr) - 1

	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else {
			arr[high], arr[mid] = arr[mid], arr[high]
			high--
		}
	}
	return arr
}

func main() {
	arr := []int{0, 2, 1, 2, 0, 1, 2, 0, 1, 0, 2, 0, 2, 0, 1}
	fmt.Println(DutchSort(arr))
}
