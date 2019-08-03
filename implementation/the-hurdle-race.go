package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var maxheight int
	for i := 0; i < n; i++ {
		var h int
		fmt.Scan(&h)

		if h > maxheight {
			maxheight = h
		}
	}

	if maxheight > k {
		fmt.Println(maxheight - k)
	} else {
		fmt.Println(0)
	}
}
