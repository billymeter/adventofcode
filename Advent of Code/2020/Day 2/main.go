package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	min    int
	max    int
	letter byte
	line   string
}

func main() {
	entries := readInput("input")
	valid := part1(entries)
	fmt.Println("Part 1:", valid)
	valid = part2(entries)
	fmt.Println("Part 2:", valid)
}

func part1(entries []entry) int {
	validCount := 0
	for _, e := range entries {
		if validateEntry(e) {
			validCount++
		}
	}
	return validCount
}

func part2(entries []entry) int {
	validCount := 0
	for _, e := range entries {
		first := e.line[e.min-1]
		last := e.line[e.max-1]
		if first != last {
			if first == e.letter || last == e.letter {
				validCount++
			}
		}
	}
	return validCount
}

func validateEntry(e entry) bool {
	count := strings.Count(e.line, string(e.letter))
	if count >= e.min && count <= e.max {
		return true
	}
	return false
}

func parseLine(line string) entry {
	// example line: 1-3 a: abcde
	components := strings.Split(line, " ")
	limits := strings.Split(components[0], "-")
	min, _ := strconv.Atoi(limits[0])
	max, _ := strconv.Atoi(limits[1])
	return entry{min, max, components[1][0], components[2]}
}

func readInput(filename string) []entry {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var entries []entry
	for scanner.Scan() {
		e := parseLine(scanner.Text())
		entries = append(entries, e)
	}
	return entries
}
