package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func openFile() [][2]int {
	ruleNumber := regexp.MustCompile("[0-9]+")
	buffer, err := os.ReadFile("day6/day6.txt")
	if err != nil {
		panic(err)
	}
	races := [][2]int{}
	line := ""
	i := 0
	for string(buffer[i]) != "\n" {
		line += string(buffer[i])
		i++
	}
	i++
	times := ruleNumber.FindAllString(line, -1)
	line = ""
	for string(buffer[i]) != "\n" {
		line += string(buffer[i])
		i++
	}
	distances := ruleNumber.FindAllString(line, -1)
	if len(distances) != len(times) {
		panic("Lecture !")
	}
	for i := 0; i < len(distances); i++ {
		time, err := strconv.ParseInt(times[i], 10, 0)
		if err != nil {
			panic(err)
		}
		distance, err := strconv.ParseInt(distances[i], 10, 0)
		if err != nil {
			panic(err)
		}
		races = append(races, [2]int{int(time), int(distance)})
	}
	return races
}

func openFile2() [][2]int {
	ruleNumber := regexp.MustCompile("[0-9]+")
	buffer, err := os.ReadFile("day6/day6.txt")
	if err != nil {
		panic(err)
	}
	races := [][2]int{}
	line := ""
	i := 0
	for string(buffer[i]) != "\n" {
		if string(buffer[i]) != " " {
			line += string(buffer[i])
		}
		i++
	}
	i++
	times := ruleNumber.FindAllString(line, -1)
	line = ""
	for string(buffer[i]) != "\n" {
		if string(buffer[i]) != " " {
			line += string(buffer[i])
		}
		i++
	}
	distances := ruleNumber.FindAllString(line, -1)
	if len(distances) != len(times) {
		panic("Lecture !")
	}
	for i := 0; i < len(distances); i++ {
		time, err := strconv.ParseInt(times[i], 10, 0)
		if err != nil {
			panic(err)
		}
		distance, err := strconv.ParseInt(distances[i], 10, 0)
		if err != nil {
			panic(err)
		}
		races = append(races, [2]int{int(time), int(distance)})
	}
	return races
}

func part1(races [][2]int) {
	product := 1
	for _, race := range races {
		count := 0
		time, distance := race[0], race[1]
		for t := 0; t < time; t++ {
			travel := t * (time - t)
			if travel > distance {
				count++
			}
		}
		product *= count
	}
	fmt.Println(product)
}

func part2(races [][2]int) {
	product := 1
	for _, race := range races {
		time, distance := race[0], race[1]
		tMin:= 0
		for tMin * (time - tMin) < distance {
			tMin++
		}
		tMax:= time
		for tMax * (time - tMax) < distance {
			tMax--
		}
		product *= (tMax-tMin+1)
	}
	fmt.Println(product)
}

func main() {
	races := openFile()
	part1(races)
	races2 := openFile2()
	part2(races2)
}
