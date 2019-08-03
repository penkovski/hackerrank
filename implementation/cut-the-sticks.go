package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int64
	var sticks []int64

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		n, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	if scanner.Scan() {
		s := scanner.Text()
		nums := strings.Fields(s)
		for _, numstr := range nums {
			num, _ := strconv.ParseInt(numstr, 10, 64)
			sticks = append(sticks, num)
		}
	}

	looping(n, &sticks)
}

func looping(stickNum int64, ptrSticks *[]int64) {
	sticks := *ptrSticks
	cutround := 0
	for {
		s := smallest(sticks)
		if s == 0 {
			break
		}

		for i := 0; i < int(stickNum); i++ {
			if sticks[i] > 0 {
				sticks[i] -= s
				cutround++
			}
		}

		if cutround == 0 {
			break
		}

		fmt.Println(cutround)
		cutround = 0
	}
}

func smallest(sticks []int64) int64 {
	s := sticks[0]
	for _, num := range sticks {
		if num > 0 && s == 0 {
			s = num
		}

		if num > 0 && num < s {
			s = num
		}
	}
	return s
}
