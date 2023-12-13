package main

import "fmt"

func main() {
	mn1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mn2 := []int{5, 7, 6}
	fmt.Println(intersection(mn1, mn2))
}

func intersection(m1, m2 []int) []int {
	mapa := make(map[int]int)

	for _, elem := range append(m1, m2...) {
		mapa[elem] += 1
	}

	res := make([]int, 0)
	
	for i := range mapa {
		if mapa[i] > 1 {
			res = append(res, i)
		}
	}

	return res
}