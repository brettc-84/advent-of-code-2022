package day06

import (
	"strconv"
	"strings"
)

func hasUnique(value string) bool {
	for _, c := range value {
		if strings.Count(value, string(c)) > 1 {
			return false
		}
	}
	return true
}

func Part1(input []string) string {
	line := input[0]
	start := 0
	for i := 4; i < len(line); i++ {
		mark := line[start:i]
		if hasUnique(mark) {
			return strconv.Itoa(i)
		}
		start += 1
	}
	return ""
}
