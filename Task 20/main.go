package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "snow dog sun 123"
	fmt.Println(reverseWordsInString(text))
}

func reverseWordsInString(s string) string {
	words := strings.Split(s, " ")

	for i := len(words) - 1; i >= len(words)/2; i-- {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}

	return strings.Join(words, " ")
}

func reverse(s string) string {
	runes := []rune(s)

	for i := len(runes) - 1; i >= len(runes)/2; i-- {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}

	return string(runes)
}