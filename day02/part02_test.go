package day02

import (
	"fmt"
	"testing"
)

func TestPart2(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{input: []string{
			"A Y",
			"B X",
			"C Z",
		},
			expected: 12},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("Expect %d", testCase.expected)
		t.Run(testName, func(t *testing.T) {
			actual := Part2(testCase.input)
			if actual != testCase.expected {
				t.Errorf("Expected: %d, actual: %d", testCase.expected, actual)
			}
		})
	}

}
