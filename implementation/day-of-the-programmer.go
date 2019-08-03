// https://www.hackerrank.com/challenges/day-of-the-programmer
package main

import (
	"fmt"
	"time"
)

const DAY = time.Minute * 60 * 24

func main() {
	var year int
	fmt.Scan(&year)

	var day time.Time
	if year > 1918 {
		start := time.Date(year, 1, 1, 0, 0, 1, 0, time.UTC)
		day = start.Add(255 * DAY)
	} else if year == 1918 {
		start := time.Date(year, 2, 14, 0, 0, 1, 0, time.UTC)
		day = start.Add((256 - 32) * DAY)
	} else {
		start := time.Date(year, 1, 1, 0, 0, 1, 0, time.UTC)
		day = start.Add(255 * DAY)

		if year == 1700 || year == 1800 || year == 1900 {
			day = day.Add(-1 * DAY)
		}
	}

	fmt.Printf("%02d.%02d.%04d\n", day.Day(), day.Month(), day.Year())
}
