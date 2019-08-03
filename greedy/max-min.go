package main

import (
	"fmt"
	"sort"
)

func main() {

	var n, k int
	fmt.Scanln(&n)
	fmt.Scanln(&k)

	var arr []int
	for i := 0; i < n; i++ {
		var f int
		fmt.Scanln(&f)
		arr = append(arr, f)
	}

	fmt.Println(maxMin(k, arr))
}

// Complete the maxMin function below.
func maxMin(k int, arr []int) int {
	sort.Ints(arr)

	minfair := -1

	for i := k - 1; i < len(arr); i++ {
		f := fairness(arr[i-k+1 : i+1])
		if minfair == -1 || minfair > f {
			minfair = f
		}
	}

	return minfair
}

func fairness(subarr []int) int {
	if len(subarr) < 2 {
		return -1
	}

	min, max := subarr[0], subarr[len(subarr)-1]

	return max - min
}

//1234
//
//10
//100
//300
//200
//1000
//20
//30
//

//k = 3
//
//10, 20, 30, 100, 200, 300, 1000
