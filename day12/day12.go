package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func openFile() ([]string, [][]int) {
	buffer, err := os.ReadFile("day12/day12.txt")
	ruleNumber := regexp.MustCompile("\\d+")
	if err != nil {
		panic(err)
	}
	records := []string{}
	values := [][]int{}
	line := ""
	for i := 0; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line += string(buffer[i])
		} else {
			parts := strings.Split(line, " ")
			records = append(records, parts[0])
			numbers := ruleNumber.FindAllString(parts[1], -1)
			lineValue := []int{}
			for _, number := range numbers {
				value, err := strconv.ParseInt(number, 10, 0)
				if err != nil {
					panic(err)
				}
				lineValue = append(lineValue, int(value))
			}
			values = append(values, lineValue)
			line = ""
		}
	}
	return records, values
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func sumArray(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func recordValid(record string, values []int) bool {
	currentvalue := 0
	i := 0
	for i < len(record) {
		if string(record[i]) == "?" {
			return true
		}
		if string(record[i]) == "#" {
			if currentvalue == len(values) {
				return false
			}
			stopIndex := i
			for ; i < stopIndex+values[currentvalue] && i < len(record); i++ {
				if string(record[i]) == "." {
					return false
				}
				if string(record[i]) == "?" {
					return true
				}
			}
			if i != stopIndex+values[currentvalue] {
				return false
			}
			currentvalue++
			if i == len(record){
				if currentvalue != len(values) {
					return false
				}
				return true
			}
			if string(record[i]) == "#" {
				return false
			}
		}
		i++;
	}
	if currentvalue != len(values) {
		return false
	}
	return true
}

func backtracking(record string, values []int, listIndex []int) int {
	value := 0
	if len(listIndex) == 0 {
		return 1
	}
	recordShard := replaceAtIndex(record, '#', listIndex[0])
	recordDot := replaceAtIndex(record, '.', listIndex[0])
	if recordValid(recordShard, values) {
		value += backtracking(recordShard, values, listIndex[1:])
	}
	if recordValid(recordDot, values) {
		value += backtracking(recordDot, values, listIndex[1:])
	}
	return value
}

func part1(records []string, values [][]int) {
	ruleQuestion := regexp.MustCompile("\\?")
	sum := 0
	if len(records) != len(values) {
		panic("Size Problem")
	}
	for i, record := range records {
		questionIndex := []int{}
		allIndex := ruleQuestion.FindAllStringIndex(record, -1)
		if allIndex == nil {
			continue
		}
		for _, index := range allIndex {
			questionIndex = append(questionIndex, index[0])
		}
		sum += backtracking(record, values[i], questionIndex)
	}
	fmt.Println(sum)
}

func main() {
	records, values := openFile()
	part1(records, values)
}
