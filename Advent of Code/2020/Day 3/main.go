package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	x int
	y int
}

func main() {
	grid := readInput("input")
	trees := part1(grid)
	fmt.Println("Part 1:", trees)
}

func part1(grid map[point]bool) int {
	return traverseSlope(grid, 3, 1)
}

func traverseSlope(grid map[point]bool, right, down int) int {
	treesEncountered := 0
	x, y := right, down
	for y <= 322 {
		p := point{x, y}
		if grid[p] {
			treesEncountered++
		}
		x += right % 31
		y += down
	}
	return treesEncountered
}

func readInput(filename string) map[point]bool {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	y := 0
	grid := make(map[point]bool)
	for scanner.Scan() {
		x := 0
		for _, c := range scanner.Text() {
			p := point{x, y}
			if c == '#' {
				grid[p] = true
			} else {
				grid[p] = false
			}
			x++
		}
		y++
	}
	return grid
}
