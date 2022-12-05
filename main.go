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
	fmt.Println("***** AoC 2022 *****")
	fmt.Println("***** Lets GOOO *****")

	seed := rand.NewSource((time.Now().UnixNano()))
	random := rand.New(seed)
	challenges := []challenge{
		{1, d1, colours[random.Intn(len(colours))]},
		{2, d2, colours[random.Intn(len(colours))]},
		{3, d3, colours[random.Intn(len(colours))]},
		{4, d4, colours[random.Intn(len(colours))]},
		{5, d5, colours[random.Intn(len(colours))]},
	}

	for _, dayChallenge := range challenges {
		fmt.Printf("%sRunning day %d...\n", dayChallenge.colour, dayChallenge.number)
		dayChallenge.work()
	}

}

func d1() {
	input := utils.ReadMultiLine("./day01/input.txt")
	fmt.Println("Part1:", day01.Part1(input))
	fmt.Println("Part2:", day01.Part2(input))
}

func d2() {
	input := utils.ReadMultiLine("./day02/input.txt")
	fmt.Println("Part1:", day02.Part1(input))
	fmt.Println("Part2:", day02.Part2(input))
}

func d3() {
	input := utils.ReadMultiLine("./day03/input.txt")
	fmt.Println("Part1:", day03.Part1(input))
	fmt.Println("Part2:", day03.Part2(input))

}

func d4() {
	input := utils.ReadMultiLine("./day04/input.txt")
	fmt.Println("Part1:", day04.Part1(input))
	fmt.Println("Part2:", day04.Part2(input))
}

func d5() {
	input := utils.ReadMultiLine("./day05/input.txt")
	fmt.Println("Part1:", day05.Part1(input))
	fmt.Println("Part2:", day05.Part2(input))
}
