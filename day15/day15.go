package main

import (
	"fmt"
	"os"
)

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
		sum +=  hash(instruction)
	}
	fmt.Println(sum)
}

func main() {
	sequence := openFile()
	part1(sequence)
}
