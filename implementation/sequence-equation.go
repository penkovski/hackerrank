package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	mapping := make(map[int]int, 0)

	for i := 1; i <= n; i++ {
		var num int
		fmt.Scan(&num)
		mapping[num] = i
	}

	for x := 1; x <= n; x++ {
		var p = mapping[x]
		fmt.Println(mapping[p])
	}
}
