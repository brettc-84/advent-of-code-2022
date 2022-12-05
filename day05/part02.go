package day05

import (
	"regexp"
	"strconv"
	"strings"
)

func Part2(input []string) string {
	result := ""

	stacks := [][]string{}
	stackNumber := 0

	stackLines := 0
	crateRegex := regexp.MustCompile(`\[|\]`)
	// create stacks
	for lineNum, line := range input {
		if !strings.Contains(line, "[") {
			stackLines = lineNum + 2
			break
		}
		for i := 0; i < len(line)-2; i += 4 {
			if len(stacks) <= stackNumber {
				stacks = append(stacks, make([]string, 0))
			}
			potentialCrate := crateRegex.ReplaceAllString(line[i:i+3], "")
			if strings.Trim(potentialCrate, " ") != "" {
				stacks[stackNumber] = append(stacks[stackNumber], potentialCrate)
			}
			stackNumber += 1
		}
		stackNumber = 0
	}

	input = input[stackLines:]

	// handle moves
	var movePattern = regexp.MustCompile(`move\s(\d+)\sfrom\s(\d+)\sto\s(\d+)`)
	for _, moveString := range input {
		moveSet := movePattern.FindAllStringSubmatch(moveString, -1)
		count, _ := strconv.Atoi(moveSet[0][1])
		from, _ := strconv.Atoi(moveSet[0][2])
		to, _ := strconv.Atoi(moveSet[0][3])
		from -= 1
		to -= 1
		var things = make([]string, count)
		copy(things, stacks[from][0:count])
		stacks[from] = stacks[from][count:len(stacks[from])]
		stacks[to] = append(things, stacks[to]...)
	}

	for _, stack := range stacks {
		result += stack[0]
	}

	return result
}
