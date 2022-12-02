package day02

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{input: []string{
			"A Y",
			"B X",
			"C Z",
		},
			expected: 15},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("Expect %d", testCase.expected)
		t.Run(testName, func(t *testing.T) {
			actual := Part1(testCase.input)
			if actual != testCase.expected {
				t.Errorf("Expected: %d, actual: %d", testCase.expected, actual)
			}
		})
	}

}
