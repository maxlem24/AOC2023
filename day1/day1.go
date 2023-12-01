package main

import (
	"fmt"
	"os"
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

func main() {
	list := openFile()
	part1(list)
}
