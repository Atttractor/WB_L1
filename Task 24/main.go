package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func newPoint(x float64, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func get_distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.x - p2.x), 2) + math.Pow(float64(p1.y - p2.y), 2))
}

func main() {
	p1 := newPoint(10, 20)
	p2 := newPoint(0, 0)
	fmt.Println(get_distance(p1, p2))
}