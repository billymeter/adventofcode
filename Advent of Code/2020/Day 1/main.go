package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	nums := readInput("input.txt")
	ans := part1(nums)
	fmt.Println("Part 1:", ans)
	ans = part2(nums)
	fmt.Println("Part 2:", ans)
}

func part1(nums []int) int {
	for i, v := range nums {
		for _, x := range nums[i+1:] {
			if v+x == 2020 {
				return v * x
			}
		}
	}
	return 0
}

func part2(nums []int) int {
	for i, x := range nums {
		for j, y := range nums[i+1:] {
			for _, z := range nums[i+j+1:] {
				if x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}
	return 0
}

func readInput(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var nums []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nums
}
