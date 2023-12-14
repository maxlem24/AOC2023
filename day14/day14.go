package main

import (
	"fmt"
	"os"
)

func openFile() [][]string {
	buffer, err := os.ReadFile("day14/day14.txt")
	if err != nil {
		panic(err)
	}
	platform := [][]string{}
	line := []string{}
	for i := 0; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line = append(line, string(buffer[i]))
		} else {
			platform = append(platform, line)
			line = []string{}
		}
	}
	return platform
}

func moveNorth(platform [][]string) {
	for i := 0; i < len(platform[0]); i++ {
		slideDestination := 0
		for j := 0; j < len(platform); j++ {
			switch platform[j][i] {
			case "#":
				slideDestination = j + 1
			case "O":
				platform[j][i] = "."
				platform[slideDestination][i] = "O"
				slideDestination++
			}
		}
	}

}

func part1(platform [][]string) {
	moveNorth(platform)
	sum := 0
	for j := 0; j < len(platform); j++ {
		for i := 0; i < len(platform[0]); i++ {
			if platform[j][i] == "O" {
				sum += len(platform) - j
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	platform := openFile()
	part1(platform)
}
