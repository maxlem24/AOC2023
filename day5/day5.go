package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func openFile() ([]int64, [][][3]int64) {
	ruleNumber := regexp.MustCompile("[0-9]+")
	buffer, err := os.ReadFile("day5/day5.txt")
	if err != nil {
		panic(err)
	}
	listSeeds := []int64{}
	seeds := ""
	i := 0
	for string(buffer[i]) != "\n" {
		seeds += string(buffer[i])
		i++
	}
	seeds = strings.Split(seeds, ":")[1]
	for _, number := range ruleNumber.FindAllString(seeds, -1) {
		seed, err := strconv.ParseInt(number, 10, 0)
		if err != nil {
			panic(err)
		}
		listSeeds = append(listSeeds, seed)
	}
	line := ""
	maps := make([][][3]int64, 0)
	categorie := make([][3]int64, 0)
	for ; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line += string(buffer[i])
		} else {
			if line == "" {
				continue
			}
			numbers := ruleNumber.FindAllString(line, -1)
			if numbers == nil {
				if len(categorie) != 0 {
					maps = append(maps, categorie)
					categorie = make([][3]int64, 0)
				}
			} else {
				rule := [3]int64{}
				for index, number := range numbers {
					value, err := strconv.ParseInt(number, 10, 0)
					if err != nil {
						panic(err)
					}
					rule[index] = value
				}
				categorie = append(categorie, rule)
			}
			line = ""
		}
	}
	if len(categorie) != 0 {
		maps = append(maps, categorie)
	}
	return listSeeds, maps
}

func part1(listSeeds []int64, maps [][][3]int64) {
	min := int64(math.MaxInt64)
	for _, seed := range listSeeds {
		finalValue := seed
		for _, categorie := range maps {
			for _, rule := range categorie {
				convert, start, interval := rule[0], rule[1], rule[2]
				if finalValue >= start && finalValue < start+interval {
					gap := finalValue - start
					finalValue = convert + gap
					break
				}
			}
		}
		if finalValue < min {
			min = finalValue
		}
	}
	fmt.Println(min)
}

func main() {
	listSeeds, maps := openFile()
	part1(listSeeds, maps)
}
