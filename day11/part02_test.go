package day11

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
			"Monkey 0:",
			"	Starting items: 79, 98",
			"	Operation: new = old * 19",
			"	Test: divisible by 23",
			"	  If true: throw to monkey 2",
			"	  If false: throw to monkey 3",
			"",
			"Monkey 1:",
			"	Starting items: 54, 65, 75, 74",
			"	Operation: new = old + 6",
			"	Test: divisible by 19",
			"	  If true: throw to monkey 2",
			"	  If false: throw to monkey 0",
			"",
			"Monkey 2:",
			"	Starting items: 79, 60, 97",
			"	Operation: new = old * old",
			"	Test: divisible by 13",
			"	  If true: throw to monkey 1",
			"	  If false: throw to monkey 3",
			"",
			"Monkey 3:",
			"	Starting items: 74",
			"	Operation: new = old + 3",
			"	Test: divisible by 17",
			"	  If true: throw to monkey 0",
			"	  If false: throw to monkey 1",
		},
			expected: "2713310158"},
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
