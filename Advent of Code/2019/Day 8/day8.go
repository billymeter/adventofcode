package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	width  = 25
	height = 6
)

func main() {
	input := readFile("input.txt")
	img := loadData(input)
	// fmt.Println(img)
	fmt.Println("Advent of Code Day 8")
	p1 := partOne(img)
	fmt.Println("Part One:", p1)
}

func partOne(img []string) int {
	// Returns the layer number which has the most zeros.
	// Layer numbers start with layer 1 (not zero).
	var zerosInLayer []int
	for _, l := range img {
		fmt.Println(l)
		count := countCharacterInLayer(l, '0')
		zerosInLayer = append(zerosInLayer, count)
	}
	// fmt.Println("zeros in layer:", zerosInLayer)
	minZerosInLayer := findMix(zerosInLayer) + 1
	onesInLayer := countCharacterInLayer(img[minZerosInLayer], '1')
	twosInLayer := countCharacterInLayer(img[minZerosInLayer], '2')
	return onesInLayer * twosInLayer
}

func findMix(nums []int) int {
	index, min := 0, nums[0]
	for i, n := range nums {
		if n < min {
			min = n
			index = i
		}
	}
	return index
}

func countCharacterInLayer(layer string, char rune) int {
	count := 0
	for _, c := range layer {
		if c == char {
			count++
		}
	}
	return count
}

func loadData(d string) []string {
	layers := len(d) / width * height
	var img []string
	for l := 0; l < layers; l++ {
		var sb strings.Builder
		for wh := 0; wh < width*height; wh++ {
			sb.WriteByte(d[l+wh])
		}

		img = append(img, sb.String())
	}
	return img
}

func readFile(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("ya ded")
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Scan()
	return s.Text()
}
