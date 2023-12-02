package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const RED int = 12
const GREEN int = 13
const BLUE int = 14

func openFile() []string {
	buffer, err := os.ReadFile("day2/day2.txt")
	if err != nil {
		panic(err)
	}
	list := []string{}
	line := ""
	for i := 0; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line += string(buffer[i])
		} else {
			list = append(list, line)
			line = ""
		}
	}
	return list
}

func part1(list []string) {
	sum := 0
	for _, element := range list {
		game := strings.Split(element, ":")
		id,_ := strconv.ParseInt(regexp.MustCompile("[0-9]+").FindString(game[0]),10,0)
		sets := strings.Split(game[1],";")
		mapColor := map[string]int{"red":0,"green":1,"blue":2}
		maxCubes := [3]int{RED,GREEN,BLUE}
		valid := true
		for _, set := range sets {
			cubes := regexp.MustCompile("[0-9]+|red|green|blue").FindAllString(set,-1)
			for i := 0; i < len(cubes); i+=2 {
				value,_ := strconv.ParseInt(cubes[i],10,0)
				maxValue := maxCubes[mapColor[cubes[i+1]]]
				if int(value)>maxValue {
					valid = false
					break
				}
			}
			if !valid{
				break
			}
		}
		if valid {
			sum+= int(id)
		}
	}
	fmt.Println(sum)
}

func main() {
	list := openFile()
	part1(list)
}