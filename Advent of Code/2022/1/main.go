package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func readInput(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totals []int

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			totals = append(totals, sum)
			sum = 0
			continue
		}

		num, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}

		sum += num
	}

	return totals
}

func part1(totals []int) int {
	max := 0
	for _, v := range totals {
		if max < v {
			max = v
		}
	}
	return max
}

func part2(totals []int) int {
	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})

	return totals[0] + totals[1] + totals[2]
}

func main() {
	totals := readInput("input")
	println("day 1 part 1:", part1(totals))
	println("day 1 part 2:", part2(totals))
}
