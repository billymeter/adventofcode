package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation int

const (
	ADD Operation = iota
	NOP
)

type Instruction struct {
	op              Operation
	value           int
	cyclesRemaining int
}

func readInput(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []string
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	return instructions
}

func parseInstructions(lines []string) []Instruction {
	var instructions []Instruction
	for _, l := range lines {
		components := strings.Split(l, " ")
		switch components[0] {
		case "noop":
			instructions = append(instructions, Instruction{NOP, 0, 1})
		case "addx":
			value, _ := strconv.Atoi(components[1])
			instructions = append(instructions, Instruction{ADD, value, 2})
		}
	}
	return instructions
}

func tick(instruction Instruction) {
	instruction.cyclesRemaining--
}

func day10p1() {
	input := readInput("input")
	instructions := parseInstructions(input)
	X := 1
	cycle := 0
	sumOfSignalStrength := 0
	for _, instr := range instructions {
		for i := instr.cyclesRemaining; i > 0; i-- {
			tick(instr)
			cycle++
			if cycle == 20 || (cycle+20)%40 == 0 {
				sumOfSignalStrength += cycle * X
			}
		}
		if instr.op == ADD {
			X += instr.value
		}
	}

	fmt.Printf("%d\n", sumOfSignalStrength)
}

func day10p2() {
	input := readInput("input")
	instructions := parseInstructions(input)

	X := 1
	cycle := 0
	for _, instr := range instructions {
		for i := instr.cyclesRemaining; i > 0; i-- {
			if (cycle%40) == X-1 || (cycle%40) == X || (cycle%40) == X+1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}

			tick(instr)
			cycle++

			if cycle%40 == 0 {
				fmt.Println()
			}
		}
		if instr.op == ADD {
			X += instr.value
		}
	}
}

func main() {
	day10p1()
	day10p2()
}
