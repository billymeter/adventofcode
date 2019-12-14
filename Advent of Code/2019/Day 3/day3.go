package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type operation struct {
	direction byte
	magnitude int
}

func main() {
	input := readInput("input.txt")

	fmt.Println("Advent of Code 2019 Day 3")
	p1 := partOne(input)
	fmt.Println("Part One:", p1)

}

func partOne(input [][]operation) int {
	cursor := point{0, 0}
	grid := make(map[point]map[int]bool)
	for i, l := range input {
		for _, o := range l {
			cursor = drawLine(o, cursor, i, grid)
		}
		cursor = point{0, 0}
	}

	intersections := findIntersections(grid)

	var distances []int
	for _, i := range intersections {
		distances = append(distances, calculateManhattanDistance(i))
	}
	sort.Ints(distances)
	return distances[0]
}

func calculateManhattanDistance(p point) int {
	centralPort := point{0, 0}
	x := abs(p.x - centralPort.x)
	y := abs(p.y - centralPort.y)
	return x + y
}

func findIntersections(grid map[point]map[int]bool) []point {
	var candidates []point
	for k := range grid {
		if grid[k][1] && grid[k][2] {
			candidates = append(candidates, k)
		}
	}
	return candidates
}

func drawLine(o operation, start point, wire int, grid map[point]map[int]bool) point {
	shift := 1 << wire
	switch o.direction {
	case 'U':
		for y := 1; y <= o.magnitude; y++ {
			p := point{start.x, start.y + y}
			if grid[p] == nil {
				grid[p] = make(map[int]bool)
			}
			grid[p][shift] = true
		}
		return point{start.x, start.y + o.magnitude}
	case 'D':
		for y := 1; y <= o.magnitude; y++ {
			p := point{start.x, start.y - y}
			if grid[p] == nil {
				grid[p] = make(map[int]bool)
			}
			grid[p][shift] = true
		}
		return point{start.x, start.y - o.magnitude}
	case 'R':
		for x := 1; x <= o.magnitude; x++ {
			p := point{start.x + x, start.y}
			if grid[p] == nil {
				grid[p] = make(map[int]bool)
			}
			grid[p][shift] = true
		}
		return point{start.x + o.magnitude, start.y}
	case 'L':
		for x := 1; x <= o.magnitude; x++ {
			p := point{start.x - x, start.y}
			if grid[p] == nil {
				grid[p] = make(map[int]bool)
			}
			grid[p][shift] = true
		}
		return point{start.x - o.magnitude, start.y}
	}
	fmt.Println("i came into a bad area")
	return point{0, 0}
}

func parseLine(s string) []operation {
	moves := strings.Split(s, ",")
	var ops []operation
	for _, m := range moves {
		direction := m[0]
		magnitude, _ := strconv.Atoi(m[1:])
		ops = append(ops, operation{direction, magnitude})
	}
	return ops
}

func readInput(filename string) [][]operation {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	var ops [][]operation
	s := bufio.NewScanner(file)
	for s.Scan() {
		l := s.Text()
		ops = append(ops, parseLine(l))
	}

	return ops
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
