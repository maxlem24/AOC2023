package main

import (
	"fmt"
	"os"
)

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

func openFile() [][]string {
	buffer, err := os.ReadFile("day10/day10.txt")
	if err != nil {
		panic(err)
	}
	laby := [][]string{}
	line := []string{}
	for i := 0; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line = append(line, string(buffer[i]))
		} else {
			laby = append(laby, line)
			line = []string{}
		}
	}
	return laby
}

func part1(laby [][]string) {
	xMax := len(laby[0])
	yMax := len(laby)
	values := make([][]int, yMax)
	for i := 0; i < yMax; i++ {
		values[i] = make([]int, xMax)
		for j := 0; j < xMax; j++ {
			values[i][j] = -1
		}
	}
	x, y := 0, 0
	for ; y < yMax; y++ {
		for x = 0; x < xMax; x++ {
			if laby[y][x] == "S" {
				break
			}
		}
		if x != xMax {
			break
		}
	}
	q := newQueue()
	values[y][x] = 0
	if x+1 < xMax && (laby[y][x+1] == "-" || laby[y][x+1] == "7" || laby[y][x+1] == "J") {
		q.push(x+1, y)
		values[y][x+1] = 1
	}
	if x-1 >= 0 && (laby[y][x-1] == "-" || laby[y][x-1] == "F" || laby[y][x-1] == "L") {
		q.push(x-1, y)
		values[y][x-1] = 1
	}
	if y-1 >= 0 && (laby[y-1][x] == "|" || laby[y-1][x] == "F" || laby[y-1][x] == "7") {
		q.push(x, y-1)
		values[y-1][x] = 1
	}
	if y+1 < yMax && (laby[y+1][x] == "|" || laby[y+1][x] == "J" || laby[y+1][x] == "L") {
		q.push(x, y+1)
		values[y+1][x] = 1
	}
	value := 0
	for !q.isEmpty() {
		x, y = q.top()
		value = values[y][x]
		value++
		q.pop()
		switch laby[y][x] {
		case "|":
			if y+1 < yMax && values[y+1][x] == -1 {
				q.push(x, y+1)
				values[y+1][x] = value
			}
			if y-1 >= 0 && values[y-1][x] == -1 {
				q.push(x, y-1)
				values[y-1][x] = value
			}
		case "-":
			if x-1 >= 0 && values[y][x-1] == -1 {
				q.push(x-1, y)
				values[y][x-1] = value
			}
			if x+1 < xMax && values[y][x+1] == -1 {
				q.push(x+1, y)
				values[y][x+1] = value
			}
		case "F":
			if y+1 < yMax && values[y+1][x] == -1 {
				q.push(x, y+1)
				values[y+1][x] = value
			}
			if x+1 < xMax && values[y][x+1] == -1 {
				q.push(x+1, y)
				values[y][x+1] = value
			}
		case "7":
			if y+1 < yMax && values[y+1][x] == -1 {
				q.push(x, y+1)
				values[y+1][x] = value
			}
			if x-1 < xMax && values[y][x-1] == -1 {
				q.push(x-1, y)
				values[y][x-1] = value
			}
		case "L":
			if y-1 >= 0 && values[y-1][x] == -1 {
				q.push(x, y-1)
				values[y-1][x] = value
			}
			if x+1 >= 0 && values[y][x+1] == -1 {
				q.push(x+1, y)
				values[y][x+1] = value
			}
		case "J":
			if y-1 >= 0 && values[y-1][x] == -1 {
				q.push(x, y-1)
				values[y-1][x] = value
			}
			if x-1 >= 0 && values[y][x-1] == -1 {
				q.push(x-1, y)
				values[y][x-1] = value
			}
		}

	}
	fmt.Println(value - 1)
}

func part2(laby [][]string) {
	xMax := len(laby[0])
	yMax := len(laby)
	values := make([][]int, yMax)
	for i := 0; i < yMax; i++ {
		values[i] = make([]int, xMax)
		for j := 0; j < xMax; j++ {
			values[i][j] = -1
		}
	}
	xS, yS := 0, 0
	for ; yS < yMax; yS++ {
		for xS = 0; xS < xMax; xS++ {
			if laby[yS][xS] == "S" {
				break
			}
		}
		if xS != xMax {
			break
		}
	}
	q := newQueue()
	values[yS][xS] = 1
	if xS+1 < xMax && (laby[yS][xS+1] == "-" || laby[yS][xS+1] == "7" || laby[yS][xS+1] == "J") {
		q.push(xS+1, yS)
		values[yS][xS+1] = 1
	} else if xS-1 >= 0 && (laby[yS][xS-1] == "-" || laby[yS][xS-1] == "F" || laby[yS][xS-1] == "L") {
		q.push(xS-1, yS)
		values[yS][xS-1] = 1
	} else if yS-1 >= 0 && (laby[yS-1][xS] == "|" || laby[yS-1][xS] == "F" || laby[yS-1][xS] == "7") {
		q.push(xS, yS-1)
		values[yS-1][xS] = 1
	} else if yS+1 < yMax && (laby[yS+1][xS] == "|" || laby[yS+1][xS] == "J" || laby[yS+1][xS] == "L") {
		q.push(xS, yS+1)
		values[yS+1][xS] = 1
	}
	polygon := [][2]int{}
	for !q.isEmpty() {
		x, y := q.top()
		q.pop()
		switch laby[y][x] {
		case "|":
			if y+1 < yMax && values[y+1][x] == -1 {
				q.push(x, y+1)
				values[y+1][x] = 1
			}
			if y-1 >= 0 && values[y-1][x] == -1 {
				q.push(x, y-1)
				values[y-1][x] = 1
			}
		case "-":
			if x-1 >= 0 && values[y][x-1] == -1 {
				q.push(x-1, y)
				values[y][x-1] = 1
			}
			if x+1 < xMax && values[y][x+1] == -1 {
				q.push(x+1, y)
				values[y][x+1] = 1
			}
		case "F":
			polygon = append(polygon, [2]int{x, y})
			if y+1 < yMax && values[y+1][x] == -1 {
				q.push(x, y+1)
				values[y+1][x] = 1

			}
			if x+1 < xMax && values[y][x+1] == -1 {
				q.push(x+1, y)
				values[y][x+1] = 1
			}
		case "7":
			polygon = append(polygon, [2]int{x, y})
			if y+1 < yMax && values[y+1][x] == -1 {
				q.push(x, y+1)
				values[y+1][x] = 1
			}
			if x-1 < xMax && values[y][x-1] == -1 {
				q.push(x-1, y)
				values[y][x-1] = 1
			}
		case "L":
			polygon = append(polygon, [2]int{x, y})
			if y-1 >= 0 && values[y-1][x] == -1 {
				q.push(x, y-1)
				values[y-1][x] = 1
			}
			if x+1 >= 0 && values[y][x+1] == -1 {
				q.push(x+1, y)
				values[y][x+1] = 1
			}
		case "J":
			polygon = append(polygon, [2]int{x, y})
			if y-1 >= 0 && values[y-1][x] == -1 {
				q.push(x, y-1)
				values[y-1][x] = 1
			}
			if x-1 >= 0 && values[y][x-1] == -1 {
				q.push(x-1, y)
				values[y][x-1] = 1
			}
		}
	}
	begin, end := polygon[0], polygon[len(polygon)-1]
	if begin[0] != end[0] && begin[1] != end[1] {
		polygon = append(polygon, [2]int{xS, yS})
	}
	polygon = append(polygon, polygon[0])
	totalIn := 0
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			if values[y][x] == -1 {
				if crossingNumber(polygon,x,y) {
					totalIn++
				}
			}
		}
	}
	fmt.Println(totalIn)
}

func crossingNumber(polygon [][2]int, x int, y int) bool {
	cn := 0
	for i := 0; i < len(polygon)-1; i++ {
		a, b := polygon[i], polygon[i+1]
		if (a[1] <= y && y < b[1]) || (b[1] <= y && y < a[1]) {
			abX, abY := b[0]-a[0], b[1]-a[1]
			if abY == 0 {
				break
			}
			ax, ay := x-a[0], y-a[1]
			t := ay / abY
			if ax < t*abX {
				cn++
			}
		}
	}
	return cn%2 == 1
}

func main() {
	laby := openFile()
	part1(laby)
	part2(laby)
}
