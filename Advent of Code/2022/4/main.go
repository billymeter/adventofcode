package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func readInput(filename string) ([]Range, []Range) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var elf1, elf2 []Range
	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")
		first := strings.Split(assignments[0], "-")
		second := strings.Split(assignments[1], "-")
		s1, _ := strconv.Atoi(first[0])
		e1, _ := strconv.Atoi(first[1])
		s2, _ := strconv.Atoi(second[0])
		e2, _ := strconv.Atoi(second[1])
		elf1 = append(elf1, Range{s1, e1})
		elf2 = append(elf2, Range{s2, e2})
	}
	return elf1, elf2
}

func day4p1(elf1, elf2 []Range) int {
	numContained := 0

	for i := range elf1 {
		if elf1[i].start <= elf2[i].start && elf1[i].end >= elf2[i].end {
			numContained++
		} else if elf1[i].start >= elf2[i].start && elf1[i].end <= elf2[i].end {
			numContained++
		}
	}

	return numContained
}

func day4p2(elf1, elf2 []Range) int {
	overlapped := 0

	for i := range elf1 {
		if (elf2[i].start <= elf1[i].end && elf2[i].start >= elf1[i].start) ||
			(elf2[i].end >= elf1[i].start && elf2[i].end <= elf1[i].end) {
			overlapped++
		} else if (elf1[i].start <= elf2[i].end && elf1[i].start >= elf2[i].start) ||
			(elf1[i].end >= elf2[i].start && elf1[i].end <= elf2[i].end) {
			overlapped++
		}
	}

	return overlapped
}

func main() {
	elf1, elf2 := readInput("input")
	fmt.Println(day4p1(elf1, elf2))
	fmt.Println(day4p2(elf1, elf2))
}
