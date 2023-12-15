package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type lens struct {
	name  string
	focal int
}

func openFile() []string {
	buffer, err := os.ReadFile("day15/day15.txt")
	if err != nil {
		panic(err)
	}
	sequence := []string{}
	instruction := ""
	for i := 0; string(buffer[i]) != "\n"; i++ {
		if string(buffer[i]) != "," {
			instruction += string(buffer[i])
		} else {
			sequence = append(sequence, instruction)
			instruction = ""
		}
	}
	sequence = append(sequence, instruction)
	return sequence
}

func hash(instruction string) int {
	currentValue := 0
	for i := 0; i < len(instruction); i++ {
		currentValue += int(instruction[i])
		currentValue *= 17
		currentValue = currentValue % 256
	}
	return currentValue
}

func part1(sequence []string) {
	sum := 0
	for _, instruction := range sequence {
		sum += hash(instruction)
	}
	fmt.Println(sum)
}

func removeLens(boxs [256][]lens, label string) [256][]lens {
	boxNumber := hash(label)
	for i, lensInTheBox := range boxs[boxNumber] {
		if lensInTheBox.name == label {
			boxs[boxNumber] = append(boxs[boxNumber][:i], boxs[boxNumber][i+1:]...)
			break
		}
	}
	return boxs
}

func addLens(boxs [256][]lens, label string, focal int) [256][]lens {
	boxNumber := hash(label)
	toAdd := true
	for i, lensInTheBox := range boxs[boxNumber] {
		if lensInTheBox.name == label {
			boxs[boxNumber][i] = lens{label, focal}
			toAdd = false
			break
		}
	}
	if toAdd {
		boxs[boxNumber] = append(boxs[boxNumber], lens{label, focal})
	}
	return boxs
}

func part2(sequence []string) {
	ruleNumber := regexp.MustCompile("\\d")
	ruleLabel := regexp.MustCompile("[a-zA-Z]+")
	sum := 0
	boxs := [256][]lens{}
	for _, instruction := range sequence {
		label := ruleLabel.FindString(instruction)
		focalString := ruleNumber.FindString(instruction)
		if focalString == "" {
			boxs = removeLens(boxs, label)
		} else {
			focal, err := strconv.ParseInt(focalString, 10, 0)
			if err != nil {
				panic(err)
			}
			boxs = addLens(boxs, label, int(focal))
		}
	}
	for i, box := range boxs {
		for j, lens := range box {
			sum += (i+1)*(j+1)*lens.focal
		}
	}
	fmt.Println(sum)
}

func main() {
	sequence := openFile()
	part1(sequence)
	part2(sequence)
}
