package main

import (
	"fmt"
	"sort"
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

	sort.Ints(nums)

	fmt.Println(quartiles(nums[:len(nums)/2]))
	fmt.Println(quartiles(nums))

	if len(nums)%2 == 0 {
		fmt.Println(quartiles(nums[len(nums)/2:]))
	} else {
		fmt.Println(quartiles(nums[len(nums)/2+1:]))
	}
}

func quartiles(nums []int) int {
	if len(nums)%2 == 0 {
		n1 := nums[len(nums)/2-1]
		n2 := nums[len(nums)/2]
		return (n1 + n2) / 2
	}
	return nums[len(nums)/2]
}
