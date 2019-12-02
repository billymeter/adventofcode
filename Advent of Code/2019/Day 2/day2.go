package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput("input.txt")
	fmt.Println("Advent of Code 2019 Day 2")
	p1 := runProgram(12, 2, input)
	fmt.Println("Part One", p1)

	input = readInput("input.txt")
	p2 := partTwo(input)
	fmt.Println("Part Two:", p2)
}

func runProgram(noun, verb int, nums []int) int {
	ip := 0
	nums[1] = noun
	nums[2] = verb
	for processOpcode(ip, nums) {
		ip += 4
	}
	return nums[0]
}

func partTwo(nums []int) int {
	noun, verb := 0, 0
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			memory := make([]int, len(nums))
			copy(memory, nums)
			val := runProgram(n, v, memory)
			if val == 19690720 {
				noun, verb = n, v
				break
			}
		}
	}
	return 100*noun + verb
}

func processOpcode(ip int, data []int) bool {
	if ip > len(data) || data[ip] > len(data) {
		return false
	}
	switch data[ip] {
	case 1:
		data[data[ip+3]] = data[data[ip+1]] + data[data[ip+2]]
	case 2:
		data[data[ip+3]] = data[data[ip+1]] * data[data[ip+2]]
	case 99:
		return false
	}
	return true
}

func readInput(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var codes []int
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		l := strings.Split(line, ",")
		for _, v := range l {
			n, _ := strconv.Atoi(v)
			codes = append(codes, n)
		}

	}
	return codes
}
