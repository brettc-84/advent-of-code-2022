package main

import (
	"fmt"
	"strconv"

	"github.com/brettc-84/advent-of-code-2022/utils"
)

func main() {
	fmt.Println("***** AoC 2022 *****")
	fmt.Println("***** Lets GOOO *****")

	day01Part1()
	day01Part2()
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
