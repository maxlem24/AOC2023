package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func openFile() [][]int {
	buffer, err := os.ReadFile("day9/day9.txt")
	if err != nil {
		panic(err)
	}
	report := [][]int{}
	line := ""
	for i := 0; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line += string(buffer[i])
		} else {
			history := []int{}
			for _, value := range strings.Split(line, " ") {
				data, err := strconv.ParseInt(value, 10, 0)
				if err != nil {
					panic(err)
				}
				history = append(history, int(data))
			}
			report = append(report, history)
			line = ""
		}
	}

	return report
}

func allZeros(line []int) bool {
	for i := 0; i < len(line); i++ {
		if line[i] != 0 {
			return false
		}
	}
	return true
}

func part1(report [][]int) {
	sum := 0
	for _, history := range report {
		pyramid := [][]int{history}
		last := history
		for !allZeros(last) {
			line := make([]int, len(last)-1)
			for i := 0; i < len(last)-1; i++ {
				line[i] = last[i+1] - last[i]
			}
			pyramid = append(pyramid, line)
			last = line
		}
		pyramid[len(pyramid)-1] = append(pyramid[len(pyramid)-1], 0)
		for i := len(pyramid) - 2; i >= 0; i-- {
			placeholder := pyramid[i][len(pyramid[i])-1] + pyramid[i+1][len(pyramid[i+1])-1]
			pyramid[i] = append(pyramid[i], placeholder)
		}
		sum += pyramid[0][len(pyramid[0])-1]
	}
	fmt.Println(sum)
}

func main() {
	report := openFile()
	part1(report)
}
