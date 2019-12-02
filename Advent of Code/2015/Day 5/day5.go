package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readInput("input.txt")
	fmt.Println("Advent of Code 2015 Day 5")
	p1 := partOne(input)
	fmt.Println("Part One:", p1)
	input = []string{"qjhvhtzxzqqjkmpb", "xxyxx", "uurcxstgmygtbstg", "ieodomkazucvgmuy"}
	p2 := partTwo(input)
	fmt.Println("Part Two:", p2)
}

func partOne(input []string) int {
	niceStrings := 0
	for _, s := range input {
		if isNice(s) {
			niceStrings++
		}
	}
	return niceStrings
}

func partTwo(input []string) int {
	niceStrings := 0
	for _, s := range input {
		if isNice2(s) {
			fmt.Println(s, "is nice")
			niceStrings++
		}
	}
	return niceStrings
}

func isNice2(s string) bool {
	return hasPairOfRepeats(s) && hasRepeatLetterWithAnotherLetterInBetween(s)
}

func hasPairOfRepeats(s string) bool {
	pairs := make(map[string]int)
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return false
		}
		pairs[s[i:i+2]]++
	}

	for k := range pairs {
		fmt.Println(k)
		if pairs[k] > 1 {
			fmt.Println(s, "has repeating pair:", k)
			return true
		}
	}
	return false
}

func hasRepeatLetterWithAnotherLetterInBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func isNice(s string) bool {
	return hasThreeVowels(s) && hasDoubleLetters(s) && hasNoBadPairs(s)
}

func hasThreeVowels(s string) bool {
	vowels := 0
	for _, c := range s {
		switch string(c) {
		case "a":
			vowels++
		case "e":
			vowels++
		case "i":
			vowels++
		case "o":
			vowels++
		case "u":
			vowels++
		}
		if vowels >= 3 {
			return true
		}
	}
	return false
}

func hasDoubleLetters(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func hasNoBadPairs(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if pair == "ab" {
			return false
		}
		if pair == "cd" {
			return false
		}
		if pair == "pq" {
			return false
		}
		if pair == "xy" {
			return false
		}
	}
	return true
}

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
