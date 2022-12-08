package main

import (
	"fmt"
	"time"

	"github.com/brettc-84/advent-of-code-2022/day01"
	"github.com/brettc-84/advent-of-code-2022/day02"
	"github.com/brettc-84/advent-of-code-2022/day03"
	"github.com/brettc-84/advent-of-code-2022/day04"
	"github.com/brettc-84/advent-of-code-2022/day05"
	"github.com/brettc-84/advent-of-code-2022/day06"
	"github.com/brettc-84/advent-of-code-2022/day07"
	"github.com/brettc-84/advent-of-code-2022/day08"
	"github.com/brettc-84/advent-of-code-2022/utils"
)

const colorReset = "\033[0m"

type Day struct {
	number int
	input  []string
	part1  func(input []string) string
	part2  func(input []string) string
	colour string
}

func newDay(number int, part1 func([]string) string, part2 func([]string) string) *Day {

	filePath := fmt.Sprintf("./day%02d/input.txt", number)
	input := utils.ReadMultiLine(filePath)

	return &Day{number, input, part1, part2, utils.GetRandomColour()}
}

func (d *Day) Run() {
	start := time.Now()
	part1Result := d.part1(d.input)
	elapsed := time.Since(start)
	fmt.Printf("Part1 answer: %v done in %s\n", part1Result, elapsed)
	start = time.Now()
	part2Result := d.part2(d.input)
	elapsed = time.Since(start)
	fmt.Printf("Part2 answer: %v done in %s\n", part2Result, elapsed)
}

func main() {
	fmt.Printf("%s****** AoC 2022 ******\n", utils.GetRandomColour())
	fmt.Println("***** Lets GOOOO *****")

	completedChallenges := []Day{
		*newDay(1, day01.Part1, day01.Part2),
		*newDay(2, day02.Part1, day02.Part2),
		*newDay(3, day03.Part1, day03.Part2),
		*newDay(4, day04.Part1, day04.Part2),
		*newDay(5, day05.Part1, day05.Part2),
		*newDay(6, day06.Part1, day06.Part2),
		*newDay(7, day07.Part1, day07.Part2),
		*newDay(8, day08.Part1, day08.Part2),
	}

	for _, dayChallenge := range completedChallenges {
		fmt.Printf("%sRunning day %d...\n", dayChallenge.colour, dayChallenge.number)
		dayChallenge.Run()
	}
}
