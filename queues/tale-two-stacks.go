package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var queue []int
	for i := 0; i < n; i++ {
		var t, x int
		fmt.Scanf("%d %d\n", &t, &x)

		switch t {
		case 1:
			queue = append(queue, x)
		case 2:
			if len(queue) < 2 {
				queue = []int{}
			} else {
				queue = queue[1:]
			}
		case 3:
			if len(queue) > 0 {
				fmt.Println(queue[0])
			}
		}
	}
}
