package day09

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
			"R 4",
			"U 4",
			"L 3",
			"D 1",
			"R 4",
			"D 1",
			"L 5",
			"R 2",
		},
			expected: "1"},
		{input: []string{
			"R 5",
			"U 8",
			"L 8",
			"D 3",
			"R 17",
			"D 10",
			"L 25",
			"U 20",
		},
			expected: "36"},
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
