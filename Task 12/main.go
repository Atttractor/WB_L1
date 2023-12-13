package main

import "fmt"

func main() {
	m := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(m)
	fmt.Println(set(m))
}

func set(m []string) []string {
	mapa := make(map[string]int)
	result := []string{}

	for _, elem := range m {
		if _, ok := mapa[elem]; ok {
			continue
		}
		mapa[elem] = 0
	}

	for k := range mapa {
		result = append(result, k)
	}

	return result
}