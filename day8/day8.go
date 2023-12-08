package main

import (
	"fmt"
	"os"
	"regexp"
)

func openFile() ([]int, map[string][2]string) {
	buffer, err := os.ReadFile("day8/day8.txt")
	nodeNames := regexp.MustCompile("\\w{3}")
	if err != nil {
		panic(err)
	}
	instructions := []int{}
	i := 0
	for string(buffer[i]) != "\n" {
		if string(buffer[i]) == "R" {
			instructions = append(instructions, 1)
		} else {
			instructions = append(instructions, 0)
		}
		i++
	}
	i++
	names := map[string][2]string{}
	line := ""
	for ; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line += string(buffer[i])
		} else {
			nodes := nodeNames.FindAllString(line, 3)
			if len(nodes) == 3 {
				names[nodes[0]] = [2]string{nodes[1], nodes[2]}
			}
			line = ""
		}
	}
	return instructions, names
}

func part1(instructions []int, names map[string][2]string) {
	i := 0
	for node := "AAA"; node != "ZZZ"; {
		side := instructions[i % len(instructions)]
		node = names[node][side]
		i++
	}
	fmt.Println(i)
}

func main() {
	instructions, names := openFile()
	part1(instructions, names)
}
