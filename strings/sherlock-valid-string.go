package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(valid(s))
}

func valid(s string) string {
	found := make(map[rune]int)
	for _, c := range s {
		found[c]++
	}

	var seen int
	var seen2 int
	deviations := 0
	for _, freq := range found {
		if seen == 0 {
			seen = freq
			continue
		}

		if freq != seen {
			if seen2 == 0 {
				deviations++
				seen2 = freq
				continue
			}

			if abs(seen-seen2) > 1 {
				return "NO"
			}

			if freq == seen2 {
				deviations++
				continue
			} else {
				return "NO"
			}
		}
	}

	if deviations > 1 {
		return "NO"
	}

	return "YES"
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
