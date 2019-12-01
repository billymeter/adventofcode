package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Advent of Code 2015 Day One")
	input := readInput("input.txt")
	floor := elevator1(input)
	fmt.Println("Part One:", floor)
	floor = elevator2(input)
	fmt.Println("Part Two:", floor)
}

func readInput(filename string) []rune {
	var instructions []rune
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			instructions = append(instructions, c)
		}
	}

	return instructions
}

func elevator1(instructions []rune) int {
	floor := 0
	for _, v := range instructions {
		if string(v) == "(" {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func elevator2(instructions []rune) int {
	floor := 0
	for i, v := range instructions {
		if string(v) == "(" {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			return i + 1 // zero based indexing messed me up. Add one.
		}
	}
	return 0
}
