package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	fmt.Fscanln(os.Stdin, &input)

	output := input
	tmp := output

	for {
		output = reduce(output)
		if strings.Compare(output, tmp) == 0 {
			break
		}
		tmp = output
	}

	if len(output) == 0 {
		fmt.Println("Empty String")
	} else {
		fmt.Println(output)
	}
}

func reduce(str string) string {
	result := str
	for _, r := range str {
		var tmp string
		tmp += string(r)
		tmp += string(r)
		if strings.Contains(str, tmp) {
			result = strings.Replace(str, string(tmp), "", 1)
			break
		}
	}
	return result
}
