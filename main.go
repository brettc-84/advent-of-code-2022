package main

import (
	"fmt"
	"strings"

	"github.com/brettc-84/advent-of-code-2022/day01"
	"github.com/brettc-84/advent-of-code-2022/utils"
)

type challenge struct {
	number int
	work   func()
}

func main() {
	fmt.Println("***** AoC 2022 *****")
	fmt.Println("***** Lets GOOO *****")

	challenges := []challenge{
		{1, d1},
	}

	for _, dayChallenge := range challenges {
		fmt.Printf("Running day %d...\n", dayChallenge.number)
		dayChallenge.work()
	}

	day02Part1()
	day02Part2()
}

func d1() {
	input := utils.ReadMultiLine("./day01/input.txt")
	fmt.Println("Part1:", day01.Part1(input))
	fmt.Println("Part2:", day01.Part2(input))
}

func day02Part1() {

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
	stratGuide := utils.ReadMultiLine("./day02/input.txt")

	// stratGuide := []string{"A Y", "B X", "C Z"}

	myScore := 0
	for _, v := range stratGuide {
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
	fmt.Println("My score:", myScore)
}

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

func day02Part2() {

	/*
		Rock (A) 1pts
		Paper (B) 2pts
		Scissors (C) 3pts

		X = lose
		Y = draw
		Z = win

		Win = 6 pts
		Draw = 3 pts
		Loss = 0 pts

		A > C
		B > A
		C > B
	*/
	stratGuide := utils.ReadMultiLine("./day02/input.txt")

	// stratGuide := []string{"A Y", "B X", "C Z"}

	var myScore = 0
	for _, v := range stratGuide {
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
	fmt.Println("My score:", myScore)
}
