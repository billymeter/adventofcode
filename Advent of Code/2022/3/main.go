package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

type Elves int

const (
	ElfOne   Elves = 1
	ElfTwo         = 2
	ElfThree       = 4
)

func readInput(filename string) []string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	data := string(file)
	return strings.Split(data, "\n")
}

func getItemPriority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item - 'A' + 27)
	}
	return int(item - 'a' + 1)
}

func day3p1(rucksackContents []string) int {
	// this is pretty sucky. there's got to be a better way of doing this
	sum := 0
	for _, pack := range rucksackContents {
		amountOfItems := len(pack)
		compartmentOne := pack[:amountOfItems/2]
		compartmentTwo := pack[amountOfItems/2:]

		c1items := make(map[rune]bool)
		c2items := make(map[rune]bool)
		for _, item := range compartmentOne {
			c1items[item] = true
		}
		for _, item := range compartmentTwo {
			c2items[item] = true
		}

		for i := range c1items {
			if c2items[i] {
				sum += getItemPriority(i)
			}
		}
	}
	return sum
}

func day3p2(rucksackContents []string) int {
	sum := 0
	for i := 0; i < len(rucksackContents)-3; i += 3 {
		itemsInPacks := make(map[rune]int)
		for _, item := range rucksackContents[i] {
			itemsInPacks[item] |= int(ElfOne)
		}
		for _, item := range rucksackContents[i+1] {
			itemsInPacks[item] |= int(ElfTwo)
		}
		for _, item := range rucksackContents[i+2] {
			itemsInPacks[item] |= int(ElfThree)
		}

		for item := range itemsInPacks {
			if itemsInPacks[item] == int(ElfOne|ElfTwo|ElfThree) {
				sum += getItemPriority(item)
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(day3p1(readInput("input")))
	fmt.Println(day3p2(readInput("input")))
}
