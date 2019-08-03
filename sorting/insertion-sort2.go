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

	for i := 0; i < int(arrsize)-1; i++ {
		if arr[i] > arr[i+1] {
			for j := i + 1; j > 0; j-- {
				if arr[j] < arr[j-1] {
					arr[j], arr[j-1] = arr[j-1], arr[j]
				}
			}
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

