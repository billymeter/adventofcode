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

	var t Point
	// 8 possibilities
	// lagging from the right
	if head.x == tail.x-2 && head.y == tail.y {
		t = Point{tail.x - 1, tail.y}
	}
	// lagging from the left
	if head.x == tail.x+2 && head.y == tail.y {
		t = Point{tail.x + 1, tail.y}
	}

	// lagging from the top
	if head.x == tail.x && head.y == tail.y+2 {
		t = Point{tail.x, tail.y + 1}
	}
	// lagging from the bottom
	if head.x == tail.x && head.y == tail.y-2 {
		t = Point{tail.x, tail.y - 1}
	}

	// lagging from the lower diagonals
	if head.x == tail.x-1 && head.y == tail.y+2 || head.x == tail.x+1 && head.y == tail.y+2 {
		t = Point{head.x, tail.y + 1}
	}

	// lagging from the upper diagonals
	if head.x == tail.x-1 && head.y == tail.y-2 || head.x == tail.x+1 && head.y == tail.y-2 {
		t = Point{head.x, tail.y - 1}
	}

	// lagging from the right diagonals
	if head.x == tail.x-2 && head.y == tail.y+1 || head.x == tail.x-2 && head.y == tail.y-1 {
		t = Point{tail.x - 1, head.y}
	}

	// lagging from the left diagonals
	if head.x == tail.x+2 && head.y == tail.y+1 || head.x == tail.x+2 && head.y == tail.y-1 {
		t = Point{tail.x + 1, head.y}
	}
	return t
}

func simulateRopeMoves(moves []string) map[Point]struct{} {
	head := Point{0, 0}
	tail := Point{0, 0}
	visited := make(map[Point]struct{})
	visited[tail] = struct{}{}

	for _, m := range moves {
		component := strings.Split(m, " ")
		amount, _ := strconv.Atoi(component[1])
		for i := 0; i < amount; i++ {
			// fmt.Printf("h: %v, t: %v --> ", head, tail)
			switch component[0] {
			case "D":
				head.y--
			case "L":
				head.x--
			case "R":
				head.x++
			case "U":
				head.y++
			default:
				log.Fatal("unrecognized move", component[0])
			}
			tail = makeTailAdjacentToHead(head, tail)
			// fmt.Printf("h: %v, t: %v\n", head, tail)
			visited[tail] = struct{}{}
		}
	}

	return visited
}

func day9p1() {
	moves := readInput("input")
	visited := simulateRopeMoves(moves)
	fmt.Printf("%d\n", len(visited))
}

func day9p2() {
	moves := readInput("test")
	visited := simulateRopeMoves(moves)
	fmt.Printf("%d\n", len(visited))
}

func main() {
	day9p1()
	day9p2()
}
