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

func sortRanges(ranges [][2]int64) [][2]int64 {
	sortedRanges := make([][2]int64, len(ranges))
	copy(sortedRanges, ranges)
	for i := 0; i < len(ranges); i++ {
		min := i
		for j := i + 1; j < len(ranges); j++ {
			if sortedRanges[j][0] < sortedRanges[min][0] {
				min = j
			}
		}
		temp := sortedRanges[i]
		sortedRanges[i] = sortedRanges[min]
		sortedRanges[min] = temp
	}
	return sortedRanges

}

func mergeRanges(ranges [][2]int64) [][2]int64 {
	sortedRanges := sortRanges(ranges)
	mergedRanges := make([][2]int64, 0)

	mergedRanges = append(mergedRanges, sortedRanges[0])
	indexMerge := 0
	for indexSort := 1; indexSort < len(sortedRanges); indexSort++ {
		min := sortedRanges[indexSort][0]
		if min >= mergedRanges[indexMerge][0] && min-1 <= mergedRanges[indexMerge][1] {
			mergedRanges[indexMerge][1] = sortedRanges[indexSort][1] 
		} else {
			mergedRanges = append(mergedRanges, sortedRanges[indexSort])
			indexMerge++
		}

	}
	return mergedRanges
}

func part2(listSeeds []int64, maps [][][3]int64) {
	ranges := [][2]int64{}
	for i := 0; i < len(listSeeds); i += 2 {
		start := listSeeds[i]
		end := start + listSeeds[i+1] - 1
		ranges = append(ranges, [2]int64{start, end})
	}
	ranges = mergeRanges(ranges)
	for _, categorie := range maps {
		newRange := [][2]int64{}
		for _, rule := range categorie {
			convert, start, interval := rule[0], rule[1], rule[2]
			for i := 0; i < len(ranges); i++ {
				begin, end := ranges[i][0], ranges[i][1]
				gap := convert - start
				if begin >= start && end < start+interval { // All the interval
					newRange = append(newRange, [2]int64{begin + gap, end + gap})
					ranges = append(ranges[:i], ranges[i+1:]...)
					i--
				} else if begin < start && end < start+interval && end >= start { // left side
					newRange = append(newRange, [2]int64{start + gap, end + gap})
					ranges = append(ranges[:i], ranges[i+1:]...)
					ranges = append(ranges, [2]int64{begin, start - 1})
					i--
				} else if begin >= start && begin < start+interval && end >= start+interval { // right side
					newRange = append(newRange, [2]int64{begin + gap, start + interval - 1 + gap})
					ranges = append(ranges[:i], ranges[i+1:]...)
					ranges = append(ranges, [2]int64{start + interval, end})
					i--
				} else if begin < start && end >= start+interval { // outside
					ranges = append(ranges[:i], ranges[i+1:]...)
					newRange = append(newRange, [2]int64{start + gap, start + interval - 1 + gap})
					ranges = append(ranges, [2]int64{begin, start - 1})
					ranges = append(ranges, [2]int64{start + interval, end})
					i--
				}
			}
		}
		ranges = append(newRange, ranges...)
		ranges = mergeRanges(ranges)
	}
	fmt.Println(ranges[0][0])
}

func main() {
	listSeeds, maps := openFile()
	part1(listSeeds, maps)
	part2(listSeeds, maps)
}
