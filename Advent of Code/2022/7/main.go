package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Item uint8

const (
	FILE Item = iota
	DIR
)

type Node struct {
	name     string
	kind     Item
	size     uint
	parent   *Node
	children []*Node
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

func buildFileSystem(lines []string) *Node {
	root := &Node{"/", DIR, 0, nil, nil}
	currNode := root

	for _, line := range lines {
		components := strings.Split(line, " ")
		if components[0] == "$" {
			switch components[1] {
			case "cd":
				if components[2] == "/" {
					currNode = root
					continue
				}

				if components[2] == ".." {
					currNode = currNode.parent
					continue
				}

				for _, node := range currNode.children {
					if node.name == components[2] {
						currNode = node
						break
					}
				}
				continue

			case "ls":
				continue
				// not needed?
			}
		} else if components[0] == "dir" {
			node := &Node{components[1], DIR, 0, currNode, nil}
			currNode.children = append(currNode.children, node)
		} else {
			size, err := strconv.Atoi(components[0])
			if err != nil {
				log.Fatal(err)
			}
			node := &Node{components[1], FILE, uint(size), currNode, nil}
			currNode.children = append(currNode.children, node)
		}
	}
	return root
}

func printFileSystem(node *Node, level int) {
	for i := 0; i < level; i++ {
		fmt.Printf("  ")
	}
	if node.kind == FILE {
		fmt.Printf("- %s (file, size=%d)\n", node.name, node.size)
	} else {
		fmt.Printf("- %s (dir, content size=%d)\n", node.name, node.size)
	}

	for _, n := range node.children {
		printFileSystem(n, level+1)
	}
}

func calculateDirectorySizes(node *Node) {
	for _, n := range node.children {
		calculateDirectorySizes(n)
	}

	if node.parent != nil {
		node.parent.size += node.size
	}
}

func sumOfDirectoriesWithContentsOfFilesizesLessThan100000(node *Node) uint {
	var sum uint = 0
	for _, n := range node.children {
		if n.kind == DIR {
			sum += sumOfDirectoriesWithContentsOfFilesizesLessThan100000(n)
		}
	}

	if node.kind == DIR && node.size <= 100000 {
		sum += node.size
	}
	return sum
}

func deletionCandidates(node *Node, spaceNeeded uint, candidates *[]*Node) {
	if node.kind == DIR {
		if node.size >= spaceNeeded {
			*candidates = append(*candidates, node)
		}

		for _, n := range node.children {
			if n.kind == DIR {
				deletionCandidates(n, spaceNeeded, candidates)
			} else {
				continue
			}
		}
	}
}

func day7p1() {
	input := readInput("input")
	root := buildFileSystem(input)
	calculateDirectorySizes(root)
	fmt.Printf("%d\n", sumOfDirectoriesWithContentsOfFilesizesLessThan100000(root))
}

func day7p2() {
	input := readInput("input")
	root := buildFileSystem(input)
	calculateDirectorySizes(root)

	freeSpace := 70000000 - root.size

	var candidates []*Node
	deletionCandidates(root, (30000000 - freeSpace), &candidates)

	var min uint = 70000000
	for _, n := range candidates {
		if n.size < min {
			min = n.size
		}
	}
	fmt.Printf("%d\n", min)
}

func main() {
	day7p1()
	day7p2()
}
