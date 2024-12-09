package main

import (
	"fmt"
)

func main() {
	// read number of candles
	var n int
	fmt.Scanln(&n)

	// read candle heights
	heights := map[int]int{}
	for i := 0; i < n; i++ {
		var h int
		fmt.Scanf("%d", &h)
		heights[h]++
	}

	// find number of max height candles
	maxHeight, maxHeightCount := 0, 0
	for h, count := range heights {
		if h > maxHeight {
			maxHeight = h
			maxHeightCount = count
		}
	}

	fmt.Println(maxHeightCount)
}
