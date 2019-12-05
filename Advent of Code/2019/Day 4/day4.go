package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "272091-815432"
	ranges := strings.Split(input, "-")
	lower, _ := strconv.Atoi(ranges[0])
	upper, _ := strconv.Atoi(ranges[1])

	fmt.Println("Advent of Code 2019 Day 4")
	p1 := partOne(lower, upper)
	fmt.Println("Part One:", p1)
	p2 := partTwo(lower, upper)
	fmt.Println("Part Two:", p2)
}

func partOne(lower, upper int) int {
	count := 0
	for i := lower; i <= upper; i++ {
		num := intStringToIntSlice(strconv.Itoa(i))
		if hasDigitsNeverDecreaseLeftToRight(num) && hasTwoSameAdjacentDigits(num) {
			count++
		}
	}
	return count
}

func partTwo(lower, upper int) int {
	count := 0
	// for i := lower; i <= upper; i++ {
	// 	num := intStringToIntSlice(strconv.Itoa(i))
	// 	if hasDigitsNeverDecreaseLeftToRight(num) && hasTwoSameAdjacentDigits(num) {
	// 		count++
	// 	}
	// }
	tests := [][]int{[]int{1, 1, 2, 2, 3, 3}, []int{1, 2, 3, 4, 4, 4}, []int{1, 1, 1, 1, 3, 3}}
	for _, t := range tests {
		fmt.Println(t, "returns ", hasTwoAdjacentDigitsNotPartOfLargerGroup(t))
		// hasTwoAdjacentDigitsNotPartOfLargerGroup(t)
	}

	return count
}

func hasTwoAdjacentDigitsNotPartOfLargerGroup(num []int) bool {
	groups := make(map[int]bool)
	streak := 1
	lastDigit := num[0]
	for i, d := range num {

		if i == 0 {
			continue
		}
		fmt.Println("i:", i, "; d:", d)
		if lastDigit == d {
			streak++
		} else {
			groups[streak] = true
			streak = 1
		}
		lastDigit = d
	}
	fmt.Println(num, ":", groups)
	if groups[2] {
		return true
	}
	return false
}

func hasDigitsNeverDecreaseLeftToRight(num []int) bool {
	for i := 0; i < len(num)-1; i++ {
		if num[i] > num[i+1] {
			return false
		}
	}
	return true
}

func hasTwoSameAdjacentDigits(num []int) bool {
	for i := 0; i < len(num)-1; i++ {
		if num[i] == num[i+1] {
			return true
		}
	}
	return false
}

func intStringToIntSlice(s string) []int {
	var nums []int
	for _, c := range s {
		v, _ := strconv.Atoi(string(c))
		nums = append(nums, v)
	}
	return nums
}
