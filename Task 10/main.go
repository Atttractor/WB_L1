package main

import (
	"fmt"
)

func main() {
	t := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -30}
	fmt.Println(combineTemperature(t))
}

func combineTemperature(t []float64) map[int][]float64 {
	ct := make(map[int][]float64)

	for _, elem := range t {
		k := int(elem / 10) * 10
		ct[k] = append(ct[k], elem)
	}

	return ct
}