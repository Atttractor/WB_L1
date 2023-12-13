package main

import "fmt"

type Earthling struct {
	weight float32
}

func (p *Earthling) WeightOnEarth() float32 {
	return p.weight
}

type Martian interface {
	WeightOnMars() float32
}

type Adapter struct {
	e Earthling
}

func (a *Adapter) WeightOnMars() float32 {
	return a.e.WeightOnEarth() / 2.64
}

func main() {
	misha := Earthling{weight: 264}

	var vasia Martian = &Adapter{
		e: misha,
	}

	fmt.Println(vasia.WeightOnMars())
}