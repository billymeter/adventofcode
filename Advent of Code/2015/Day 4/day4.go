package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
)

func main() {
	input := "yzbqklnj"

	fmt.Println("Advent of Code Day Four")
	p1 := mine(input, 5)
	fmt.Println("Part One:", p1)
	p2 := mine(input, 6)
	fmt.Println("Part One:", p2)
}

func mine(input string, difficulty int) int {
	for i := 0; ; i++ {
		h := md5.New()
		io.WriteString(h, input+strconv.Itoa(i))
		hash := hex.EncodeToString(h.Sum(nil))
		if difficulty == 5 {
			if isCoinMined(hash) {
				return i
			}
		} else {
			if isCoinMinedWithHigherDifficulty(hash) {
				return i
			}
		}
	}
}

func isCoinMined(hash string) bool {
	return hash[:5] == "00000"
}

func isCoinMinedWithHigherDifficulty(hash string) bool {
	return hash[:6] == "000000"
}
