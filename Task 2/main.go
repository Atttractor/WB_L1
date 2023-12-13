package main

import (
	"fmt"
	"sync"
)

func main() {
	m := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(5)

	for _, elem := range m {
		go func(num int) {
			fmt.Println(num * num)
			wg.Done()
		}(elem)
	}

	wg.Wait()
}