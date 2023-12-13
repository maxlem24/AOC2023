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

func mergeMaps(m1 map[string]int, m2 map[string]int) map[string]int {
	merged := make(map[string]int)
	for key, value := range m1 {
		merged[key] = value
	}
	for key, value := range m2 {
		merged[key] = value
	}
	return merged
}

func recordValid(record string, values []int) bool {
	currentvalue := 0
	i := 0
	if sumArray(values)+len(values)-1 > len(record) {
		return false
	}
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
			if i == len(record) {
				if currentvalue != len(values) {
					return false
				}
				return true
			}
			if string(record[i]) == "#" {
				return false
			}
		}
		i++
	}
	if currentvalue != len(values) {
		return false
	}
	return true
}

func backtracking(record string, values []int, memory *map[string]int) int {
	ruleShard := regexp.MustCompile("#+")
	ruleQuestion := regexp.MustCompile("\\?")
	value := 0
	index := ruleQuestion.FindStringIndex(record)
	if index == nil {
		return 1
	}
	recordShard := replaceAtIndex(record, '#', index[0])
	if recordValid(recordShard, values) {
		offset := 0
		alreadyCheck := 0
		if index[0] !=0 && string(recordShard[index[0]-1]) == "." {
			offset = index[0]
			alreadyCheck = len(ruleShard.FindAllString(recordShard[:index[0]-1], -1))
		}
		key := recordShard[offset:]
		for _, length := range values[alreadyCheck:] {
			key += strconv.Itoa(length)
		}
		val, exist := (*memory)[key]
		if !exist {
			val = backtracking(recordShard[offset:], values[alreadyCheck:], memory)
			(*memory)[key] = val
		}
		value += val
	}
	recordDot := replaceAtIndex(record, '.', index[0])
	if recordValid(recordDot, values) {
		alreadyCheck := len(ruleShard.FindAllString(recordDot[:index[0]], -1))
		key := recordDot[index[0]+1:]
		for _, length := range values[alreadyCheck:] {
			key += strconv.Itoa(length)
		}
		val, exist := (*memory)[key]
		if !exist {
			val = backtracking(recordDot[index[0]+1:], values[alreadyCheck:], memory)
			(*memory)[key] = val
		}
		value += val
	}
	return value
}

func part1(records []string, values [][]int) {

	sum := 0
	if len(records) != len(values) {
		panic("Size Problem")
	}
	for i, record := range records {
		memory := map[string]int{}
		sum += backtracking(record, values[i], &memory)
	}
	fmt.Println(sum)
}

func part2(records []string, values [][]int) {
	sum := 0
	if len(records) != len(values) {
		panic("Size Problem")
	}
	for i, record := range records {
		fmt.Print(i, "/", len(records), "\r")
		newValue := values[i]
		newRecord := record
		for k := 0; k < 4; k++ {
			newRecord += "?" + record
			newValue = append(newValue, values[i]...)
		}
		memory := map[string]int{}
		sum += backtracking(newRecord, newValue, &memory)
	}
	fmt.Println(sum)
}

func main() {
	records, values := openFile()
	part1(records, values)
	part2(records, values)
}
