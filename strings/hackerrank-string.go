package main

import (
	"fmt"
)

func main() {
	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var input string
		fmt.Scan(&input)
		if check(input) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func check(s string) bool {
	hackrank := []string{"h", "a", "c", "k", "e", "r", "r", "a", "n", "k"}
	hindex := 0

	for _, r := range s {
		if string(r) == hackrank[hindex] {
			hindex++
		}
        
        if hindex == len(hackrank) {
            break
        }
	}

	if hindex == len(hackrank) {
		return true
	}

	return false
}
