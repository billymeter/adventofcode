package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Point struct {
	x int
	y int
}

func main() {
	fmt.Println("Advent of Code 2015 Day Three")
	p1 := partOne("input.txt")
	fmt.Println("Part One:", p1)
	p2 := partTwo("input.txt")
	fmt.Println("Part Two:", p2)

}

func partOne(filename string) int {
	x, y := 0, 0
	visited := make(map[Point]int)

	visited[Point{x, y}]++

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)
	for {
		c, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		switch string(c) {
		case ">":
			x++
		case "<":
			x--
		case "^":
			y++
		case "v":
			y--
		default:
			fmt.Println("something goofed:")
		}
		visited[Point{x, y}]++
	}

	return len(visited)
}

func partTwo(filename string) int {
	sx, sy := 0, 0
	rx, ry := 0, 0
	robot := false
	visited := make(map[Point]int)

	visited[Point{sx, sy}] = 2

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)
	for {
		c, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if robot {
			switch string(c) {
			case ">":
				rx++
			case "<":
				rx--
			case "^":
				ry++
			case "v":
				ry--
			}
			visited[Point{rx, ry}]++
			robot = false
		} else {
			switch string(c) {
			case ">":
				sx++
			case "<":
				sx--
			case "^":
				sy++
			case "v":
				sy--
			}
			visited[Point{sx, sy}]++
			robot = true
		}
	}
	return len(visited)
}
