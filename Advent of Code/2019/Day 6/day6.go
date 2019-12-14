package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	name     string
	children []node
}

func main() {
	fmt.Println("Advent of Code 2019 Day 6")
	orbits := readFile("input.txt")
	fmt.Println(orbits)
	fmt.Println(orbits["COM"])
}

func readFile(f string) map[string][]string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal("ded")
	}
	defer file.Close()
	orbits := make(map[string][]string)
	s := bufio.NewScanner(file)
	for s.Scan() {
		l := s.Text()
		planets := strings.Split(l, ")")
		p1, p2 := planets[0], planets[1]
		orbits[p1] = append(orbits[p1], p2)
	}
	return orbits
}
