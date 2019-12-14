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
	// input := readInput("test.txt")

	fmt.Println("Advent of Code 2019 Day 3")
	p1 := partOne(input)
	fmt.Println("Part One:", p1)
	p2 := partTwo(input)
	fmt.Println("Part Two:", p2)
}

func partOne(input [][]operation) int {
	cursor := point{0, 0}
	grid := make(map[point]map[int]int)
	for i, l := range input {
		count := 0
		for _, o := range l {
			cursor = drawLine(o, cursor, i, count, grid)
			count += o.magnitude
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

func partTwo(input [][]operation) int {
	cursor := point{0, 0}
	grid := make(map[point]map[int]int)
	for i, l := range input {
		count := 0
		for _, o := range l {
			cursor = drawLine(o, cursor, i, count, grid)
			count += o.magnitude
		}
		cursor = point{0, 0}
	}

	intersections := findIntersections(grid)
	return findMinLengthAtIntersection(intersections, grid)
}

func findMinLengthAtIntersection(points []point, grid map[point]map[int]int) int {
	min := grid[points[0]][1] + grid[points[0]][2]
	for _, p := range points {
		if min > grid[p][1]+grid[p][2] {
			min = grid[p][1] + grid[p][2]
		}
	}
	return min
}

func calculateManhattanDistance(p point) int {
	centralPort := point{0, 0}
	x := abs(p.x - centralPort.x)
	y := abs(p.y - centralPort.y)
	return x + y
}

func findIntersections(grid map[point]map[int]int) []point {
	var candidates []point
	for k := range grid {
		if grid[k][1] > 0 && grid[k][2] > 0 {
			candidates = append(candidates, k)
		}
	}
	return candidates
}

func drawLine(o operation, start point, wire, count int, grid map[point]map[int]int) point {
	shift := 1 << wire
	switch o.direction {
	case 'U':
		for y := 1; y <= o.magnitude; y++ {
			count++
			p := point{start.x, start.y + y}
			if grid[p] == nil {
				grid[p] = make(map[int]int)
			}
			if grid[p][shift] == 0 {
				grid[p][shift] = count
			}
		}
		return point{start.x, start.y + o.magnitude}
	case 'D':
		for y := 1; y <= o.magnitude; y++ {
			count++
			p := point{start.x, start.y - y}
			if grid[p] == nil {
				grid[p] = make(map[int]int)
			}
			if grid[p][shift] == 0 {
				grid[p][shift] = count
			}
		}
		return point{start.x, start.y - o.magnitude}
	case 'R':
		for x := 1; x <= o.magnitude; x++ {
			count++
			p := point{start.x + x, start.y}
			if grid[p] == nil {
				grid[p] = make(map[int]int)
			}
			if grid[p][shift] == 0 {
				grid[p][shift] = count
			}
		}
		return point{start.x + o.magnitude, start.y}
	case 'L':
		for x := 1; x <= o.magnitude; x++ {
			count++
			p := point{start.x - x, start.y}
			if grid[p] == nil {
				grid[p] = make(map[int]int)
			}
			if grid[p][shift] == 0 {
				grid[p][shift] = count
			}
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
