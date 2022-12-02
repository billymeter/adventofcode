package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Move int
type Result int

const (
	Rock     Move = 1
	Paper         = 2
	Scissors      = 3
)

const (
	Lose Result = 0
	Draw        = 3
	Win         = 6
)

func readInput(filename string) ([]Move, []Move, []Move) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var enemy, player, strat []Move

	for scanner.Scan() {
		t := scanner.Text()
		moves := strings.Split(t, " ")
		var enemyMove Move
		switch moves[0] {
		case "A":
			enemy = append(enemy, Rock)
			enemyMove = Rock
		case "B":
			enemy = append(enemy, Paper)
			enemyMove = Paper
		case "C":
			enemy = append(enemy, Scissors)
			enemyMove = Scissors
		default:
			log.Print("unknown opponent value")
		}

		switch moves[1] {
		case "X":
			player = append(player, Rock)
		case "Y":
			player = append(player, Paper)
		case "Z":
			player = append(player, Scissors)
		default:
			log.Print("unknown player value")
		}

		switch moves[1] {
		case "X": // lose
			if enemyMove == Rock {
				strat = append(strat, Scissors)
			} else if enemyMove == Paper {
				strat = append(strat, Rock)
			} else {
				strat = append(strat, Paper)
			}
		case "Y": // draw
			if enemyMove == Rock {
				strat = append(strat, Rock)
			} else if enemyMove == Paper {
				strat = append(strat, Paper)
			} else {
				strat = append(strat, Scissors)
			}
		case "Z": // win
			if enemyMove == Rock {
				strat = append(strat, Paper)
			} else if enemyMove == Paper {
				strat = append(strat, Scissors)
			} else {
				strat = append(strat, Rock)
			}
		default:
			log.Print("unknown player value")
		}
	}

	return enemy, player, strat
}

// func convertMovesToStrategies(moves []Move) []Result {
// 	var result []Result
// 	for _, move := range moves {
// 		switch move {
// 		// X was set to Rock, but this should mean lose
// 		// Y was set to Paper, but this should mean draw
// 		// Z was set to Scissors, but this should mean win
// 		case Rock:
// 			result = append(result, Lose)
// 		case Paper:
// 			result = append(result, Draw)
// 		case Scissors:
// 			result = append(result, Win)
// 		}
// 	}
// 	return result
// }

func playerPointTotalForRound(enemy, player Move) int {
	score := int(player)

	switch enemy {
	case Rock:
		if player == Rock {
			score += int(Draw)
		}
		if player == Paper {
			score += int(Win)
		}
		if player == Scissors {
			score += int(Lose)
		}
	case Paper:
		if player == Rock {
			score += int(Lose)
		}
		if player == Paper {
			score += int(Draw)
		}
		if player == Scissors {
			score += int(Win)
		}
	case Scissors:
		if player == Rock {
			score += int(Win)
		}
		if player == Paper {
			score += int(Lose)
		}
		if player == Scissors {
			score += int(Draw)
		}
	}
	return score
}

func day2p1(enemy, player []Move) int {
	score := 0
	length := len(player)
	for i := 0; i < length; i++ {
		// log.Printf("enemy: %d\tplayer: %d\tplayer points: %d\n", enemy[i], player[i], playerPointTotalForRound(enemy[i], player[i]))
		score += playerPointTotalForRound(enemy[i], player[i])
	}

	return score
}

func day2p2(enemy, strat []Move) int {
	score := 0
	for i := range enemy {
		score += playerPointTotalForRound(enemy[i], strat[i])
	}
	return score
}

func main() {
	enemy, player, strats := readInput("input")
	println(day2p1(enemy, player))
	println(day2p2(enemy, strats))

}
