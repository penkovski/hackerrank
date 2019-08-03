package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	var nums []int
	var sum int
	for i := 0; i < n; i++ {
		var f int
		fmt.Scan(&f)
		nums = append(nums, f)

		sum += f
	}

	// print the mean
	if len(nums) > 0 {
		fmt.Printf("%.1f\n", float64(sum)/float64(len(nums)))

		// print the median
		fmt.Printf("%.1f\n", median(nums))

		// print the mode
		fmt.Println(mode(nums))
	} else {
		fmt.Println(0)
		fmt.Println(0)
		fmt.Println(0)
	}

}

func median(nums []int) float64 {
	sort.Ints(nums)

	if len(nums)%2 == 0 {
		var n1, n2 int
		n1 = nums[len(nums)/2-1]
		n2 = nums[len(nums)/2]
		return float64(n1+n2) / 2
	}

	return float64(nums[len(nums)/2])
}

func mode(nums []int) int {
	frequency := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		frequency[nums[i]]++
	}

	maxFreqNum := nums[0]
	maxFreq, startFreq := frequency[nums[0]], frequency[nums[0]]

	for k, n := range frequency {
		if n > maxFreq {
			maxFreq = n
			maxFreqNum = k
		} else if n == maxFreq {
			if k < maxFreqNum {
				maxFreqNum = k
			}
		}
	}

	if startFreq == maxFreq {
		sort.Ints(nums)
		return nums[0]
	}

	return maxFreqNum
}
