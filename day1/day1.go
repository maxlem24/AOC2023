package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func openFile() []string {
	buffer, err := os.ReadFile("day1/day1.txt")
	if err != nil {
		panic(err)
	}
	list := []string{}
	line := ""
	for i := 0; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line += string(buffer[i])
		} else {
			list = append(list, line)
			line = ""
		}
	}
	return list
}

func part1(list []string) {
	sum := int64(0)
	for _, element := range list {
		i := 0
		firstDigit, err := strconv.ParseInt(string(element[i]), 10, 0)
		for ; err != nil; firstDigit, err = strconv.ParseInt(string(element[i]), 10, 0) {
			i++
		}
		i = len(element) - 1
		secondDigit, err := strconv.ParseInt(string(element[i]), 10, 0)
		for ; err != nil; secondDigit, err = strconv.ParseInt(string(element[i]), 10, 0) {
			i--
		}
		sum += 10*firstDigit + secondDigit
	}
	fmt.Println(sum)
}

func part2(list []string) {
	sum := 0
	mapValue := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	ruleRegex := ""
	for value := range mapValue {
		ruleRegex += string(value) + "|"
	}
	ruleRegex = ruleRegex[:len(ruleRegex)-1]
	fmt.Println(ruleRegex)
	for _, element := range list {
		firstValue := regexp.MustCompile(ruleRegex).FindString(element)
		firstDigit := mapValue[firstValue]
		secondDigit := 0
		for i := len(element); i >= 0; i-- {
			secondValue := regexp.MustCompile(ruleRegex).FindString(element[i:])
			if secondValue != "" {
				secondDigit = mapValue[secondValue]
				break
			}
		}
		sum += 10*firstDigit + secondDigit
	}
	fmt.Println(sum)
}

func main() {
	list := openFile()
	part1(list)
	part2(list)
}
