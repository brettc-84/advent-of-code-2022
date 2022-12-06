package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/brettc-84/advent-of-code-2022/day01"
	"github.com/brettc-84/advent-of-code-2022/day02"
	"github.com/brettc-84/advent-of-code-2022/day03"
	"github.com/brettc-84/advent-of-code-2022/day04"
	"github.com/brettc-84/advent-of-code-2022/day05"
	"github.com/brettc-84/advent-of-code-2022/day06"
	"github.com/brettc-84/advent-of-code-2022/utils"
)

const colorReset = "\033[0m"

var colours = [...]string{
	"\033[31m",
	"\033[32m",
	"\033[33m",
	"\033[34m",
}

type challenge struct {
	number int
	work   func()
	colour string
}

func main() {
	fmt.Printf("%s****** AoC 2022 ******\n", colours[2])
	fmt.Println("***** Lets GOOOO *****")

	seed := rand.NewSource((time.Now().UnixNano()))
	random := rand.New(seed)
	challenges := []challenge{
		{1, d1, colours[random.Intn(len(colours))]},
		{2, d2, colours[random.Intn(len(colours))]},
		{3, d3, colours[random.Intn(len(colours))]},
		{4, d4, colours[random.Intn(len(colours))]},
		{5, d5, colours[random.Intn(len(colours))]},
		{6, d6, colours[random.Intn(len(colours))]},
	}

	for _, dayChallenge := range challenges {
		fmt.Printf("%sRunning day %d...\n", dayChallenge.colour, dayChallenge.number)
		dayChallenge.work()
	}
}

func d1() {
	input := utils.ReadMultiLine("./day01/input.txt")
	start := time.Now()
	part1 := day01.Part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part1 answer: %v done in %s\n", part1, elapsed)
	start = time.Now()
	part2 := day01.Part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part2 answer: %v done in %s\n", part2, elapsed)
}

func d2() {
	input := utils.ReadMultiLine("./day02/input.txt")
	start := time.Now()
	part1 := day02.Part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part1 answer: %v done in %s\n", part1, elapsed)
	start = time.Now()
	part2 := day02.Part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part2 answer: %v done in %s\n", part2, elapsed)
}

func d3() {
	input := utils.ReadMultiLine("./day03/input.txt")
	start := time.Now()
	part1 := day03.Part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part1 answer: %v done in %s\n", part1, elapsed)
	start = time.Now()
	part2 := day03.Part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part2 answer: %v done in %s\n", part2, elapsed)
}

func d4() {
	input := utils.ReadMultiLine("./day04/input.txt")
	start := time.Now()
	part1 := day04.Part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part1 answer: %v done in %s\n", part1, elapsed)
	start = time.Now()
	part2 := day04.Part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part2 answer: %v done in %s\n", part2, elapsed)
}

func d5() {
	input := utils.ReadMultiLine("./day05/input.txt")
	start := time.Now()
	part1 := day05.Part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part1 answer: %v done in %s\n", part1, elapsed)
	start = time.Now()
	part2 := day05.Part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part2 answer: %v done in %s\n", part2, elapsed)
}

func d6() {
	input := utils.ReadSingleLine("./day06/input.txt")
	start := time.Now()
	part1 := day06.Part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part1 answer: %v done in %s\n", part1, elapsed)
	start = time.Now()
	part2 := day06.Part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part2 answer: %v done in %s\n", part2, elapsed)
}
