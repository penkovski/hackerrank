package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	var sum int
	var nums []int
	for i := 0; i < n; i++ {
		var f int
		fmt.Scan(&f)
		nums = append(nums, f)
		sum += f
	}

	// find the mean
	mean := float64(sum) / float64(len(nums))

	distance := 0.0
	for i := 0; i < n; i++ {
		d := float64(nums[i]) - mean
		distance += math.Pow(d, 2)
	}

	result := math.Sqrt(distance / float64(len(nums)))
	fmt.Printf("%.1f\n", result)
}
