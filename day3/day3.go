package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type star struct {
	X     int
	Y     int
	Value []int
}

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
				right = len(element)
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

func addStar(galaxy []star, newStar star) []star {
	for i := 0; i < len(galaxy); i++ {
		currentStar := galaxy[i]
		if currentStar.X == newStar.X && currentStar.Y == newStar.Y {
			currentStar.Value = append(currentStar.Value, newStar.Value[0])
			galaxy[i] = currentStar
			return galaxy
		}
	}
	galaxy = append(galaxy, newStar)
	return galaxy
}

func part2(list []string) {
	sum := 0
	ruleNumber := regexp.MustCompile("[0-9]+")
	ruleStar := regexp.MustCompile("\\*")
	galaxy := make([]star, 0)
	for listIndex, element := range list {
		i := 0
		for i < len(element) {
			index := ruleNumber.FindStringIndex(element[i:])
			if index == nil {
				break
			}
			value := ruleNumber.FindString(element[i:])
			number, err := strconv.ParseInt(value, 10, 0)
			if err != nil {
				panic(err)
			}
			left := i + index[0] - 1
			right := i + index[1] + 1
			if left < 0 {
				left = 0
			}
			if right >= len(element) {
				right = len(element)
			}
			coord := ruleStar.FindStringIndex(element[left:right])
			if coord != nil {
				star := star{left + coord[0], listIndex, []int{int(number)}}
				galaxy = addStar(galaxy, star)
			}
			if listIndex != 0 {
				coord := ruleStar.FindStringIndex(list[listIndex-1][left:right])
				if coord != nil {
					star := star{left + coord[0], listIndex - 1, []int{int(number)}}
					galaxy = addStar(galaxy, star)
				}
			}
			if listIndex != len(list)-1 {
				coord := ruleStar.FindStringIndex(list[listIndex+1][left:right])
				if coord != nil {
					star := star{left + coord[0], listIndex + 1, []int{int(number)}}
					galaxy = addStar(galaxy, star)
				}
			}
			i = right
		}
	}
	for _, star := range galaxy {
		if len(star.Value) == 2 {
			sum += star.Value[0] * star.Value[1]
		}
	}
	fmt.Println(sum)
}

func main() {
	list := openFile()
	part1(list)
	part2(list)
}
