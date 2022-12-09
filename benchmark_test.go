package main

import (
	"testing"

	"github.com/brettc-84/advent-of-code-2022/day01"
	"github.com/brettc-84/advent-of-code-2022/day02"
	"github.com/brettc-84/advent-of-code-2022/day03"
	"github.com/brettc-84/advent-of-code-2022/day04"
	"github.com/brettc-84/advent-of-code-2022/day05"
	"github.com/brettc-84/advent-of-code-2022/day06"
	"github.com/brettc-84/advent-of-code-2022/day07"
	"github.com/brettc-84/advent-of-code-2022/day08"
	"github.com/brettc-84/advent-of-code-2022/day09"
	"github.com/brettc-84/advent-of-code-2022/utils"
)

func BenchmarkDay01Part1(b *testing.B) {
	input := utils.ReadMultiLine("./day01/input.txt")
	for i := 0; i < b.N; i++ {
		day01.Part1(input)
	}
}
func BenchmarkDay01Part2(b *testing.B) {
	input := utils.ReadMultiLine("./day01/input.txt")
	for i := 0; i < b.N; i++ {
		day01.Part2(input)
	}
}
func BenchmarkDay02Part1(b *testing.B) {
	input := utils.ReadMultiLine("./day02/input.txt")
	for i := 0; i < b.N; i++ {
		day02.Part1(input)
	}
}
func BenchmarkDay02Part2(b *testing.B) {
	input := utils.ReadMultiLine("./day02/input.txt")
	for i := 0; i < b.N; i++ {
		day02.Part2(input)
	}
}
func BenchmarkDay03Part1(b *testing.B) {
	input := utils.ReadMultiLine("./day03/input.txt")
	for i := 0; i < b.N; i++ {
		day03.Part1(input)
	}
}
func BenchmarkDay03Part2(b *testing.B) {
	input := utils.ReadMultiLine("./day03/input.txt")
	for i := 0; i < b.N; i++ {
		day03.Part2(input)
	}
}
func BenchmarkDay04Part1(b *testing.B) {
	input := utils.ReadMultiLine("./day04/input.txt")
	for i := 0; i < b.N; i++ {
		day04.Part1(input)
	}
}
func BenchmarkDay04Part2(b *testing.B) {
	input := utils.ReadMultiLine("./day04/input.txt")
	for i := 0; i < b.N; i++ {
		day04.Part2(input)
	}
}
func BenchmarkDay05Part1(b *testing.B) {
	input := utils.ReadMultiLine("./day05/input.txt")
	for i := 0; i < b.N; i++ {
		day05.Part1(input)
	}
}
func BenchmarkDay05Part2(b *testing.B) {
	input := utils.ReadMultiLine("./day05/input.txt")
	for i := 0; i < b.N; i++ {
		day05.Part2(input)
	}
}
func BenchmarkDay06Part1(b *testing.B) {
	input := utils.ReadMultiLine("./day06/input.txt")
	for i := 0; i < b.N; i++ {
		day06.Part1(input)
	}
}
func BenchmarkDay06Part2(b *testing.B) {
	input := utils.ReadMultiLine("./day06/input.txt")
	for i := 0; i < b.N; i++ {
		day06.Part2(input)
	}
}
func BenchmarkDay07Part1(b *testing.B) {
	input := utils.ReadMultiLine("./day07/input.txt")
	for i := 0; i < b.N; i++ {
		day07.Part1(input)
	}
}
func BenchmarkDay07Part2(b *testing.B) {
	input := utils.ReadMultiLine("./day07/input.txt")
	for i := 0; i < b.N; i++ {
		day07.Part2(input)
	}
}
func BenchmarkDay08Part1(b *testing.B) {
	input := utils.ReadMultiLine("./day08/input.txt")
	for i := 0; i < b.N; i++ {
		day08.Part1(input)
	}
}
func BenchmarkDay08Part2(b *testing.B) {
	input := utils.ReadMultiLine("./day08/input.txt")
	for i := 0; i < b.N; i++ {
		day08.Part2(input)
	}
}
func BenchmarkDay09Part1(b *testing.B) {
	input := utils.ReadMultiLine("./day09/input.txt")
	for i := 0; i < b.N; i++ {
		day09.Part1(input)
	}
}
func BenchmarkDay09Part2(b *testing.B) {
	input := utils.ReadMultiLine("./day09/input.txt")
	for i := 0; i < b.N; i++ {
		day09.Part2(input)
	}
}
