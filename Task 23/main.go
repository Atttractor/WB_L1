package main

import "fmt"

// Без сохранения порядка
func removeElem1(s []int, i int) []int {
	s[5] = s[0]
	return s[1:]
}

// C сохранением порядка
func removeElem2(s []int, i int) []int {
	return append(s[:i], s[i + 1:]...)
}

// C сохранением порядка, но записываем в новый слайс
func removeElem3(s []int, i int) []int {
	ns := make([]int, 0)
	ns = append(ns, s[:i]...)
	return append(ns, s[i + 1:]...)
}

func main() {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sl)

	fmt.Println(removeElem2(sl, 5))
	fmt.Println(sl)
}