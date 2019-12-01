package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	input := readInput("input.txt")
	sum := computeFuel1(input)
	fmt.Println("Day One Part One Fuel needed:", sum)
	sum = computeFuel2(input)
	fmt.Println("Day One Part Two Fuel needed:", sum)
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

func fuelRequirements(mass int) int {
	fuel := int(math.Floor(float64(mass) / 3.0))
	fuel -= 2
	return fuel
}

func computeFuel1(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += fuelRequirements(v)
	}
	return sum
}

func computeFuel2(nums []int) int {
	sum := 0
	for _, v := range nums {
		fuel := fuelRequirements(v)
		for {
			if fuel <= 0 {
				break
			} else {
				sum += fuel
				fuel = fuelRequirements(fuel)
			}
		}
	}
	return sum
}
