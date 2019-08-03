package main

import (
	"fmt"
	"strings"
)

func main() {
	var p int
	fmt.Scan(&p)

	for i := 0; i < p; i++ {
		var str1, str2 string
		fmt.Scanln(&str1)
		fmt.Scanln(&str2)

		fmt.Println(twoStrings(str1, str2))
	}
}

func twoStrings(s1 string, s2 string) string {
	checked := make(map[rune]bool)
	for _, char := range s1 {
		if _, ok := checked[char]; !ok {
			checked[char] = true
			if strings.Contains(s2, string(char)) {
				return "YES"
			}
		}
	}
	return "NO"
}
