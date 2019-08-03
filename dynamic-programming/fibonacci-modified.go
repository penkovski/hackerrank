package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
    "math/big"
)

func main() {
    var t1, t2, n int64

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s := scanner.Text()
		args := strings.Fields(s)
		t1, _ = strconv.ParseInt(args[0], 10, 64)
		t2, _ = strconv.ParseInt(args[1], 10, 64)
		n, _ = strconv.ParseInt(args[2], 10, 64)
	}

	//fmt.Printf("t1 = %d, t2 = %d, n = %d\n", t1, t2, n)

	fib := fibm(t1, t2, n)
    fmt.Println(fib)
}

func fibm(t1, t2, n int64) *big.Int {
	tmp := big.NewInt(0)
	result := big.NewInt(0)

	var t1inc, t2inc *big.Int
	t1inc = big.NewInt(t1)
	t2inc = big.NewInt(t2)

	for i := 3; i <= int(n); i++ {
		tmp.Mul(t2inc, t2inc)
		result.Add(t1inc, tmp)
		t1inc.Set(t2inc)
		t2inc.Set(result)

		// NOTE this is what the arithmetic looks like without big numbers
		//		result = t1inc + t2inc*t2inc
		//		t1inc = t2inc
		//		t2inc = result
	}

	return result
}

