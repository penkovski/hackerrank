package main

import (
	"fmt"
)

func main() {
	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var s string
		fmt.Scanln(&s)
		fmt.Println(alternate(s))
	}
}

func alternate(s string) int {
	deletions := 0
	repeated := 0

	for i, c := range s {
		if i == 0 {
			continue
		}

		if uint8(c) == s[i-1] {
			repeated++
			continue
		}

		if repeated > 0 {
			deletions += repeated
			repeated = 0
		}
	}

	if repeated > 0 {
		deletions += repeated
	}

	return deletions
}
