package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type trool struct {
	// lame attempt at a boolean-like value but with three
	// states
	val int
}

var (
	none = trool{0}
	one  = trool{1}
	both = trool{2}
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
	fmt.Println(input)
}

func drawLine(o operation, start point, grid map[point]trool) {
	switch o.direction {
	case 'U':
		for y := start.y; y <= o.magnitude; y++ {

		}
	}
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
