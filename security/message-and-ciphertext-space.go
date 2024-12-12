package main

import (
	"fmt"
	"strings"
)

func main() {
	var msg string
	fmt.Scanln(&msg)

	ciphertext := encrypt(msg)

	fmt.Println(ciphertext)
}

// encrypt shifts the given plaintext by incrementing char code +1
// plaintext is digits only. treating them as ascii codes means min
// number is 0 = 48 (decimal) and max number is 9 = 57 (decimal)
func encrypt(plaintext string) string {
	var buf strings.Builder
	for _, ch := range plaintext {
		if ch < 57 {
			ch = ch + 1
		} else {
			ch = 48
		}
		buf.WriteRune(ch)
	}

	return buf.String()
}
