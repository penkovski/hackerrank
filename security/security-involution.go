package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	funcs := make(map[int]int)
	for i := 1; i <= n; i++ {
		var f int
		fmt.Scan(&f)
		funcs[i] = f
	}

	for i := 1; i <= n; i++ {
		if funcs[funcs[i]] != i {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
