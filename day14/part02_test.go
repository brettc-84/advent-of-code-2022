package day14

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
			"498,4 -> 498,6 -> 496,6",
			"503,4 -> 502,4 -> 502,9 -> 494,9",
		},
			expected: "93"},
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
