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

// using the stack implementation from https://towardsdev.com/implementing-stack-in-golang-30e4c4c9f941
// since i don't understand Go interfaces yet
type Stack struct {
	top        int
	capacity   int
	stackArray []interface{}
}

func (stack *Stack) init(capacity int) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.stackArray = make([]interface{}, capacity)

	return stack
}

func newStack(capacity int) *Stack {
	return new(Stack).init(capacity)
}

func (stack *Stack) push(data interface{}) error {
	if stack.isFull() {
		return errors.New("stack overflow")
	}
	stack.top++
	stack.stackArray[stack.top] = data
	return nil
}

func (stack *Stack) pop() (interface{}, error) {
	if stack.isEmpty() {
		return nil, errors.New("stack underflow")
	}
	temp := stack.stackArray[stack.top]
	stack.top--
	return temp, nil
}

func (stack *Stack) peek() (interface{}, error) {
	if stack.isEmpty() {
		return nil, errors.New("stack underflow")
	}
	temp := stack.stackArray[stack.top]
	return temp, nil
}

func (stack *Stack) isFull() bool {
	return stack.top == int(stack.capacity)-1
}

func (stack *Stack) isEmpty() bool {
	return stack.top == -1
}

// end copy pasta of stack implementation

func readInput(filename string) ([]Stack, []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var stacks []Stack
	var stackData []string
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	t := scanner.Text()

	numLines := 1
	numberOfStacks := len(t)/4 + 1
	stackData = append(stackData, t)

	// we need to count the number of lines in the stacks
	// so we can build our stack data structures
	for scanner.Scan() {
		t := scanner.Text()
		// this is ugly
		temp := strings.TrimSpace(t)[0]
		if temp == '1' {
			// trash newline then break
			scanner.Scan()
			break
		}
		numLines++
		stackData = append(stackData, t)
	}

	// build the stacks
	for i := 0; i < numberOfStacks; i++ {
		s := newStack(numLines * numLines)
		stacks = append(stacks, *s)
	}

	// populate the stacks
	for i := len(stackData) - 1; i >= 0; i-- {
		stackNum := 0
		for cursor := 1; cursor < numberOfStacks*4; cursor += 4 {
			letter := stackData[i][cursor]
			if letter != ' ' {
				stacks[stackNum].push(letter)
			}
			stackNum++
		}
	}

	// return the slice of operations
	var operations []string
	for scanner.Scan() {
		operations = append(operations, scanner.Text())
	}
	return stacks, operations
}

func day5p1() {
	stacks, ops := readInput("input")

	for _, op := range ops {
		steps := strings.Split(op, " ")
		quantity, _ := strconv.Atoi(steps[1])
		source, _ := strconv.Atoi(steps[3])
		destination, _ := strconv.Atoi(steps[5])
		// adjust source and destinations to start at 0
		source--
		destination--

		for i := quantity; i > 0; i-- {
			t, _ := stacks[source].pop()
			stacks[destination].push(t)
		}
	}

	for i := 0; i < len(stacks); i++ {
		t, _ := stacks[i].peek()
		fmt.Printf("%c", t)
	}
	fmt.Println()
}

func day5p2() {
	stacks, ops := readInput("input")

	for _, op := range ops {
		steps := strings.Split(op, " ")
		quantity, _ := strconv.Atoi(steps[1])
		source, _ := strconv.Atoi(steps[3])
		destination, _ := strconv.Atoi(steps[5])
		// adjust source and destinations to start at 0
		source--
		destination--

		l := len(stacks)
		temp := newStack(l * l)

		for i := quantity; i > 0; i-- {
			t, _ := stacks[source].pop()
			temp.push(t)
		}

		for i := 0; i < quantity; i++ {
			t, _ := temp.pop()
			stacks[destination].push(t)
		}
	}

	for i := 0; i < len(stacks); i++ {
		t, _ := stacks[i].peek()
		fmt.Printf("%c", t)
	}
	fmt.Println()
}

func main() {
	day5p1()
	day5p2()
}
