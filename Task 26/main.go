package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "qwertgaQ"
	fmt.Println(unique(str))
}

func unique(str string) bool {
	m := map[rune]struct{}{}
	str = strings.ToLower(str)

	for _, elem := range str {
		if _, ok := m[elem]; ok {
			return false
		}
		m[elem] = struct{}{}
	}
	return true
}