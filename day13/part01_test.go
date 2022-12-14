package day13

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
			"[1,1,3,1,1]",
			"[1,1,5,1,1]",
		},
			expected: "1"},
		{input: []string{
			"[[1],[2,3,4]]",
			"[[1],4]",
		},
			expected: "1"},
		// {input: []string{
		// 	"[1,1,3,1,1]",
		// 	"[1,1,5,1,1]",
		// 	"",
		// 	"[[1],[2,3,4]]",
		// 	"[[1],4]",
		// 	"",
		// 	"[9]",
		// 	"[[8,7,6]]",
		// 	"",
		// 	"[[4,4],4,4]",
		// 	"[[4,4],4,4,4]",
		// 	"",
		// 	"[7,7,7,7]",
		// 	"[7,7,7]",
		// 	"",
		// 	"[]",
		// 	"[3]",
		// 	"",
		// 	"[[[]]]",
		// 	"[[]]",
		// 	"",
		// 	"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		// 	"[1,[2,[3,[4,[5,6,0]]]],8,9]",
		// },
		// 	expected: "13"},
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
