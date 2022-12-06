package day01

import (
	"strconv"
)

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

func Part2(input []string) string {

	var top3 [3]int

	var elfCalories = 0

	for i := 0; i < len(input); i++ {
		calValue, _ := strconv.Atoi(input[i])
		elfCalories += calValue
		if calValue == 0 || i == len(input)-1 {
			arrangeTop3(&top3, elfCalories)
			elfCalories = 0
		}
	}
	return strconv.Itoa(sumArray(top3))
}
