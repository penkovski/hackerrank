package main

import (
	"fmt"
)

var words = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	21: "twenty one",
	22: "twenty two",
	23: "twenty three",
	24: "twenty four",
	25: "twenty five",
	26: "twenty six",
	27: "twenty seven",
	28: "twenty eight",
	29: "twenty nine",
	30: "thirty",
}

func main() {
	var h, m int
	fmt.Scanln(&h)
	fmt.Scanln(&m)

	if m == 0 {
		fmt.Println(words[h] + " o' clock")
	} else if m == 30 {
		fmt.Println("half past " + words[h])
	} else if m%15 == 0 {
		if m > 30 {
			fmt.Println("quarter to " + words[h+1])
		} else {
			fmt.Println("quarter past " + words[h])
		}
	} else {
		if m == 1 {
			fmt.Println(words[m] + " minute past " + words[h])
		} else if m == 59 {
			fmt.Println(words[(60-m)] + " minute to " + words[h+1])
		} else if m > 30 {
			fmt.Println(words[(60-m)] + " minutes to " + words[h+1])
		} else {
			fmt.Println(words[m] + " minutes past " + words[h])
		}
	}
}
