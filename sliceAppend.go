package main

import "fmt"

func main() {

	var nums []int

	printSliceDetails(nums)

	nums = append(nums, 10)
	printSliceDetails(nums)

	nums = append(nums, 100)
	printSliceDetails(nums)

	nums = append(nums, 1000, 10000, 100000)
	printSliceDetails(nums)

	nums1 := make([]int, len(nums), (cap(nums) * 2))
	copy(nums1, nums)

	printSliceDetails(nums1)

	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println("Before: ", colors)
	colors = append(colors[1:len(colors)])
	fmt.Println("Items after removig 1st element: ", colors)
}

func printSliceDetails(x []int) {
	fmt.Printf("Length=%d Capacity=%d Slice=%v\n", len(x), cap(x), x)
}
