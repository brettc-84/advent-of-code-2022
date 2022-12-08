package day08

import "strconv"

func Part2(input []string) string {
	gridH, gridW := len(input), len(input[0])
	result := 0

	for x := 1; x < gridW-1; x++ {
		for y := 1; y < gridH-1; y++ {
			treeScore := getScenicScoreForTree(x, y, input)
			if treeScore > result {
				result = treeScore
			}
		}
	}

	return strconv.Itoa(result)
}

func getScenicScoreForTree(x, y int, grid []string) int {

	score := 0
	row := make([]int, 0)
	for _, t := range grid[y] {
		row = append(row, int(t-'0'))
	}
	column := make([]int, 0)

	for i := 0; i < len(grid); i++ {
		column = append(column, int(grid[i][x]-'0'))
	}
	tree := row[x]

	// up
	scoreUp := calculateDownScore(tree, column[0:y])
	// left
	scoreLeft := calculateDownScore(tree, row[0:x])
	// right
	scoreRight := calculateUpScore(tree, row[x+1:])
	// down
	scoreDown := calculateUpScore(tree, column[y+1:])

	score = scoreUp * scoreLeft * scoreRight * scoreDown
	return score
}

func calculateUpScore(tree int, otherTrees []int) int {
	score := 0
	for i := 0; i < len(otherTrees); i++ {
		score += 1
		if otherTrees[i] >= tree {
			break
		}
	}
	return score
}

func calculateDownScore(tree int, otherTrees []int) int {
	score := 0
	for i := len(otherTrees) - 1; i >= 0; i-- {
		score += 1
		if otherTrees[i] >= tree {
			break
		}
	}
	return score
}
