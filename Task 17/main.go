package main

import "fmt"

func main() {
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 500, 1000}
	fmt.Println(binarySearch(m, 1))
}

func binarySearch(m []int, k int) int {
	left := 0
	right := len(m) - 1
	mid := right / 2

	for left <= right {
		if k > m[mid] {
			left = mid + 1
		} else if k < m[mid] {
			right = mid - 1
		} else {
			return mid
		}
		mid = (left + right) / 2
	}

	return -1
}