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

	//fmt.Println(naive(nums))
	fmt.Println(greedy(nums))
}

func greedy(nums []int) int {
	sort.Ints(nums)

	mindif := -1
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 {
			dif := absdif(nums[i-1], nums[i])
			if mindif == -1 || dif < mindif {
				mindif = dif
			}
		}
		dif := absdif(nums[i], nums[i+1])
		if mindif == -1 || dif < mindif {
			mindif = dif
		}
	}
	return mindif
}

func naive(nums []int) int {
	checked := make(map[int]int)
	mindif := -1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if checked[nums[i]] == nums[j] || checked[nums[j]] == nums[i] {
				continue
			}
			checked[nums[i]] = nums[j]

			diff := absdif(nums[i], nums[j])
			if mindif == -1 || diff < mindif {
				mindif = diff
			}
		}
	}
	if mindif < 0 {
		return 0
	}
	return mindif
}

func absdif(a, b int) int {
	r := a - b
	if r < 0 {
		return -r
	}
	return r
}
