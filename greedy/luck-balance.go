package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var major []int
	var minor []int
	for i := 0; i < n; i++ {
		var luck, importance int
		fmt.Scan(&luck, &importance)

		if importance == 1 {
			major = append(major, luck)
		} else {
			minor = append(minor, luck)
		}
	}

	sort.Ints(major)

	fmt.Println("major =", major)
	fmt.Println("minor =", minor)

	maxluck := 0
	for i := len(major) - 1; i >= 0; i-- {
		if i >= len(major)-k {
			maxluck += major[i]
		} else {
			maxluck -= major[i]
		}
	}

	for i := 0; i < len(minor); i++ {
		maxluck += minor[i]
	}

	fmt.Println(maxluck)
}
