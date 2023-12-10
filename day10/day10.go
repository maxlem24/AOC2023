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
		// for i := 0; i < yMax; i++ {
		// 	for j := 0; j < xMax; j++ {
		// 		if values[i][j] == -1 {
		// 			fmt.Print("  ")
		// 		} else {
		// 			fmt.Printf("%2d", values[i][j])
		// 		}

		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println("=============")
		// time.Sleep(time.Second)
	}
	fmt.Println(value - 1)

}

func main() {
	laby := openFile()
	part1(laby)
}
