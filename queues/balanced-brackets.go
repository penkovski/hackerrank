package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var brackets string
		fmt.Scanln(&brackets)
		fmt.Println(balanced(brackets))
	}
}

func balanced(s string) string {
	opened := 0
	closed := 0
	stack := make([]rune, len(s))
	for _, r := range s {
		if opening(r) {
			opened++
			stack = append(stack, r)
			continue
		}

		if len(stack) == 0 {
			return "NO"
		}

		last := stack[len(stack)-1]
		if match(last, r) {
			closed++
			stack = stack[:len(stack)-1]
		} else {
			return "NO"
		}
	}
	if opened == closed {
		return "YES"
	}
	return "NO"
}

func opening(r rune) bool {
	if r == '(' || r == '{' || r == '[' {
		return true
	}
	return false
}

func match(a, b rune) bool {
	if (a == '{' && b == '}') || (a == '(' && b == ')') || (a == '[' && b == ']') {
		return true
	}
	return false
}
