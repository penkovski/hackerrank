package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n int64 // numbers
	var k int64 // rotations
	var q int64 // queries

	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	line, isPrefix, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}

	if isPrefix == false {
		s := string(line)
		split := strings.Fields(s)
		n, _ = strconv.ParseInt(split[0], 10, 64)
		k, _ = strconv.ParseInt(split[1], 10, 64)
		q, _ = strconv.ParseInt(split[2], 10, 64)
	} else {
		fmt.Println("buffer size is too small")
		return
	}

	var rotations = k % n

	//fmt.Printf("n = %v, k = %v, q = %v\n", n, k, q)

	var numbers []int64

	line, isPrefix, err = reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err == nil && isPrefix == false {
		s := string(line)
		split := strings.Fields(s)
		//		fmt.Println(split)
		for _, strnum := range split {
			num, _ := strconv.ParseInt(strnum, 10, 64)
			numbers = append(numbers, num)
		}
	}

	for i := 0; i < int(q); i++ {
		line, isPrefix, err = reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			return
		}

		s := string(line)
		query, _ := strconv.ParseInt(s, 10, 64)

		var index int
		if query-rotations >= 0 {
			index = int(query) - int(rotations)
		} else {
			index = int(query) - int(rotations) + len(numbers)
		}

		fmt.Println(numbers[int(index)])
	}

	// this solution does not really rotate the elements of the slice, but
	// insted uses a trick to print the elements as if they are rotated
	// it is super fast in comparison with the commented solution below

	// NOTE: the commented solution below performs real array elements rotation
	// and because of that is slow and takes more memory

	//	rotateNaive(&numbers, int(k))
	//	rotateFaster(&numbers, int(k))

	//	fmt.Println(numbers)
	//	for i := 0; i < int(q); i++ {
	//		line, isPrefix, err = reader.ReadLine()
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}

	//		s := string(line)
	//		query, _ := strconv.ParseInt(s, 10, 64)
	//		//		fmt.Println(numbers[int(query)])
	//	}
}

func rotateNaive(arr *[]int64, rotations int) {
	a := *arr
	for i := 0; i < rotations; i++ {
		for i := len(a) - 1; i > 0; i-- {
			a[i], a[i-1] = a[i-1], a[i]
		}
	}
}

func rotateFaster(arr *[]int64, rotations int) {
	a := *arr

	for i := 0; i < rotations; i++ {
		// get last element
		last := a[len(a)-1]

		// shift right the elements by one position
		copy(a[1:], a[0:len(a)-1])

		// put the last element as first (circular right rotatation)
		a[0] = last
	}
}
