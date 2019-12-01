package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dimension struct {
	length int
	width  int
	height int
}

func main() {
	totalPaper := 0
	ribbon := 0
	input := readInput("input.txt")
	for _, v := range input {
		paper, rib := calcBoxArea(v)
		totalPaper += paper
		ribbon += rib
	}
	fmt.Println("Advent of Code 2015 Day Two")
	fmt.Println("Part One:", totalPaper)
	fmt.Println("Part Two:", ribbon)
}

func calcBoxArea(d Dimension) (int, int) {
	lw := calcArea(d.length, d.width)
	lh := calcArea(d.length, d.height)
	wh := calcArea(d.width, d.height)
	e1, e2 := slack(d.length, d.width, d.height)
	extra := e1 * e2
	ribbon := 2*e1 + 2*e2 + d.length*d.width*d.height
	return lw + lh + wh + extra, ribbon
}

func calcArea(x, y int) int {
	return 2 * x * y
}

func slack(a, b, c int) (int, int) {
	if a < b {
		if b < c {
			return a, b
		}
		return a, c
	}
	if a < c {
		return a, b
	}
	return b, c
}

func readInput(filename string) []Dimension {
	var dimensions []Dimension
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(file)

	for s.Scan() {
		line := s.Text()
		dims := strings.Split(line, "x")
		l, _ := strconv.Atoi(dims[0])
		w, _ := strconv.Atoi(dims[1])
		h, _ := strconv.Atoi(dims[2])
		d := Dimension{
			length: l,
			width:  w,
			height: h,
		}
		dimensions = append(dimensions, d)
	}
	return dimensions
}
