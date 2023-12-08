package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func openFile() [][2]string {
	buffer, err := os.ReadFile("day7/day7.txt")
	if err != nil {
		panic(err)
	}
	cardsBid := [][2]string{}
	line := ""
	for i := 0; i < len(buffer); i++ {
		if string(buffer[i]) != "\n" {
			line += string(buffer[i])
		} else {
			parts := strings.Split(line, " ")
			cardsBid = append(cardsBid, [2]string{parts[0], parts[1]})
			line = ""
		}
	}
	return cardsBid
}

func order(hand1 string, hand2 string) bool { // true if card 1 > card 2
	cards1 := map[string]int{}
	for i := 0; i < len(hand1); i++ {

		if cards1[string(hand1[i])] != 0 {
			cards1[string(hand1[i])]++
		} else {
			cards1[string(hand1[i])] = 1
		}
	}
	cards2 := map[string]int{}
	for i := 0; i < len(hand2); i++ {
		if cards2[string(hand2[i])] != 0 {
			cards2[string(hand2[i])]++
		} else {
			cards2[string(hand2[i])] = 1
		}
	}
	if len(cards1) < len(cards2) {
		return true
	} else if len(cards1) > len(cards2) {
		return false
	} else {
		max1 := 0
		for _, nbcard := range cards1 {
			if nbcard > max1 {
				max1 = nbcard
			}
		}
		max2 := 0
		for _, nbcard := range cards2 {
			if nbcard > max2 {
				max2 = nbcard
			}
		}
		if max1 < max2 {
			return false
		} else if max1 > max2 {
			return true
		} else {
			value := map[string]int{"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2}
			for i := 0; i < 5; i++ {
				if value[string(hand1[i])] < value[string(hand2[i])] {
					return false
				} else if value[string(hand1[i])] > value[string(hand2[i])] {
					return true
				}
			}
		}
	}
	return false
}

func part1(cardsBid [][2]string) {
	orderedCard := [][2]string{}
	for _, hand := range cardsBid {
		toAdd := true
		for i := 0; i < len(orderedCard); i++ {
			if order(orderedCard[i][0], hand[0]) {
				orderedCard = append(orderedCard[:i+1], orderedCard[i:]...)
				orderedCard[i] = hand
				toAdd = false
				break
			}
		}
		if toAdd {
			orderedCard = append(orderedCard, hand)
		}
	}
	value := 0
	for index, hand := range orderedCard {
		bid, err := strconv.ParseInt(hand[1], 10, 0)
		if err != nil {
			panic(err)
		}
		value += int(bid) * (index + 1)
	}
	fmt.Println(value)
}

func order2(hand1 string, hand2 string) bool { // true if card 1 > card 2
	cards1 := map[string]int{}
	for i := 0; i < len(hand1); i++ {

		if cards1[string(hand1[i])] != 0 {
			cards1[string(hand1[i])]++
		} else {
			cards1[string(hand1[i])] = 1
		}
	}
	cards2 := map[string]int{}
	for i := 0; i < len(hand2); i++ {
		if cards2[string(hand2[i])] != 0 {
			cards2[string(hand2[i])]++
		} else {
			cards2[string(hand2[i])] = 1
		}
	}
	max1c1 := 0
	max2c1 := 0
	for card, nbcard := range cards1 {
		if nbcard > max1c1 && card != "J" {
			max2c1 = max1c1
			max1c1 = nbcard
		} else if nbcard > max2c1 && card != "J" {
			max2c1 = nbcard
		}
	}
	max1c2 := 0
	max2c2 := 0
	for card, nbcard := range cards2 {
		if nbcard > max1c2 && card != "J" {
			max2c2 = max1c2
			max1c2 = nbcard
		} else if nbcard > max2c2 && card != "J" {
			max2c2 = nbcard
		}
	}
	max1c1 += cards1["J"]
	max1c2 += cards2["J"]
	if max1c1 > max1c2 {
		return true
	} else if max1c1 < max1c2 {
		return false
	} else if max2c1 > max2c2 {
		return true
	} else if max2c1 < max2c2 {
		return false
	} else {
		value := map[string]int{"A": 14, "K": 13, "Q": 12, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2, "J": 0}
		for i := 0; i < 5; i++ {
			if value[string(hand1[i])] < value[string(hand2[i])] {
				return false
			} else if value[string(hand1[i])] > value[string(hand2[i])] {
				return true
			}
		}
	}
	return false
}

func part2(cardsBid [][2]string) {
	orderedCard := [][2]string{}
	for _, hand := range cardsBid {
		toAdd := true
		for i := 0; i < len(orderedCard); i++ {
			if order2(orderedCard[i][0], hand[0]) {
				orderedCard = append(orderedCard[:i+1], orderedCard[i:]...)
				orderedCard[i] = hand
				toAdd = false
				break
			}
		}
		if toAdd {
			orderedCard = append(orderedCard, hand)
		}
	}
	value := 0
	for index, hand := range orderedCard {

		bid, err := strconv.ParseInt(hand[1], 10, 0)
		if err != nil {
			panic(err)
		}
		value += int(bid) * (index + 1)
	}
	fmt.Println(value)
}

func main() {
	cardBid := openFile()
	part1(cardBid)
	part2(cardBid)
}
