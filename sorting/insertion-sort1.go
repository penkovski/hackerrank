package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var arrsize int64
	var arr []int64

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s := scanner.Text()
		arrsize, _ = strconv.ParseInt(s, 10, 64)
	}

	if scanner.Scan() {
		s := scanner.Text()
		args := strings.Fields(s)
		for _, numstr := range args {
			num, _ := strconv.ParseInt(numstr, 10, 64)
			arr = append(arr, num)
		}
	}

	el := arr[arrsize-1]
	for i := arrsize - 1; i >= 0; i-- {
		if i > 0 {
			if arr[i-1] > el {
				arr[i] = arr[i-1]
			} else if arr[i-1] < el {
				arr[i] = el
				printarr(arr)
				break
			}
		} else {
			arr[0] = el
		}
		printarr(arr)
	}
}

func printarr(arr []int64) {
	for _, n := range arr {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

