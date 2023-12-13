package main

import (
	"fmt"
	"os"
)

func openFile() [][][]string {
	buffer, err := os.ReadFile("day13/day13.txt")
	if err != nil {
		panic(err)
	}
	patterns := [][][]string{}
	line := []string{}
	pattern := [][]string{}
	for i := 0; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line = append(line, string(buffer[i]))
		} else {
			if len(line) == 0 {
				patterns = append(patterns, pattern)
				pattern = [][]string{}
			} else {
				pattern = append(pattern, line)
				line = []string{}
			}
		}
	}
	patterns = append(patterns, pattern)
	return patterns
}

func isSameRow(row1 []string, row2 []string) bool {
	for i := 0; i < len(row1); i++ {
		if row1[i] != row2[i] {
			return false
		}
	}
	return true
}

func isSameColumn(grid [][]string, index1 int, index2 int) bool {
	for i := 0; i < len(grid); i++ {
		if grid[i][index1] != grid[i][index2] {
			return false
		}
	}
	return true
}

func part1(patterns [][][]string) {
	sum := 0
	for _, pattern := range patterns {
		height := len(pattern)
		width := len(pattern[0])
		addValue := 0
		for middleRow := 0; middleRow < height-1; middleRow++ {
			row1, row2 := middleRow, middleRow+1
			for row1 >= 0 && row2 < height && isSameRow(pattern[row1], pattern[row2]) {
				row1--
				row2++
			}
			if row1 < 0 || row2 >= height {
				addValue = middleRow + 1
				break
			}
		}
		if addValue != 0 {
			sum += addValue * 100
		} else {
			for middleColumn := 0; middleColumn < width-1; middleColumn++ {
				column1, column2 := middleColumn, middleColumn+1
				for column1 >= 0 && column2 < width && isSameColumn(pattern, column1, column2) {
					column1--
					column2++
				}
				if column1 < 0 || column2 >= width {
					addValue = middleColumn + 1
					break
				}
			}
			sum += addValue
		}

	}
	fmt.Println(sum)
}

func main() {
	patterns := openFile()
	part1(patterns)
}
