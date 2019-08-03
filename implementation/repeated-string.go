package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s string
	var n int64

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s = scanner.Text()
	}

	if scanner.Scan() {
		ns := scanner.Text()
		n, _ = strconv.ParseInt(ns, 10, 64)
	}

	// count the letters 'a' in the input string
	countA := 0
	for _, let := range s {
		if let == 'a' {
			countA++
		}
	}

	repeated := int(n) / len(s)
	result := repeated * countA

	// now we should only calculate the last part
	index := 0
	for _, let := range s {
		if index >= int(n)%len(s) {
			break
		}

		index += 1
		if let == 'a' {
			result += 1
		}
	}

	fmt.Println(result)
}
