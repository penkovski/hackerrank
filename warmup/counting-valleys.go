package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countingValleys function below.
func countingValleys(n int, s string) int {
	var currLevel int
	var valleys int
	var inValley bool
	for i := 0; i < n; i++ {
		step := s[i]
		if step == 'U' {
			currLevel += 1
			if currLevel == 0 {
				if inValley {
					valleys += 1
				}
				inValley = false
			}
			continue
		}

		currLevel -= 1
		if currLevel < 0 {
			inValley = true
		}
	}

	return valleys
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	s := readLine(reader)

	result := countingValleys(n, s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
