package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func readInput(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func isTailAdjacentToHead(head, tail Point) bool {
	// 9 positions to check
	// upper left diagonal
	if head.x-1 == tail.x && head.y+1 == tail.y {
		return true
	}
	// directly above
	if head.x == tail.x && head.y+1 == tail.y {
		return true
	}
	// upper right diagonal
	if head.x+1 == tail.x && head.y+1 == tail.y {
		return true
	}

	// directly left
	if head.x-1 == tail.x && head.y == tail.y {
		return true
	}
	// directly on
	if head.x == tail.x && head.y == tail.y {
		return true
	}
	// directly right
	if head.x+1 == tail.x && head.y == tail.y {
		return true
	}

	// lower left diagonal
	if head.x-1 == tail.x && head.y-1 == tail.y {
		return true
	}
	// directly below
	if head.x == tail.x && head.y-1 == tail.y {
		return true
	}
	// lower right diagonal
	if head.x+1 == tail.x && head.y-1 == tail.y {
		return true
	}

	return false
}

func makeTailAdjacentToHead(head, tail Point) Point {
	// if already adjacent, don't adjust tail position
	if isTailAdjacentToHead(head, tail) {
		return tail
	}

	// 8 possibilities
	// lagging from the right
	if head.x == tail.x-2 && head.y == tail.y {
		return Point{tail.x - 1, tail.y}
	}
	// lagging from the left
	if head.x == tail.x+2 && head.y == tail.y {
		return Point{tail.x + 1, tail.y}
	}

	// lagging from the top
	if head.x == tail.x && head.y == tail.y+2 {
		return Point{tail.x, tail.y + 1}
	}
	// lagging from the bottom
	if head.x == tail.x && head.y == tail.y-2 {
		return Point{tail.x, tail.y - 1}
	}

	// lagging from the lower diagonals
	if head.x == tail.x-1 && head.y == tail.y+2 || head.x == tail.x+1 && head.y == tail.y+2 {
		return Point{head.x, tail.y + 1}
	}

	// lagging from the upper diagonals
	if head.x == tail.x-1 && head.y == tail.y-2 || head.x == tail.x+1 && head.y == tail.y-2 {
		return Point{head.x, tail.y - 1}
	}

	// lagging from the right diagonals
	if head.x == tail.x-2 && head.y == tail.y+1 || head.x == tail.x-2 && head.y == tail.y-1 {
		return Point{tail.x - 1, head.y}
	}

	// lagging from the left diagonals
	if head.x == tail.x+2 && head.y == tail.y+1 || head.x == tail.x+2 && head.y == tail.y-1 {
		return Point{tail.x + 1, head.y}
	}

	if head.x == tail.x-2 && head.y == tail.y-2 {
		return Point{tail.x - 1, tail.y - 1}
	}
	if head.x == tail.x-2 && head.y == tail.y+2 {
		return Point{tail.x - 1, tail.y + 1}
	}
	if head.x == tail.x+2 && head.y == tail.y-2 {
		return Point{tail.x + 1, tail.y - 1}
	}
	if head.x == tail.x+2 && head.y == tail.y+2 {
		return Point{tail.x + 1, tail.y + 1}
	}

	t := Point{-100, -100}
	return t
}

func simulateRopeMoves(moves []string, numKnots int) map[Point]struct{} {
	var knots []Point
	for i := 0; i < numKnots; i++ {
		knots = append(knots, Point{0, 0})
	}

	visited := make(map[Point]struct{})

	for _, m := range moves {
		component := strings.Split(m, " ")
		amount, _ := strconv.Atoi(component[1])
		for i := 0; i < amount; i++ {
			switch component[0] {
			case "D":
				knots[0].y--
			case "L":
				knots[0].x--
			case "R":
				knots[0].x++
			case "U":
				knots[0].y++
			default:
				log.Fatal("unrecognized move", component[0])
			}

			for i := 1; i < numKnots; i++ {
				knots[i] = makeTailAdjacentToHead(knots[i-1], knots[i])
			}
			visited[knots[len(knots)-1]] = struct{}{}
		}
	}

	return visited
}

func day9p1() {
	moves := readInput("input")
	visited := simulateRopeMoves(moves, 2)

	fmt.Printf("%d\n", len(visited))
}

func day9p2() {
	moves := readInput("input")
	visited := simulateRopeMoves(moves, 10)

	fmt.Printf("%d\n", len(visited))
}

func main() {
	day9p1()
	day9p2()
}
