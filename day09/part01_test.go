package day09

import (
	"fmt"
	"testing"
)

func TestAreTouching(t *testing.T) {
	var tests = []struct {
		position                   string
		headX, headY, tailX, tailY int
		expected                   bool
	}{
		{"overlapped", 0, 0, 0, 0, true},
		{"directly above", 0, 1, 0, 0, true},
		{"directly below", 0, -1, 0, 0, true},
		{"directly right", 1, 0, 0, 0, true},
		{"directly left", -1, 0, 0, 0, true},
		{"diagonal up left", -1, 1, 0, 0, true},
		{"diagonal up right", 1, 1, 0, 0, true},
		{"diagonal down right", 1, -1, 0, 0, true},
		{"diagonal down left", -1, -1, 0, 0, true},
		{"two away above", 0, 2, 0, 0, false},
		{"two away below", 0, -2, 0, 0, false},
		{"two away right", 2, 0, 0, 0, false},
		{"two away left", -2, 0, 0, 0, false},
		{"two away up and left", -1, 2, 0, 0, false},
		{"random far away", -4, -1, 0, 0, false},
	}
	for _, testCase := range tests {
		testName := fmt.Sprintf("Testing %v", testCase.position)
		t.Run(testName, func(t *testing.T) {
			actual := areTouching(testCase.headX, testCase.headY, testCase.tailX, testCase.tailY)
			if actual != testCase.expected {
				t.Errorf("Expected %v but got %v for position %v", testCase.expected, actual, testCase.position)
			}
		})
	}
}

func TestPart1(t *testing.T) {
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
			expected: "13"},
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
