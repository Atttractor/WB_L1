package main

import (
	"fmt"
)

func main() {
	txt := "123456"
	fmt.Println(reverse(txt))
}

func reverse(s string) string {
	runes := []rune(s)

	for i := len(runes) - 1; i >= len(runes) / 2; i-- {
		runes[i], runes[len(runes) - i - 1] = runes[len(runes) - i - 1], runes[i]
	}

	return string(runes)
}