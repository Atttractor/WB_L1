package main

import "fmt"

func main() {
	m := []int{2, 4, 6, 8, 10}
	c := make(chan int)

	for _, elem := range m {
		go square_num(c, elem)
	}

	var sum int

	for i := 0; i < len(m); i++ {
		sum += <-c
	}

	fmt.Println(sum)
}

func square_num(c chan int, num int) {
	c <- num * num
}