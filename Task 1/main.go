package main

import "fmt"

type Human struct {
	Name  string
	Age   int
	Phone string
}

func (h Human) calcAge() int {
	return h.Age
}

type Action struct {
	Human
}

func (a Action) calcAge() int {
	return a.Human.calcAge()
}


func main() {
	h := Human{
		Name:  "Вася Пупкин",
		Age:   18,
		Phone: "+79021234567",
	}

	a := Action{
		Human: h,
	}
	fmt.Println(a.calcAge())
	fmt.Println(h.calcAge())
}