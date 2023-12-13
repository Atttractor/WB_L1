package main

import "fmt"

func main() {
	mas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	in := make(chan int)
	out := make(chan int)

	// Запись
	go func(c chan int, m []int) {
		defer close(c)
		for _, elem := range m{
			c <- elem
		}
	}(in, mas)

	// Чтение
	go func(c, o chan int) {
		defer close(out)
		for elem := range in {
			out <- elem * elem
		}
	}(in, out)

	// Вывод
	for elem := range out {
		fmt.Println(elem)
	}
}