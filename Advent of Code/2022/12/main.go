package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type Node struct {
	number    int
	elevation int
}

type Edge struct {
	start Node
	end   Node
}

type SuckyEdge struct {
	start int
	end   int
}

func readInput(filename string) []string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func getElevation(c rune) int {
	switch c {
	case 'S':
		return 0
	case 'E':
		return 25
	default:
		return int(c - 'a')
	}
}

func getEdges(node int, width, height int) []SuckyEdge {
	var edges []SuckyEdge

	up, right, down := node-width, node+1, node+width
	// edge to up node
	if up >= 0 {
		edges = append(edges, SuckyEdge{node, up})
	}
	// edge to right node
	if (right % width) != 0 {
		edges = append(edges, SuckyEdge{node, right})
	}

	// edge to bottom node
	if down < height*width {
		edges = append(edges, SuckyEdge{node, down})
	}

	// edge to left node
	if (node % width) != 0 {
		edges = append(edges, SuckyEdge{node, node - 1})
	}

	return edges
}

func inputToGraph(lines []string) (graph map[Node][]Edge, start, end Node) {
	graph = make(map[Node][]Edge)
	mapping := make(map[int]Node)
	nodeNum, s, e := 0, 0, 0

	temp := make(map[int][]SuckyEdge)
	width := len(lines[0])
	height := len(lines)

	for _, line := range lines {
		for _, c := range line {
			if c == 'S' {
				s = nodeNum
			}
			if c == 'E' {
				e = nodeNum
			}

			mapping[nodeNum] = Node{nodeNum, getElevation(c)}
			edges := getEdges(nodeNum, width, height)
			temp[nodeNum] = edges
			nodeNum++
		}
	}

	for n, edges := range temp {
		node := mapping[n]
		var _edges []Edge
		for _, e := range edges {
			elevation := mapping[e.end].elevation - mapping[e.start].elevation
			if elevation <= 0 || elevation == 1 {
				_edge := Edge{mapping[e.start], mapping[e.end]}
				_edges = append(_edges, _edge)
			}
		}
		graph[node] = _edges
	}

	start = mapping[s]
	end = mapping[e]

	return graph, start, end
}

func bfs(graph map[Node][]Edge, start Node) map[Node]int {
	q := list.New()
	dist := make(map[Node]int)
	for n := range graph {
		dist[n] = 999999999999999999
	}

	dist[start] = 0
	q.PushBack(start)
	for node := q.Front(); node != nil; node = q.Front() {
		n := node.Value.(Node)
		q.Remove(node)
		for _, e := range graph[n] {
			if dist[e.end] == 999999999999999999 {
				q.PushBack(e.end)
				dist[e.end] = dist[n] + 1
			}
		}
	}

	return dist
}

func day12p1() {
	input := readInput("input")
	graph, start, end := inputToGraph(input)

	dist := bfs(graph, start)
	fmt.Printf("%d\n", dist[end])
}

func day12p2() {
	// this is not efficient
	// something like Floyd-Warshall might be better, but I'm not implementing
	// another graph algo
	input := readInput("input")
	graph, _, end := inputToGraph(input)

	min := 999999999999999999
	for n := range graph {
		if n.elevation == 0 {
			dist := bfs(graph, n)
			if dist[end] < min {
				min = dist[end]
			}
		}
	}
	fmt.Printf("%d\n", min)
}

func main() {
	day12p1()
	day12p2()
}
