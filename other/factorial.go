package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int64
	fmt.Scan(&n)

	fmt.Println(factorial(n))
}

func factorial(n int64) int64 {
	if n < 3 {
		return n
	}

	total := big.Int{}

	for i := 2; i <= int(n); i++ {
		multi := big.Int{}
		if total.Int64() == 0 {
			multi.Mul(big.NewInt(int64(i)), big.NewInt(int64(i-1)))
		} else {
			multi.Mul(&total, big.NewInt(int64(i)))
		}
		total.Add(&total, &multi)
	}

	return total.Int64()
}
