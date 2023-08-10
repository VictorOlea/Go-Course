package main

import (
	"fmt"
)

type circle struct {
	length  float32
	breadth float32
	color   string
}

func main() {
	fmt.Println("Without Instantiating: ", circle{10.5, 25.10, "red"})

	var circle1 circle
	circle1.length = 20
	circle1.breadth = 30
	circle1.color = "Yellow"
	fmt.Println("Circle1: ", circle1)

	var circle2 = circle{length: 10, color: "Green"}
	fmt.Println("Circle2: ", circle2)

	circle3 := new(circle)
	circle3.length = 10
	circle3.breadth = 20
	circle3.color = "Red"
	fmt.Println("Circle3: ", circle3)

	var circle4 = new(circle)
	circle4.length = 40
	circle4.color = "Blue"
	fmt.Println("Circle4: ", circle4)

	var circle5 = &circle{10, 20, "Green"}
	fmt.Println(circle5)

	var circle6 = &circle{}
	circle6.length = 10
	circle6.color = "Red"
	fmt.Println(circle6)

	var circle7 = &circle{}
	(*circle7).breadth = 10
	(*circle7).color = "Blue"
	fmt.Println(circle7)
}
