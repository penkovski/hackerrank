package main

import "fmt"

func main() {
	var s1, s2 string

	fmt.Scanln(&s1)
	fmt.Scanln(&s2)

	fmt.Println(makeAnagram(s1, s2))
}

func makeAnagram(s1 string, s2 string) int {
	var countDelete int
	s1Frequency := make(map[rune]int)
	s2Frequency := make(map[rune]int)

	for _, char := range s1 {
		s1Frequency[char]++
	}

	for _, char := range s2 {
		s2Frequency[char]++
	}

	for r, freq := range s1Frequency {
		countDelete += abs(freq - s2Frequency[r])
		delete(s1Frequency, r)
		delete(s2Frequency, r)
	}

	for r, freq := range s2Frequency {
		countDelete += abs(freq - s1Frequency[r])
	}

	return countDelete
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
