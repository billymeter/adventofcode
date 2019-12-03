package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type op struct {
	val int
}

var (
	on     = op{0}
	off    = op{1}
	toggle = op{2}
)

type operation struct {
	opcode op
	start  point
	end    point
}

func main() {
	opcodes := readInput("input.txt")
	fmt.Println("Advent of Code 2015 Day 6")
	p1 := partOne(opcodes)
	fmt.Println("Part One:", p1)
	p2 := partTwo(opcodes)
	fmt.Println("Part Two:", p2)
}

func partOne(opcodes []operation) int {
	grid := make(map[point]bool)

	for _, o := range opcodes {
		executeOperation(o, grid)
	}

	lightsLit := 0
	for k := range grid {
		if grid[k] {
			lightsLit++
		}
	}
	return lightsLit
}

func partTwo(opcodes []operation) int {
	grid := make(map[point]int)

	for _, o := range opcodes {
		executeOperation2(o, grid)
	}

	brightness := 0
	for k := range grid {
		brightness += grid[k]
	}
	return brightness
}

func executeOperation(o operation, grid map[point]bool) {
	for x := o.start.x; x <= o.end.x; x++ {
		for y := o.start.y; y <= o.end.y; y++ {
			switch o.opcode {
			case on:
				grid[point{x, y}] = true
			case off:
				grid[point{x, y}] = false
			default:
				grid[point{x, y}] = !grid[point{x, y}]
			}
		}
	}
}

func executeOperation2(o operation, grid map[point]int) {
	for x := o.start.x; x <= o.end.x; x++ {
		for y := o.start.y; y <= o.end.y; y++ {
			switch o.opcode {
			case on:
				grid[point{x, y}]++
			case off:
				if grid[point{x, y}] <= 0 {
					grid[point{x, y}] = 0
				} else {
					grid[point{x, y}]--
				}
			default:
				grid[point{x, y}] += 2
			}
		}
	}
}

func parseLine(s string) operation {
	tokens := strings.Split(s, " ")

	if tokens[0] == "toggle" {
		coords := strings.Split(tokens[1], ",")
		x1, _ := strconv.Atoi(coords[0])
		y1, _ := strconv.Atoi(coords[1])

		coords = strings.Split(tokens[3], ",")
		x2, _ := strconv.Atoi(coords[0])
		y2, _ := strconv.Atoi(coords[1])

		return operation{toggle, point{x1, y1}, point{x2, y2}}
	}
	coords := strings.Split(tokens[2], ",")
	x1, _ := strconv.Atoi(coords[0])
	y1, _ := strconv.Atoi(coords[1])

	coords = strings.Split(tokens[4], ",")
	x2, _ := strconv.Atoi(coords[0])
	y2, _ := strconv.Atoi(coords[1])

	if tokens[1] == "on" {
		return operation{on, point{x1, y1}, point{x2, y2}}
	}

	return operation{off, point{x1, y1}, point{x2, y2}}
}

func readInput(filename string) []operation {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	var opcodes []operation
	for _, l := range lines {
		opcodes = append(opcodes, parseLine(l))
	}

	return opcodes
}
