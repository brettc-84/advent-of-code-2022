package day06

import (
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

func Part1(input string) int {
	start := 0
	for i := 4; i < len(input); i++ {
		mark := input[start:i]
		if hasUnique(mark) {
			return i
		}
		start += 1
	}
	return -1
}
