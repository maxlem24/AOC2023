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

func moveNorth(platform [][]string) [][]string {
	platformCopy := make([][]string, len(platform))
	for i := 0; i < len(platform); i++ {
		platformCopy[i] = make([]string, len(platform[i]))
	}
	for i := 0; i < len(platformCopy[0]); i++ {
		slideDestination := 0
		for j := 0; j < len(platformCopy); j++ {
			switch platform[j][i] {
			case ".":
				platformCopy[j][i] = "."
			case "#":
				slideDestination = j + 1
				platformCopy[j][i] = "#"
			case "O":
				platformCopy[j][i] = "."
				platformCopy[slideDestination][i] = "O"
				slideDestination++
			}
		}
	}
	return platformCopy
}

func moveWest(platform [][]string) [][]string {
	platformCopy := make([][]string, len(platform))
	for i := 0; i < len(platform); i++ {
		platformCopy[i] = make([]string, len(platform[i]))
	}
	for j := 0; j < len(platformCopy); j++ {
		slideDestination := 0
		for i := 0; i < len(platformCopy[0]); i++ {
			switch platform[j][i] {
			case ".":
				platformCopy[j][i] = "."
			case "#":
				slideDestination = i + 1
				platformCopy[j][i] = "#"
			case "O":
				platformCopy[j][i] = "."
				platformCopy[j][slideDestination] = "O"
				slideDestination++
			}
		}
	}
	return platformCopy
}

func moveSouth(platform [][]string) [][]string {
	platformCopy := make([][]string, len(platform))
	for i := 0; i < len(platform); i++ {
		platformCopy[i] = make([]string, len(platform[i]))
	}
	for i := 0; i < len(platformCopy[0]); i++ {
		slideDestination := len(platformCopy) - 1
		for j := len(platformCopy) - 1; j >= 0; j-- {
			switch platform[j][i] {
			case ".":
				platformCopy[j][i] = "."
			case "#":
				slideDestination = j - 1
				platformCopy[j][i] = "#"
			case "O":
				platformCopy[j][i] = "."
				platformCopy[slideDestination][i] = "O"
				slideDestination--
			}
		}
	}
	return platformCopy
}

func moveEast(platform [][]string) [][]string {
	platformCopy := make([][]string, len(platform))
	for i := 0; i < len(platform); i++ {
		platformCopy[i] = make([]string, len(platform[i]))
	}
	for j := 0; j < len(platformCopy); j++ {
		slideDestination := len(platformCopy[0]) - 1
		for i := len(platformCopy[0]) - 1; i >= 0; i-- {
			switch platform[j][i] {
			case ".":
				platformCopy[j][i] = "."
			case "#":
				slideDestination = i - 1
				platformCopy[j][i] = "#"
			case "O":
				platformCopy[j][i] = "."
				platformCopy[j][slideDestination] = "O"
				slideDestination--
			}
		}
	}
	return platformCopy
}

func part1(platform [][]string) {
	platform = moveNorth(platform)
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

func arePlatformsEquals(p1 [][]string, p2 [][]string) bool {
	for j := 0; j < len(p1); j++ {
		for i := 0; i < len(p1[0]); i++ {
			if p1[j][i] != p2[j][i] {
				return false
			}
		}
	}
	return true
}

func part2(platform [][]string) {
	queuePlatform := [][][]string{platform}
	sizeCycle := 0
	for i := 1; i <= 1000000000 && sizeCycle == 0; i++ {
		platform = moveNorth(platform)
		platform = moveWest(platform)
		platform = moveSouth(platform)
		platform = moveEast(platform)
		for j := 0; j < len(queuePlatform); j++ {
			if arePlatformsEquals(platform, queuePlatform[j]) {
				sizeCycle = i - j
				remain := (1000000000 - i) % sizeCycle
				platform = queuePlatform[j+remain]
				break
			}
		}
		queuePlatform = append(queuePlatform, platform)
		fmt.Print(i, "/1000000000\r")
	}
	sum := 0
	for j := 0; j < len(platform); j++ {
		for i := 0; i < len(platform[0]); i++ {
			if platform[j][i] == "O" {
				sum += len(platform) - j
			}
		}
	}
	fmt.Println()
	fmt.Println(sum)
}

func main() {
	platform := openFile()
	part1(platform)
	part2(platform)
}
