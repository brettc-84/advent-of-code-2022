package day10

import (
	"strconv"
	"strings"
)

func Part1(input []string) string {

	result := 0

	cycle := 1
	x := 1
	for _, instruction := range input {
		if instruction == "noop" {
			cycleCpu(&cycle, &x, 0, 1, &result)
		} else {
			instructionValue := strings.Split(instruction, " ")[1]
			value, _ := strconv.Atoi(instructionValue)
			cycleCpu(&cycle, &x, value, 2, &result)
		}
	}
	return strconv.Itoa(result)
}

func cycleCpu(cycle *int, x *int, operationValue int, numberOfCyclesToExecute int, result *int) {
	for i := 0; i < numberOfCyclesToExecute; i++ {
		if *cycle == 20 || *cycle == 60 || *cycle == 100 || *cycle == 140 || *cycle == 180 || *cycle == 220 {
			*result += *cycle * *x
		}
		*cycle += 1
	}
	*x += operationValue
}
