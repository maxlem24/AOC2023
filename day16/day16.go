package main

import (
	"fmt"
	"os"
)

type Direction int

const (
	Top   Direction = 0
	Left  Direction = 1
	Down  Direction = 2
	Right Direction = 3
)

type node struct {
	x         int
	y         int
	direction Direction
	down      *node
}

type stack struct {
	head *node
}

func (s *stack) isEmpty() bool {
	return s.head == nil
}

func (s *stack) pop() {
	tmp := s.head
	s.head = tmp.down
}

func (s *stack) top() (int, int, Direction) {
	return s.head.x, s.head.y, s.head.direction
}

func (s *stack) push(x int, y int, dir Direction) {
	new := &node{x, y, dir, s.head}
	s.head = new
}

func newStack() stack {
	return stack{nil}
}

func openFile() [][]string {
	buffer, err := os.ReadFile("day16/test.txt")
	if err != nil {
		panic(err)
	}
	contraption := [][]string{}
	line := []string{}
	for i := 0; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line = append(line, string(buffer[i]))
		} else {
			contraption = append(contraption, line)
			line = []string{}
		}
	}
	return contraption
}

func contains(memory []node, x int, y int, dir Direction) bool {
	for _, tile := range memory {
		if tile.x == x && tile.y == y && tile.direction == dir {
			return true
		}
	}
	return false
}

func part1(contraption [][]string) {
	energized := make([][]string, len(contraption))
	for i := 0; i < len(contraption); i++ {
		energized[i] = make([]string, len(contraption[i]))
	}
	beam := newStack()
	beam.push(0, 0, Right)
	memory := []node{}
	for !beam.isEmpty() {
		x, y, dir := beam.top()
		energized[y][x] = "#"
		beam.pop()
		if contains(memory, x, y, dir) {
			continue
		}
		memory = append(memory, node{x, y, dir, nil})
		switch contraption[y][x] {
		case "/":
			switch dir {
			case Top:
				if x+1 < len(contraption[y]) {
					beam.push(x+1, y, Right)
				}
			case Right:
				if y-1 >= 0 {
					beam.push(x, y-1, Top)
				}
			case Down:
				if x-1 >= 0 {
					beam.push(x-1, y, Left)
				}
			case Left:
				if y+1 < len(contraption) {
					beam.push(x, y+1, Down)
				}
			}
		case "\\":
			switch dir {
			case Top:
				if x-1 >= 0 {
					beam.push(x-1, y, Left)
				}
			case Right:
				if y+1 < len(contraption) {
					beam.push(x, y+1, Down)
				}
			case Down:
				if x+1 < len(contraption[y]) {
					beam.push(x+1, y, Right)
				}
			case Left:
				if y-1 >= 0 {
					beam.push(x, y-1, Top)
				}
			}
		case "|":
			switch dir {
			case Right, Left:
				if y+1 < len(contraption) {
					beam.push(x, y+1, Down)
				}
				if y-1 >= 0 {
					beam.push(x, y-1, Top)
				}
			case Top:
				if y-1 >= 0 {
					beam.push(x, y-1, Top)
				}
			case Down:
				if y+1 < len(contraption) {
					beam.push(x, y+1, Down)
				}
			}
		case "-":
			switch dir {
			case Top, Down:
				if x-1 >= 0 {
					beam.push(x-1, y, Left)
				}
				if x+1 < len(contraption[y]) {
					beam.push(x+1, y, Right)
				}
			case Left:
				if x-1 >= 0 {
					beam.push(x-1, y, Left)
				}
			case Right:
				if x+1 < len(contraption[y]) {
					beam.push(x+1, y, Right)
				}
			}
		case ".":
			switch dir {
			case Top:
				if y-1 >= 0 {
					beam.push(x, y-1, Top)
				}
			case Right:
				if x+1 < len(contraption[y]) {
					beam.push(x+1, y, Right)
				}
			case Down:
				if y+1 < len(contraption) {
					beam.push(x, y+1, Down)
				}
			case Left:
				if x-1 >= 0 {
					beam.push(x-1, y, Left)
				}
			}
		}
	}
	sum := 0
	for _, line := range energized {
		for _, tile := range line {
			if tile == "#" {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	contraption := openFile()
	part1(contraption)
}
