package main

import "fmt"

func main() {
	m := []int{10, 5, 11, 2, 7, 9, 100, 100, 500, 4, 40, 150, 2}
	fmt.Println(m)
	fmt.Println(quickSort(m))
}

func quickSort(m []int) []int {
	if len(m) < 1 {
		return m
	}

	pivot := m[len(m) / 2]
	m_less := []int{}
	m_more := []int{}
	result := []int{}

	for i, elem := range m {
		if i != len(m) / 2 {
			if elem < pivot {
				m_less = append(m_less, elem)
			} else {
				m_more = append(m_more, elem)
			}
		}
	}

	result = append(result, quickSort(m_less)...)
	result = append(result, pivot)
	result = append(result, quickSort(m_more)...)
	return result
}