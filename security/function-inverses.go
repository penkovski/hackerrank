package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	num := make(map[int]int)
	for i := 1; i <= n; i++ {
		var f int
		fmt.Scan(&f)
		num[f] = i
	}

	for i := 1; i <= n; i++ {
		fmt.Println(num[i])
	}
}