package main

import (
	"fmt"
	"io"
	"os"
)

func findStartOfPacket(filename string, uniqueLength int) {
	file, _ := os.Open(filename)
	sequence, _ := io.ReadAll(file)
	file.Close()

	var i, j int

	for i = 0; i < len(sequence)-uniqueLength; i++ {
		seen := make(map[byte]struct{})
		seen[sequence[i]] = struct{}{}
		for j = i + 1; j <= i+uniqueLength; j++ {
			if len(seen) == uniqueLength {
				fmt.Printf("%d\n", i+uniqueLength)
				return
			}
			if _, ok := seen[sequence[j]]; ok {
				break
			}
			seen[sequence[j]] = struct{}{}
		}
	}
}

func main() {
	findStartOfPacket("input", 4)
	findStartOfPacket("input", 14)
}
