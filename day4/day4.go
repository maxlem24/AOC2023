package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"fmt"
)

func power(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * power(x, y-1)
}

func openFile() [][2][]int {
	ruleNumber := regexp.MustCompile("[0-9]+")
	buffer, err := os.ReadFile("day4/day4.txt")
	if err != nil {
		panic(err)
	}
	list := [][2][]int{}
	line := ""
	for i := 0; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line += string(buffer[i])
		} else {
			header := strings.Split(line,":")
			split := strings.Split(header[1], "|")
			winningString := ruleNumber.FindAllString(split[0], -1)
			chosenString := ruleNumber.FindAllString(split[1], -1)
			winningInt := make([]int, len(winningString))
			chosenInt := make([]int, len(chosenString))
			for index, win := range winningString {
				value, err := strconv.ParseInt(win, 10, 0)
				if err != nil {
					panic(err)
				}
				winningInt[index] = int(value)
			}
			for index, choice := range chosenString {
				value, err := strconv.ParseInt(choice, 10, 0)
				if err != nil {
					panic(err)
				}
				chosenInt[index] = int(value)
			}
			list = append(list, [2][]int{winningInt, chosenInt})
			line = ""
		}
	}
	return list
}

func part1(list [][2][]int) {
	sum := 0
	for _, card := range list {
		count := 0
		for _, win := range card[0] {
			for _, choice := range card[1] {
				if win == choice {
					count++
				}
			}
		}
		if count != 0 {
			sum += power(2, count-1)
		}
	}
	fmt.Println(sum)
}

func main() {
	list := openFile()
	part1(list)
}
