package day12

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {

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
			expected: "31"},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("Expect %v", testCase.expected)
		t.Run(testName, func(t *testing.T) {
			actual := Part1(testCase.input)
			if actual != testCase.expected {
				t.Errorf("Expected: %v, actual: %v", testCase.expected, actual)
			}
		})
	}
}
