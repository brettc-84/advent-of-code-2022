package day14

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
			"498,4 -> 498,6 -> 496,6",
			"503,4 -> 502,4 -> 502,9 -> 494,9",
		},
			expected: "24"},
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
