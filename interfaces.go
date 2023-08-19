package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func main() {
	var s Shape
	s = Rectangle{10, 14}
	fmt.Println("Area of Shape: ", s.Area())
	fmt.Println("Perimeter of Shape: ", s.Perimeter())
}
