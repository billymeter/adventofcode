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
	// input := readInput("input.txt")
	input := readInput("test.txt")

	cursor := point{0, 0}
	grid := make(map[point]int)
	for i, l := range input {
		fmt.Println(l)
		for _, o := range l {
			cursor = drawLine(o, cursor, i, grid)
		}
		cursor = point{0, 0}
	}

	intersections := findIntersections(grid)

	fmt.Println("intersections:", intersections)

	var distances []int
	for _, i := range intersections {
		distances = append(distances, calculateManhattanDistance(i))
	}
	sort.Ints(distances)
	fmt.Println(distances)
}

func calculateManhattanDistance(p point) int {
	centralPort := point{0, 0}
	val := (p.x - centralPort.x) + (p.y - centralPort.y)
	if val < 0 {
		val *= -1
	}
	return val
}

func findIntersections(grid map[point]int) []point {
	var candidates []point
	for k, v := range grid {
		if v > 2 {
			// if v&3 == 3 {
			fmt.Println(k, "has a value of", v)
			candidates = append(candidates, k)
		}
	}
	return candidates
}

func drawLine(o operation, start point, wire int, grid map[point]int) point {
	// 0 for no wire in path, 1 for one wire, and 11 for two wires
	shift := 1 << wire
	switch o.direction {
	case 'U':
		for y := 1; y < o.magnitude; y++ {
			p := point{start.x, start.y + y}
			fmt.Println("Point", p, "marked for wire", wire+1)
			grid[p] |= shift
		}
		return point{start.x, start.y + o.magnitude}
	case 'D':
		for y := 1; y < o.magnitude; y++ {
			p := point{start.x, start.y - y}
			fmt.Println("Point", p, "marked for wire", wire+1)
			grid[p] |= shift
		}
		return point{start.x, start.y - o.magnitude}
	case 'R':
		for x := 1; x < o.magnitude; x++ {
			p := point{start.x + x, start.y}
			fmt.Println("Point", p, "marked for wire", wire+1)
			grid[p] |= shift
		}
		return point{start.x + o.magnitude, start.y}
	case 'L':
		for x := 1; x < o.magnitude; x++ {
			p := point{start.x - x, start.y}
			fmt.Println("Point", p, "marked for wire", wire+1)
			grid[p] |= shift
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
