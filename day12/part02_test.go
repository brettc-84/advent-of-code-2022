package day12

import (
	"fmt"
	"testing"
)

func TestPart2(t *testing.T) {

	var tests = []struct {
		input    []string
		expected string
	}{
		{input: []string{
			"Sabqponm",
			"abcryxxl",
			"accszExk",
			"acctuvwj",
			"abdefghi",
		},
			expected: "29"},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("Expect %v", testCase.expected)
		t.Run(testName, func(t *testing.T) {
			actual := Part2(testCase.input)
			if actual != testCase.expected {
				t.Errorf("Expected: %v, actual: %v", testCase.expected, actual)
			}
		})
	}
}
