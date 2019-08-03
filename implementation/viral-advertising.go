package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int64
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s := scanner.Text()
		n, _ = strconv.ParseInt(s, 10, 64)
	}

	var reach int = 5
	var likes int = reach / 2

	if n == 1 {
		fmt.Println(likes)
		return
	}

	likes = recursive(1, n, 2)

	fmt.Println(likes)
}

func recursive(day int64, n int64, likes int) int {
	if day > n {
		return 0
	}
	var dayLikes = (likes * 3) / 2

	return recursive(day+1, n, dayLikes) + likes
}
