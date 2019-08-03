package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	input, _ := in.ReadString('\n')

	var alphabet = map[string]bool{
		"a": false,
		"b": false,
		"c": false,
		"d": false,
		"e": false,
		"f": false,
		"g": false,
		"h": false,
		"i": false,
		"j": false,
		"k": false,
		"l": false,
		"m": false,
		"n": false,
		"o": false,
		"p": false,
		"q": false,
		"r": false,
		"s": false,
		"t": false,
		"u": false,
		"v": false,
		"w": false,
		"x": false,
		"y": false,
		"z": false,
	}

	for _, r := range input {
		alphabet[string(unicode.ToLower(r))] = true
	}

	if pangram(alphabet) {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}

func pangram(runes map[string]bool) bool {
	for _, b := range runes {
		if !b {
			return false
		}
	}
	return true
}

