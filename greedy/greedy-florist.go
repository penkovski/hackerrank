package main

import (
	"fmt"
	"sort"
)

func main() {

	var n, k int
	fmt.Scan(&n, &k)

	var prices []int
	for i := 0; i < n; i++ {
		var p int
		fmt.Scan(&p)
		prices = append(prices, p)
	}

	fmt.Println(mincost(k, prices))
}

func mincost(k int, prices []int) int {
	sort.Ints(prices)

	purchases := make(map[int]int)
	var currPerson int
	var mincost int

	for i := len(prices) - 1; i >= 0; i-- {
		mincost += prices[i] * (1 + purchases[currPerson])
		purchases[currPerson]++
		currPerson = (currPerson + 1) % k
	}

	return mincost
}
