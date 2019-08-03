package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	fmt.Scanln(&input)

	if len(input) == 0 {
		fmt.Println(0)
		return
	}

	words := 1
	for _, r := range input {
		if unicode.IsUpper(r) {
			words += 1
		}
	}

	fmt.Println(words)
}

