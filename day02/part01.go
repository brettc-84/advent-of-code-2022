package day02

import (
	"strings"
)

func Part1(input []string) int {

	/*
		Rock (A, X) 1pts
		Paper (B, Y) 2pts
		Scissors (C, Z) 3pts

		Win = 6 pts
		Draw = 3 pts
		Loss = 0 pts

		A, X > C, Z
		B, Y > A, X
		C, Z > B, Y
	*/

	// stratGuide := []string{"A Y", "B X", "C Z"}

	myScore := 0
	for _, v := range input {
		moves := strings.Split(v, " ")
		oppMove, myMove := moves[0], moves[1]

		switch myMove {
		case "X":
			myScore += 1
			if oppMove == "C" {
				// I win
				myScore += 6
			} else if oppMove == "A" {
				myScore += 3
			}
		case "Y":
			myScore += 2
			if oppMove == "A" {
				// I win
				myScore += 6
			} else if oppMove == "B" {
				myScore += 3
			}
		case "Z":
			myScore += 3
			if oppMove == "B" {
				// I win
				myScore += 6
			} else if oppMove == "C" {
				myScore += 3
			}
		}
	}
	return myScore
}
