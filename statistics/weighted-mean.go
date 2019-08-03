package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var nums []int
	var weights []int

	for i := 0; i < n; i++ {
		var f int
		fmt.Scan(&f)
		nums = append(nums, f)
	}

	totalWeight := 0
	for i := 0; i < n; i++ {
		var f int
		fmt.Scan(&f)
		weights = append(weights, f)
		totalWeight += f
	}

	totalNums := 0
	for i := 0; i < n; i++ {
		totalNums += nums[i] * weights[i]
	}

	fmt.Printf("%.1f\n", float64(float64(totalNums)/float64(totalWeight)))
}
