package main

import (
	"fmt"
)

func main() {
	var s, t string
	var k int

	fmt.Scanln(&s)
	fmt.Scanln(&t)
	fmt.Scanln(&k)

	// find the position at which the strings start to be different
	var pos int
	for i := 0; i < len(s); i++ {
		if i >= len(s) || i >= len(t) {
			break
		}

		if s[i] != t[i] {
			break
		}

		pos++
	}

	mindel := len(s) - pos
	minapp := len(t) - pos

	if mindel+minapp > k {
		fmt.Println("No")
		return
	}

	if mindel+minapp == k {
		fmt.Println("Yes")
		return
	}

	if len(s) > len(t) {
		if k > mindel {
			if (k-mindel)%2 == 0 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	} else if len(t) > len(s) {
		if k > minapp {
			if (k-minapp)%2 == 0 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	} else {
		if pos == len(s) {
			if k%2 == 0 {
				fmt.Println("Yes")
			} else if k >= len(s)*2 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		} else {
			if k >= minapp+mindel || k%2 == 0 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
