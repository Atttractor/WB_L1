package main

import (
	"fmt"
	"reflect"
)

//????
func main() {
	m := []any{1, "", func(){}, struct{}{}}
	for _, elem := range m {
		fmt.Println(reflect.TypeOf(elem))
	}
}