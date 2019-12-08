package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	ip = 0
)

func main() {
	fmt.Println("Welcome to Thermal Environment Supervision Terminal (TEST) Computer")
	program := readFile("input.txt")
	// program := readFile("test.txt")
	for {
		width, err := processOpcode(program)
		if err != nil {
			log.Fatal(err)
		}
		if width < 0 {
			break
		}
		ip += width
	}
}

func processOpcode(mem []int) (int, error) {
	if ip > len(mem) {
		return -1, errors.New("Attempted to access out of bounds memory")
	}
	op := mem[ip] % 100
	op1Mode := mem[ip] % 1000 / 100
	op2Mode := mem[ip] % 10000 / 1000
	op3Mode := mem[ip] % 100000 / 10000
	op1, op2, op3 := 0, 0, 0

	if op == 99 {
		// halt
		return -1, nil
	}
	if op1Mode == 1 {
		op1 = mem[ip+1]
	} else {
		op1 = mem[mem[ip+1]]
	}

	if op != 3 && op != 4 {
		if op2Mode == 1 {
			op2 = mem[ip+2]
		} else {
			op2 = mem[mem[ip+2]]
		}

		if op3Mode == 1 {
			op3 = ip + 3
		} else {
			op3 = mem[ip+3]
		}
	}

	switch op {
	case 1:
		// add
		mem[op3] = op1 + op2
		return 4, nil
	case 2:
		// multiply
		mem[op3] = op1 * op2
		return 4, nil
	case 3:
		// input
		val := readLine()
		mem[mem[ip+1]] = val
		return 2, nil
	case 4:
		// output
		val := op1
		fmt.Printf("%v\n", val)
		return 2, nil
	case 5:
		// jump-if-true
		if op1 != 0 {
			ip = op2
			return 0, nil
		}
		return 3, nil
	case 6:
		// jump-if-false
		if op1 == 0 {
			ip = op2
			return 0, nil
		}
		return 3, nil
	case 7:
		// less than
		if op1 < op2 {
			mem[op3] = 1
		} else {
			mem[op3] = 0
		}
		return 4, nil
	case 8:
		// equals
		if op1 == op2 {
			mem[op3] = 1
		} else {
			mem[op3] = 0
		}
		return 4, nil
	// case 99:
	// 	// halt
	// 	fmt.Println("halt")
	// 	return -1, nil
	default:
		err := fmt.Errorf("opcode %v is not implemented", mem[ip])
		return -1, err
	}
}

func readLine() int {
	fmt.Println("Please enter an integer:")
	val := 0
	fmt.Scanf("%d", &val)
	return val
}

func readFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("cannot open file")
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	var codes []int
	for s.Scan() {
		line := s.Text()
		components := strings.Split(line, ",")
		for _, c := range components {
			v, _ := strconv.Atoi(c)
			codes = append(codes, v)
		}
	}
	return codes
}
