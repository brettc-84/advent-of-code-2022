package day06

import "strconv"

func Part2(input []string) string {
	line := input[0]
	start := 0
	for i := 14; i < len(line); i++ {
		mark := line[start:i]
		if hasUnique(mark) {
			return strconv.Itoa(i)
		}
		start += 1
	}
	return ""
}
