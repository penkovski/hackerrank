package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var nums []int
	for i := 0; i < n; i++ {
		var f int
		fmt.Scan(&f)
		nums = append(nums, f)
	}

	fmt.Println(minimumSwaps(nums))
}

func minimumSwaps(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	var swaps int
	for i := 0; i < n; i++ {
		if nums[i] == i+1 {
			continue
		}

		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		swaps++
		i--
	}

	return swaps
}
