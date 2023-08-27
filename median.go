package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Median of {56, 85, 92} is: ", median([]float64{56, 85, 92}))
	fmt.Println("Median of {56, 85, 92, 99} is: ", median([]float64{56, 85, 92, 99}))
}

func median(nums []float64) float64 {
	vals := make([]float64, len(nums))
	copy(vals, nums)
	sort.Float64s(vals)

	i := len(vals) / 2
	if len(vals)%2 == 1 {
		return vals[i]
	}
	return (vals[i-1] + vals[i]) / 2
}
