package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func Part2(input []string) string {

	var crt = make([][]rune, 6)
	for i := range crt {
		crt[i] = make([]rune, 40)
	}

	cycle := 1
	x := 1

	for _, instruction := range input {
		if instruction == "noop" {
			cycleAndDraw(crt, &cycle, 1, &x, 0)
		} else {
			instructionValue := strings.Split(instruction, " ")[1]
			value, _ := strconv.Atoi(instructionValue)
			cycleAndDraw(crt, &cycle, 2, &x, value)
		}
	}
	for _, row := range crt {
		s := string(row)
		fmt.Println(s)
	}
	return "^^ Look above ^^"
}

func cycleAndDraw(crt [][]rune, cycle *int, numberOfCyclesToExecute int, x *int, opValue int) {
	for i := 0; i < numberOfCyclesToExecute; i++ {
		crtRow := 0
		drawPos := *cycle - 1
		if *cycle >= 41 && *cycle <= 80 {
			crtRow = 1
			drawPos = *cycle - 41
		} else if *cycle >= 81 && *cycle <= 120 {
			crtRow = 2
			drawPos = *cycle - 81
		} else if *cycle >= 121 && *cycle <= 160 {
			crtRow = 3
			drawPos = *cycle - 121
		} else if *cycle >= 161 && *cycle <= 200 {
			crtRow = 4
			drawPos = *cycle - 161
		} else if *cycle >= 201 && *cycle <= 240 {
			crtRow = 5
			drawPos = *cycle - 201
		}

		spriteStart, spriteEnd := *x-1, *x+1

		if drawPos >= spriteStart && drawPos <= spriteEnd {
			crt[crtRow][drawPos] = '#'
		} else {
			crt[crtRow][drawPos] = '.'
		}

		*cycle += 1
	}
	*x += opValue
}
