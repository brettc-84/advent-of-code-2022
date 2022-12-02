package day01

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
			"1000",
			"2000",
			"3000",
			"",
			"4000",
			"",
			"5000",
			"6000",
			"",
			"7000",
			"8000",
			"9000",
			"",
			"10000",
		},
			expected: 24000},
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
