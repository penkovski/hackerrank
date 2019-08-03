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
	var k int64
	var numbers []int64

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s := scanner.Text()
		args := strings.Fields(s)
		n, _ = strconv.ParseInt(args[0], 10, 64)
		k, _ = strconv.ParseInt(args[1], 10, 64)
	}

	if scanner.Scan() {
		s := scanner.Text()
		nums := strings.Fields(s)
		for _, numstr := range nums {
			num, _ := strconv.ParseInt(numstr, 10, 64)
			numbers = append(numbers, num)
		}
	}

	var pairs int = 0

	for i := 0; i < int(n); i++ {
		for j := i + 1; j < int(n); j++ {
			sum := numbers[i] + numbers[j]
			if sum%k == 0 {
				pairs++
			}
		}
	}

	fmt.Println(pairs)
}
