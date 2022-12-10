package main

import (
	"bufio"
	"fmt"
	"os"
)

func runeToInt(c rune) int {
	if c == '0' {
		return 0
	}
	if c == '1' {
		return 1
	}
	if c == '2' {
		return 2
	}
	if c == '3' {
		return 3
	}
	if c == '4' {
		return 4
	}
	if c == '5' {
		return 5
	}
	if c == '6' {
		return 6
	}
	if c == '7' {
		return 7
	}
	if c == '8' {
		return 8
	}
	return 9

}

func readInput(filename string) [][]int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	var grid [][]int

	for scanner.Scan() {
		var line []int
		for _, c := range scanner.Text() {
			n := runeToInt(c)
			line = append(line, n)
		}
		grid = append(grid, line)
	}
	return grid
}

func isTreeVisible(x, y int, grid [][]int) bool {
	tree := grid[x][y]
	blocked := false

	for i := 0; i < y; i++ {
		if grid[x][i] >= tree {
			blocked = true
		}
	}
	if !blocked {
		return true
	}

	blocked = false
	for i := y + 1; i < len(grid[x]); i++ {
		if grid[x][i] >= tree {
			blocked = true
		}
	}
	if !blocked {
		return true
	}

	blocked = false
	for i := 0; i < x; i++ {
		if grid[i][y] >= tree {
			blocked = true
		}
	}

	if !blocked {
		return true
	}

	blocked = false
	for i := x + 1; i < len(grid[x]); i++ {
		if grid[i][y] >= tree {
			blocked = true
		}
	}
	if !blocked {
		return true
	}

	return false
}

func treeScenicScore(x, y int, grid [][]int) int {
	tree := grid[x][y]
	up, down, left, right := 0, 0, 0, 0

	for i := y - 1; i >= 0; i-- {
		left++
		if grid[x][i] >= tree {
			break
		}
	}
	for i := y + 1; i < len(grid[x]); i++ {
		right++
		if grid[x][i] >= tree {
			break
		}
	}

	for i := x - 1; i >= 0; i-- {
		up++
		if grid[i][y] >= tree {
			break
		}
	}

	for i := x + 1; i < len(grid[x]); i++ {
		down++
		if grid[i][y] >= tree {
			break
		}
	}
	return up * down * left * right
}

func day8p1() {
	input := readInput("input")

	seen := 0
	for x, i := range input {
		for y := range i {
			vis := isTreeVisible(x, y, input)
			if vis {
				seen++
			}
		}
	}

	fmt.Printf("%d\n", seen)
}

func day8p2() {
	input := readInput("input")

	var score int
	max := 0
	for x, i := range input {
		if x == 0 || x == len(i)-1 {
			continue
		}
		for y := range i {
			if y == 0 || y == len(i)-1 {
				continue
			}
			score = treeScenicScore(x, y, input)
			if max < score {
				max = score
			}
		}
	}

	fmt.Printf("%d\n", max)
}

func main() {
	day8p1()
	day8p2()
}
