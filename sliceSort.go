package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := make([]int, 5)
	intSlice[0] = 6
	intSlice[1] = 99
	intSlice[2] = 45
	intSlice[3] = 34
	intSlice[4] = 1

	fmt.Println("Original Slice: ", intSlice)
	sort.Ints(intSlice)
	fmt.Println("Sorted Slice: ", intSlice)

}
