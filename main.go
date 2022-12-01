package main

import (
	"fmt"
	"strconv"

	"github.com/brettc-84/advent-of-code-2022/utils"
)

func main() {
	fmt.Println("***** AoC 2022 *****")
	fmt.Println("***** Lets GOOO *****")

	day01()
}

func day01() {
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
