package main

import (
	"fmt"
	"strings"
)

func main() {
	var length int
	var binary string
	fmt.Scanln(&length)
	fmt.Scanln(&binary)

	if length < 3 {
		fmt.Println("0")
		return
	}

    fmt.Println(beautify2(binary))
}

func beautify2(s string) int {
	strlen := len(s)

	tmp := strings.Replace(s, "010", "", -1)

	tmplen := len(tmp)

	res := (strlen - tmplen) / 3

	//fmt.Printf("strlen = %d, tmplen = %d, res = %v\n", strlen, tmplen, res)

	//fmt.Printf("(%d - %d) / 3 = %v\n", strlen, tmplen, res)

	return res
}

func beautify(s string) (string, bool) {
	var result string
	pos := strings.Index(s, "010")
	if pos >= 0 {
		result = s[:pos+1] + "0" + s[pos+2:]
		return result, false
	}

	return s, true
}

