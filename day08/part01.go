package day08

import "strconv"

func Part1(input []string) string {
	gridH, gridW := len(input), len(input[0])
	result := (gridW * 2) + ((gridH - 2) * 2)

	for x := 1; x < gridW-1; x++ {
		for y := 1; y < gridH-1; y++ {
			if isVisible(x, y, input) {
				result += 1
			}
		}
	}

	return strconv.Itoa(result)
}

func isVisible(x, y int, grid []string) bool {
	row := make([]int, 0)
	for _, t := range grid[y] {
		row = append(row, int(t-'0'))
	}
	tree := row[x]
	// from left
	visibleFromLeft := allShorter(row[0:x], tree)
	visibleFromRight := allShorter(row[x+1:], tree)

	// make array from column
	column := make([]int, 0)

	for i := 0; i < len(grid); i++ {
		column = append(column, int(grid[i][x]-'0'))
	}

	visibleFromTop := allShorter(column[0:y], tree)
	visibleFromBottom := allShorter(column[y+1:], tree)

	return visibleFromBottom || visibleFromLeft || visibleFromRight || visibleFromTop
}

func allShorter(trees []int, myTree int) bool {
	for i := 0; i < len(trees); i++ {
		if trees[i] >= myTree {
			return false
		}
	}
	return true
}
