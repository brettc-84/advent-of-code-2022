package day04

import (
	"strconv"
	"strings"
)

func Part2(input []string) string {
	result := 0
	for _, pairs := range input {
		assignments := strings.Split(pairs, ",")
		elf1, elf2 := assignments[0], assignments[1]
		elf1Start, elf1End := getPoints(elf1)
		elf2Start, elf2End := getPoints(elf2)

		if (elf1Start >= elf2Start && elf1End <= elf2End) ||
			(elf2Start >= elf1Start && elf2End <= elf1End) ||
			(elf1End >= elf2Start && elf1End <= elf2End) ||
			(elf2End >= elf1Start && elf2End <= elf1End) ||
			(elf1Start >= elf2Start && elf1Start <= elf2End) ||
			(elf2Start >= elf1Start && elf2Start <= elf1End) {
			result += 1
		}
	}
	return strconv.Itoa(result)
}
