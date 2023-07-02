package main

import (
	"fmt"
)

func main() {

	//var marks = make([]float64, 3, 5)
	//var marks []float64
	marks := []float32{10, 12.6, 20.0, 37.56, 48.74, 50.0, 64.15, 79.63, 8.75}
	printItemOfSlice(marks)
	fmt.Println("Original Slice = ", marks)
	fmt.Println("Marks[1:5] = ", marks[1:5])
	fmt.Println("Marks[:4] = ", marks[:4])
	fmt.Println("Marks[3:] = ", marks[3:])
	marks1 := make([]float32, 0, 5)
	printItemOfSlice(marks1)
	marks2 := marks[:3]
	printItemOfSlice(marks2)
	marks3 := marks[2:3]
	printItemOfSlice(marks3)

	// if marks == nil {
	// 	fmt.Printf("Slice is Nil")
	// }

}

func printItemOfSlice(x []float32) {
	fmt.Printf("Length=%d Capacity=%d Slice=%v\n", len(x), cap(x), x)
}
