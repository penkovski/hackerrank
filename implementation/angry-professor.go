package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var tests int64

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		tests, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	for i := 0; i < int(tests); i++ {
		var k int64 // threshold

		if scanner.Scan() {
			s := scanner.Text()
			args := strings.Fields(s)
			k, _ = strconv.ParseInt(args[1], 10, 64)
		}

		if scanner.Scan() {
			var arrivals int64
			s := scanner.Text()
			arrivalStr := strings.Fields(s)
			for _, str := range arrivalStr {
				num, _ := strconv.ParseInt(str, 10, 64)
				if num <= 0 {
					arrivals++
				}
			}

			if arrivals < k {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}
