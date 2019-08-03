package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		s := scanner.Text()

		// separate the string by colons ':'
		parts := strings.FieldsFunc(s, func(r rune) bool { return r == ':' })

		//		fmt.Println(parts)

		// the first part is the hour
		hour, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		last := parts[len(parts)-1]

		// the last characters of the input are the AM/PM indicator
		ampm := last[len(last)-2:]

		//		fmt.Printf("%v\n", hour)
		//		fmt.Printf("%v\n", ampm)

		var h int64

		if ampm == "PM" {
			if hour < 12 {
				h = hour + 12
			} else {
				h = hour
			}
		} else if ampm == "AM" {
			if hour == 12 {
				h = 0
			} else {
				h = hour
			}
		} else {
			fmt.Println("error, invalid time format")
			os.Exit(1)
		}

		var newtime string
		newtime = fmt.Sprintf("%02d:%v:%v", h, parts[1], parts[2][:len(parts[2])-2])

		fmt.Println(newtime)
	}
}
