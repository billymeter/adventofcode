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
	Add Operation = iota
	Multiply
)

type Monkey struct {
	items           []int
	operation       Operation
	operationAmount int
	test            int
	testTrue        int
	testFalse       int
	timesInspected  int
}

func readInput(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func parseNotesToMonkies(lines []string) []*Monkey {
	var monkies []*Monkey
	monkey := 0
	for _, l := range lines {
		components := strings.Split(strings.Trim(l, " "), " ")

		switch components[0] {
		case "Monkey":
			monkies = append(monkies, &Monkey{[]int{}, Add, 0, 0, 0, 0, 0})
		case "Starting":
			for _, item := range components[2:] {
				num, _ := strconv.Atoi(strings.Trim(item, ","))
				monkies[monkey].items = append(monkies[monkey].items, num)
			}
		case "Operation:":
			if components[4] == "*" {
				monkies[monkey].operation = Multiply
			}
			num, _ := strconv.Atoi(components[5])
			if num == 0 {
				num = -100
			}
			monkies[monkey].operationAmount = num
		case "Test:":
			num, _ := strconv.Atoi(components[3])
			monkies[monkey].test = num
		case "If":
			if components[1] == "true:" {
				num, _ := strconv.Atoi(components[5])
				monkies[monkey].testTrue = num
			} else {
				num, _ := strconv.Atoi(components[5])
				monkies[monkey].testFalse = num
			}
		case "":
			monkey++
		}
	}
	return monkies
}

func runRound(monkies []*Monkey, divide bool) {
	divisor := 1
	for _, m := range monkies {
		// this is needed to get a ring of integers modulo Pi_i d_i.
		// that is, the product of all the tests of each monkey
		divisor *= m.test
	}

	for _, m := range monkies {
		for idx := range m.items {
			if divide {
				if m.operation == Add {
					if m.operationAmount < 0 {
						m.items[idx] += m.items[idx]
					} else {
						m.items[idx] += m.operationAmount
					}
				} else {
					if m.operationAmount < 0 {
						m.items[idx] *= m.items[idx]
					} else {
						m.items[idx] *= m.operationAmount
					}
				}

				m.items[idx] /= 3

			} else {

				if m.operation == Add {
					if m.operationAmount < 0 {
						m.items[idx] += m.items[idx]
					} else {
						m.items[idx] += m.operationAmount
					}
				} else {
					if m.operationAmount < 0 {
						m.items[idx] = (m.items[idx] * m.items[idx]) % divisor
					} else {
						m.items[idx] = (m.items[idx] * m.operationAmount) % divisor
					}
				}
			}

			if m.items[idx]%m.test == 0 {
				monkies[m.testTrue].items = append(monkies[m.testTrue].items, m.items[idx])
			} else {
				monkies[m.testFalse].items = append(monkies[m.testFalse].items, m.items[idx])
			}

			m.timesInspected++
		}
		m.items = []int{}
	}
}

func day11p1() {
	input := readInput("input")
	monkies := parseNotesToMonkies(input)

	for i := 0; i < 20; i++ {
		runRound(monkies, true)
	}

	max1, max2 := 0, 0
	for _, m := range monkies {
		if m.timesInspected > max1 {
			max1 = m.timesInspected
		}
	}

	for _, m := range monkies {
		if m.timesInspected == max1 {
			continue
		}
		if m.timesInspected > max2 {
			max2 = m.timesInspected
		}
	}

	fmt.Printf("%d\n", max1*max2)
}

func day11p2() {
	input := readInput("input")
	monkies := parseNotesToMonkies(input)

	for i := 0; i < 10000; i++ {
		runRound(monkies, false)
	}

	max1, max2 := 0, 0
	for _, m := range monkies {
		if m.timesInspected > max1 {
			max1 = m.timesInspected
		}
	}

	for _, m := range monkies {
		if m.timesInspected == max1 {
			continue
		}
		if m.timesInspected > max2 {
			max2 = m.timesInspected
		}
	}

	fmt.Printf("%d\n", max1*max2)
}

func main() {
	day11p1()
	day11p2()
}
