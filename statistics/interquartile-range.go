package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	var nums []int
	var s []int

	for i := 0; i < n; i++ {
		var f int
		fmt.Scan(&f)
		nums = append(nums, f)
	}

	for i := 0; i < n; i++ {
		var f int
		fmt.Scan(&f)

		for j := 0; j < f; j++ {
			s = append(s, nums[i])
		}
	}

	sort.Ints(s)

	var result float64
	if len(s)%2 == 0 {
		result = quartilesf(s[len(s)/2:]) - quartilesf(s[:len(s)/2])
	} else {
		result = quartilesf(s[len(s)/2+1:]) - quartilesf(s[:len(s)/2])
	}

	fmt.Printf("%.1f\n", result)
}

func quartilesf(nums []int) float64 {
	if len(nums)%2 == 0 {
		n1 := nums[len(nums)/2-1]
		n2 := nums[len(nums)/2]
		return float64(n1+n2) / float64(2)
	}
	return float64(nums[len(nums)/2])
}
