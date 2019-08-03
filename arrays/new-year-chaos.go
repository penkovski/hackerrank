package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scanln(&t)

	for test := 0; test < t; test++ {
		var n int
		fmt.Scanln(&n)

		var queue []int
		for i := 0; i < n; i++ {
			var f int
			fmt.Scan(&f)
			queue = append(queue, f)
		}

		r := solve(queue)
		if r < 0 {
			fmt.Println("Too chaotic")
			continue
		}

		fmt.Println(r)
	}
}

func solve(queue []int) int {
	var bribes int

	for i := len(queue) - 1; i >= 0; i-- {
		if queue[i]-(i+1) > 2 {
			return -1
		}

		for j := int(math.Max(float64(0), float64(queue[i]-2))); j < i; j++ {
			if queue[j] > queue[i] {
				bribes++
			}
		}
	}

	return bribes
}
