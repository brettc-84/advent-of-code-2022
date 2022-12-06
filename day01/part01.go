package day01

import (
	"strconv"
)

func Part1(input []string) string {

	var max = 0

	var elfCalories = 0

	for i := 0; i < len(input); i++ {
		calValue, _ := strconv.Atoi(input[i])
		if calValue > 0 {
			elfCalories += calValue
		} else {
			if elfCalories > max {
				max = elfCalories
			}
			elfCalories = 0
		}
	}
	return strconv.Itoa(max)
}
