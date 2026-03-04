package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	// We use two pointers to shink from the array into the exact target value
	p1 := 0
	p2 := len(arr) - 1

	for p1 <= p2 {
		mid := (p1 + p2) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			p1 = mid + 1
		} else {
			p2 = mid - 1
		}
	}
	return -1 	// return -1 if no matching found

	// time complexity -> O(log n) (each search reduce the search range into almost half of the current range)
	// space complexity -> O(1) 
}

func main() {
	arr := []int{1, 2, 4, 6, 8, 11, 34, 56, 78, 90}		// We can perform binary search only on sorted array.
	target := 100
	fmt.Println(BinarySearch(arr, target))
}
