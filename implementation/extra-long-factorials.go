package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scanln(&n)

	if n <= 2 {
		fmt.Println(n)
		return
	}

	factorial := new(big.Int)
	factorial.SetInt64(1)
	for i := 2; i <= n; i++ {
		bigI := new(big.Int)
		bigI.SetInt64(int64(i))
		factorial = factorial.Mul(factorial, bigI)
	}

	fmt.Println(factorial)
}
