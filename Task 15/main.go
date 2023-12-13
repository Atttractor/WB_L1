package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(size int) string {
	b := strings.Builder{}
	symbol := "界"

	for i := 0; i < size; i++ {
		fmt.Fprint(&b, symbol)
	}

	return b.String()
}

func someFunc() string {
	// Проблема в том, что в памяти хранится строка из 1*2^10 символов, так же срез будет по байтам, а не по символам
	size := 100
	v := createHugeString(1 << 10)
	r := []rune(v)
	justString := make([]rune, size)
	copy(justString, r)
	
	return string(justString)
}

func main() {
	fmt.Println(someFunc())
}
