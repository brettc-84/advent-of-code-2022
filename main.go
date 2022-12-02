package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brettc-84/advent-of-code-2022/utils"
)

func main() {
	fmt.Println("***** AoC 2022 *****")
	fmt.Println("***** Lets GOOO *****")

	day02Part1()
	day02Part2()
}

func day01Part1() {
	allCalories := utils.ReadMultiLine("./day01/input.txt")

	var max = 0

	var elfCalories = 0

	for i := 0; i < len(allCalories); i++ {
		calValue, _ := strconv.Atoi(allCalories[i])
		if calValue > 0 {
			elfCalories += calValue
		} else {
			if elfCalories > max {
				max = elfCalories
			}
			elfCalories = 0
		}
	}
	fmt.Println("Part1:", max)
}

func arrangeTop3(top3 *[3]int, value int) bool {
	var changed = false
	if value >= top3[0] {
		top3[2] = top3[1]
		top3[1] = top3[0]
		top3[0] = value
		changed = true
	} else if value >= top3[1] {
		top3[2] = top3[1]
		top3[1] = value
		changed = true
	} else if value >= top3[2] {
		top3[2] = value
		changed = true
	}
	return changed
}

func sumArray(arr [3]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func day01Part2() {
	allCalories := utils.ReadMultiLine("./day01/input.txt")
	var top3 [3]int

	var elfCalories = 0

	for i := 0; i < len(allCalories); i++ {
		calValue, _ := strconv.Atoi(allCalories[i])
		if calValue > 0 {
			elfCalories += calValue
		} else {
			arrangeTop3(&top3, elfCalories)
			elfCalories = 0
		}
	}
	fmt.Println("Part2:", sumArray(top3))
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
		moves := strings.Split(v," ")
		oppMove, myMove := moves[0], moves[1]

		switch myMove {
			case "X":
				myScore += 1
				if (oppMove == "C") {
					// I win
					myScore += 6
				} else if (oppMove == "A") {
					myScore += 3
				}
			case "Y":
				myScore += 2
				if (oppMove == "A") {
					// I win
					myScore += 6
				} else if (oppMove == "B") {
					myScore += 3
				}
			case "Z":
				myScore += 3
				if (oppMove == "B") {
					// I win
					myScore += 6
				} else if (oppMove == "C") {
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
		moves := strings.Split(v," ")
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