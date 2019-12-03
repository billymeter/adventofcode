package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type op struct {
	val uint
}

type operation struct {
	gate   op
	input1 string
	input2 string
	output string
}

var (
	not    = op{0}
	and    = op{1}
	or     = op{2}
	lshift = op{3}
	rshift = op{4}
	assign = op{5}
)

func main() {
	opcodes := readInput("input.txt")
	var wires = make(map[string]int)
	for _, o := range opcodes {
		fmt.Println(o)
	}
}

func executeOperation(o operation, wires map[string]int) {
	v, err := strconv.Atoi(o.input1)
	if err != nil {
		i1int := false
		input1 := o.input1
	} else {
		i1int := true
		input1 := v
	}

	v, err = strconv.Atoi(o.input2)
	if err != nil {
		i2int := false
		input2 := o.input2
	} else {
		i2int := true
		input2 := v
	}

	switch o.gate {
	case not:
		if i1int {
			wires[o.output] = !input1
		} else {
			wires[o.output] = !wires[input1]
		}
	case and:
		if i1int && i2int {
			wires[o.output] = input1 & input2
		}
		if i1int {
			wires[o.output] = input1 & wires[input2]
		}
		if i2int {
			wires[o.output] = wires[input1] & input2
		} else {
			wires[o.output] = wires[input1] & wires[input2]
		}
	case or:
		if i1int && i2int {
			wires[o.output] = input1 | input2
		}
		if i1int {
			wires[o.output] = input1 | wires[input2]
		}
		if i2int {
			wires[o.output] = wires[input1] | input2
		} else {
			wires[o.output] = wires[input1] | wires[input2]
		}
	case lshift:
		if i1int && i2int {
			wires[o.output] = input1 << input2
		}
		if i1int {
			wires[o.output] = input1 << wires[input2]
		}
		if i2int {
			wires[o.output] = wires[input1] << input2
		} else {
			wires[o.output] = wires[input1] << wires[input2]
		}
	case rshift:
		if i1int && i2int {
			wires[o.output] = input1 >> input2
		}
		if i1int {
			wires[o.output] = input1 >> wires[input2]
		}
		if i2int {
			wires[o.output] = wires[input1] >> input2
		} else {
			wires[o.output] = wires[input1] >> wires[input2]
		}
	case assign:
		wires[o.output] = input1
	default:
		log.Fatal("invalid opcode for operation:", o)
	}
}

func parseLine(s string) operation {
	tokens := strings.Split(s, " ")
	if tokens[0] == "NOT" {
		return operation{not, tokens[1], "", tokens[3]}
	}
	switch tokens[1] {
	case "AND":
		return operation{and, tokens[0], tokens[2], tokens[4]}
	case "OR":
		return operation{or, tokens[0], tokens[2], tokens[4]}
	case "LSHIFT":
		return operation{and, tokens[0], tokens[2], tokens[4]}
	case "RSHIFT":
		return operation{and, tokens[0], tokens[2], tokens[4]}
	case "->":
		return operation{assign, tokens[0], "", tokens[2]}
	default:
		log.Fatal("Parse error for line:", tokens)
		return operation{op{99}, "", "", ""}
	}
}

func readInput(filename string) []operation {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	var opcodes []operation
	for _, s := range lines {
		opcodes = append(opcodes, parseLine(s))
	}
	return opcodes
}
