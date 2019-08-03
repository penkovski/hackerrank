package main

import "fmt"

func main() {
	var arrSize int
	var arrSum int

	fmt.Scan(&arrSize)

	for i := 0; i < arrSize; i++ {
		var num int
		fmt.Scan(&num)
		arrSum += num
	}

	fmt.Println(arrSum)
}
