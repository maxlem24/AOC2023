package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Right Direction = 0
	Down  Direction = 1
	Left  Direction = 2
	Up    Direction = 3
)

type instruction struct {
	direction Direction
	length    int
	color     string
}

func (i *instruction) getValue() (Direction, int, string) {
	return i.direction, i.length, i.color
}

type node struct {
	x    int
	y    int
	next *node
}

type queue struct {
	head *node
	tail *node
}

func (q *queue) isEmpty() bool {
	return q.head == nil && q.tail == nil
}

func (q *queue) pop() {
	tmp := q.head
	q.head = tmp.next
	if q.head == nil {
		q.tail = nil
	}
}

func (q *queue) top() (int, int) {
	return q.head.x, q.head.y
}

func (q *queue) push(x int, y int) {
	new := &node{x, y, nil}
	if q.tail != nil {
		q.tail.next = new
	} else {
		q.head = new
	}
	q.tail = new
}

func newQueue() queue {
	return queue{nil, nil}
}

func openFile() ([]instruction, [2]int) {
	buffer, err := os.ReadFile("day18/day18.txt")
	if err != nil {
		panic(err)
	}
	digPlan := []instruction{}
	line := ""
	counters := [2]int{0, 0}
	for i := 0; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line += string(buffer[i])
		} else {
			splited := strings.Split(line, " ")
			direction := Up
			switch splited[0] {
			case "U":
				direction = Up
			case "R":
				direction = Right
			case "D":
				direction = Down
			case "L":
				direction = Left
			}
			length, err := strconv.ParseInt(splited[1], 10, 0)
			if err != nil {
				panic(err)
			}
			if int(direction) < 2 {
				counters[direction] += int(length)
			}
			color := splited[2][1 : len(splited[2])-1]
			digPlan = append(digPlan, instruction{direction, int(length), color})
			line = ""
		}
	}
	return digPlan, counters
}

func floodFilling(x int, y int, digArea [][]string) {
	q := newQueue()
	q.push(x, y)
	for !q.isEmpty() {
		x, y = q.top()
		q.pop()
		if digArea[y][x+1] == "." {
			digArea[y][x+1] = "#"
			q.push(x+1, y)
		}
		if digArea[y][x-1] == "." {
			digArea[y][x-1] = "#"
			q.push(x-1, y)
		}
		if digArea[y+1][x] == "." {
			digArea[y+1][x] = "#"
			q.push(x, y+1)
		}
		if digArea[y-1][x] == "." {
			digArea[y-1][x] = "#"
			q.push(x, y-1)
		}
	}
}

func part1(digPlan []instruction, counters [2]int) {
	digArea := make([][]string, 2*counters[1]+1)
	for i := 0; i < len(digArea); i++ {
		digArea[i] = make([]string, 2*counters[0]+1)
		for j := 0; j < len(digArea[i]); j++ {
			digArea[i][j] = "."
		}
	}
	x, y := counters[0], counters[1]
	for _, line := range digPlan {
		direction, length, _ := line.getValue()
		switch direction {
		case Up:
			for i := 0; i < length; i++ {
				digArea[y][x] = "#"
				y--
			}
		case Right:
			for i := 0; i < length; i++ {
				digArea[y][x] = "#"
				x++
			}
		case Down:
			for i := 0; i < length; i++ {
				digArea[y][x] = "#"
				y++
			}
		case Left:
			for i := 0; i < length; i++ {
				digArea[y][x] = "#"
				x--
			}
		}
	}
	floodFilling(counters[0]-1, counters[1]-1, digArea)
	sum := 0
	for i := 0; i < len(digArea); i++ {
		for j := 0; j < len(digArea[i]); j++ {
			if digArea[i][j] == "#" {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	digPlan, counters := openFile()
	part1(digPlan, counters)
}
