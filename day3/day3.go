package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "strings"
)

func openFile() []string {
	buffer, err := os.ReadFile("day3/day3.txt")
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
	sum := 0
	ruleNumber := regexp.MustCompile("[0-9]+")
	ruleSymbol := regexp.MustCompile("[^.\\w]")
	for listIndex, element := range list {
		i := 0
		for i < len(element) {
			index := ruleNumber.FindStringIndex(element[i:])
			if index == nil {
				break
			}
			value := ruleNumber.FindString(element[i:])
			left := i + index[0] - 1
			right := i + index[1] + 1
			if left < 0 {
				left = 0
			}
			if right >= len(element) {
				right =len(element)
			}
			add := false
			if listIndex != 0 {
				if ruleSymbol.FindString(list[listIndex-1][left:right]) != "" {
					add = true
				}
			}
			if ruleSymbol.FindString(element[left:right]) != "" {
				add = true
			}
			if listIndex != len(list)-1 {
				if ruleSymbol.FindString(list[listIndex+1][left:right]) != "" {
					add = true
				}
			}
			i = right
			if add {
				number, err := strconv.ParseInt(value, 10, 0)
				if err != nil {
					panic(err)
				}
				sum += int(number)
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	list := openFile()
	part1(list)

}
