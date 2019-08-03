package main

import (
	"fmt"
	"math"
)

func main() {
	var input string
	fmt.Scanln(&input)

	len := len(input)
	rows := int(math.Floor(math.Sqrt(float64(len))))
	cols := int(math.Ceil(math.Sqrt(float64(len))))

	if rows * cols < len {
		if rows > cols {
			cols += 1
		} else {
			rows += 1
		}
	}
    
	var pos int
	var grid [][]rune = make([][]rune, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]rune, cols)
		for j := 0; j < cols; j++ {
			if pos < len {
				grid[i][j] = []rune(input)[pos]
				pos++				
			}
		}
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if grid[j][i] > 0 {
				fmt.Print(string(grid[j][i]))				
			}
		}
		fmt.Print(" ")
	}
}