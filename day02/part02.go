package day02

import (
	"strconv"
	"strings"
)

func winGame(oppMove string, myScore int) (newScore int, myMove string) {
	newScore = myScore + 6
	switch oppMove {
	case "A":
		myMove = "B"
	case "B":
		myMove = "C"
	case "C":
		myMove = "A"
	}
	return
}

func loseGame(oppMove string, myScore int) (newScore int, myMove string) {
	newScore = myScore
	switch oppMove {
	case "A":
		myMove = "C"
	case "B":
		myMove = "A"
	case "C":
		myMove = "B"
	}
	return
}

func drawGame(oppMove string, myScore int) (newScore int, myMove string) {
	newScore = myScore + 3
	myMove = oppMove
	return
}

func Part2(input []string) string {
	var myScore = 0
	for _, v := range input {
		moves := strings.Split(v, " ")
		oppMove, result := moves[0], moves[1]
		var myMove = ""

		switch result {
		case "X":
			myScore, myMove = loseGame(oppMove, myScore)
		case "Y":
			myScore, myMove = drawGame(oppMove, myScore)
		case "Z":
			myScore, myMove = winGame(oppMove, myScore)
		}
		switch myMove {
		case "A":
			myScore += 1
		case "B":
			myScore += 2
		case "C":
			myScore += 3
		}
	}
	return strconv.Itoa(myScore)
}
