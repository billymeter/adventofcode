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
	start  Node
	end    Node
	weight int
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
			// if elevation == 0 || elevation == 1 || elevation == -1 {
			// 	_edge := Edge{mapping[e.start], mapping[e.end], elevation}
			// 	_edges = append(_edges, _edge)
			// }
			elevation += mapping[e.start].elevation
			elevation = 1
			_edge := Edge{mapping[e.start], mapping[e.end], elevation}
			_edges = append(_edges, _edge)
		}
		graph[node] = _edges
	}

	start = mapping[s]
	end = mapping[e]

	return graph, start, end
}

func dijkstra(graph map[Node][]Edge, start, target Node) (path []Node) {
	q := list.New()
	dist := make(map[Node]int)
	prev := make(map[Node]Node)

	for n := range graph {
		dist[n] = 999999999999999999
		q.PushBack(n)
	}

	dist[start] = 0

	for n := q.Front(); n != nil; n = q.Front() {
		q.Remove(n)
		node := Node(n.Value.(Node))

		for _, e := range graph[node] {
			neighbor := e.end
			val := dist[node] + e.weight
			if val < dist[neighbor] {
				dist[neighbor] = val
				prev[neighbor] = node
			}
		}
	}

	fmt.Printf("%v\n\n", dist)
	fmt.Printf("%v\n", prev)
	return path
}

func day12p1() {
	input := readInput("test")
	graph, start, end := inputToGraph(input)

	// fmt.Printf("start: %v, end: %v\n", start, end)
	dijkstra(graph, start, end)
	// for n := range graph {
	// 	fmt.Printf("%v: %v\n", n, graph[n])
	// }
}

func day12p2() {
	// input := readInput("test")
	// fmt.Printf("%v\n", input)
}

func main() {
	day12p1()
	day12p2()
}
