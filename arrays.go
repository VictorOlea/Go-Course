package main

import (
	"fmt"
)

func main() {

	var array1 = []float32{10.5, 5.2, 2.88}

	var array2 [10]int
	var i, j int

	//Initializing elements of the array
	for i = 0; i < 10; i++ {
		array2[i] = i + 50
	}

	//Print the value of each elements of array1
	fmt.Println("Elements stored in array1")
	for j = 0; j < 3; j++ {
		fmt.Printf("Element [%d] = %f \n", j, array1[j])
	}

	//Print the value of each elements of array2
	fmt.Println("Elements stored in array2")
	for j = 0; j < 10; j++ {
		fmt.Printf("Elemen [%d] = %d \n", j, array2[j])
	}

	fmt.Println("Size of array1: ", len(array1))
	fmt.Println("Size of array2: ", len(array2))

}
