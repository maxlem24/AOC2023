package main

import (
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func openFile() []point {
	buffer, err := os.ReadFile("day11/day11.txt")
	if err != nil {
		panic(err)
	}
	image := [][]string{}
	galaxies := []point{}
	x, y := 0, 0
	line := []string{}
	for i := 0; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line = append(line, string(buffer[i]))
			if string(buffer[i]) == "#" {
				galaxies = append(galaxies, point{x, y})
			}
			x++
		} else {
			x = 0
			image = append(image, line)
			line = []string{}
			y++
		}
	}
	addedRow := 0
	for y = 0; y < len(image); y++ {
		for x = 0; x < len(image[0]); x++ {
			if image[y][x] == "#" {
				break
			}
		}
		if x == len(image[y]) {
			for i, _ := range galaxies {
				if galaxies[i].y > y+addedRow {
					galaxies[i].y++
				}
			}
			addedRow++
		}
	}
	addedColumn := 0
	for x = 0; x < len(image[0]); x++ {
		for y = 0; y < len(image); y++ {
			if image[y][x] == "#" {
				break
			}
		}
		if y == len(image) {
			for i, _ := range galaxies {
				if galaxies[i].x > x+addedColumn {
					galaxies[i].x++
				}
			}
			addedColumn++
		}
	}
	return galaxies
}

func part1(galaxies []point) {
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			length := abs(galaxies[i].x-galaxies[j].x) + abs(galaxies[i].y-galaxies[j].y)
			sum += length
		}
	}
	fmt.Println(sum)
}

func main() {
	galaxies := openFile()
	part1(galaxies)
}
