package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		x int64 = 175 // 10101111 -> 10111111 = 191
		numOfBit int64 = 6
		bitLen int
		res int64
	)
	// i-ый справа
	//m := int64(math.Pow(2, float64(numOfBit)))

	// i-ый слева
	for i := 0; int64(math.Pow(2, float64(i))) <= x; i++ {
		bitLen = i
	}
	bitLen++

	m := int64(math.Pow(2, float64(int64(bitLen) - numOfBit)))

	res = x ^ m
	fmt.Println(res)
}