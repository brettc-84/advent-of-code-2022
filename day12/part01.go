package day12

import (
	"strconv"
)

type Cell struct {
	x, y     int
	value    rune
	visited  bool
	distance int
}

func Part1(input []string) string {
	// result := 0

	possibilePathQueue := make([]Cell, 0)

	grid := make([][]Cell, len(input))
	var startCell Cell
	var endCell Cell

	// create grid and initial
	for y, row := range input {
		grid[y] = make([]Cell, len(row))
		for x, cell := range row {
			coord := Cell{
				x: x, y: y, value: cell, visited: false,
			}
			coord.distance = 0
			if cell == 'S' {
				coord.visited = true
				coord.value = 'a'
				startCell = coord
			} else if cell == 'E' {
				coord.value = 'z'
				endCell = coord
			}
			grid[y][x] = coord
		}
	}

	// find path
	possibilePathQueue = append(possibilePathQueue, startCell)

	for len(possibilePathQueue) > 0 {
		checkCell := possibilePathQueue[0]
		possibilePathQueue = possibilePathQueue[1:]
		if checkCell == endCell {
			return strconv.Itoa(checkCell.distance)
		}

		// check up
		if checkCell.y-1 >= 0 {
			upNeighbour := grid[checkCell.y-1][checkCell.x]
			if isValidStep(checkCell, upNeighbour) {
				if upNeighbour == endCell {
					return strconv.Itoa(checkCell.distance + 1)
				}
				upNeighbour.distance = checkCell.distance + 1
				upNeighbour.visited = true
				grid[upNeighbour.y][upNeighbour.x] = upNeighbour
				possibilePathQueue = append(possibilePathQueue, upNeighbour)
			}
		}
		// check down
		if checkCell.y+1 < len(grid) {
			downNeighbour := grid[checkCell.y+1][checkCell.x]
			if isValidStep(checkCell, downNeighbour) {
				if downNeighbour == endCell {
					return strconv.Itoa(checkCell.distance + 1)
				}
				downNeighbour.distance = checkCell.distance + 1
				downNeighbour.visited = true
				grid[downNeighbour.y][downNeighbour.x] = downNeighbour
				possibilePathQueue = append(possibilePathQueue, downNeighbour)
			}
		}
		// check left
		if checkCell.x-1 >= 0 {
			leftNeighbour := grid[checkCell.y][checkCell.x-1]
			if isValidStep(checkCell, leftNeighbour) {
				if leftNeighbour == endCell {
					return strconv.Itoa(checkCell.distance + 1)
				}
				leftNeighbour.distance = checkCell.distance + 1
				leftNeighbour.visited = true
				grid[leftNeighbour.y][leftNeighbour.x] = leftNeighbour
				possibilePathQueue = append(possibilePathQueue, leftNeighbour)
			}
		}
		// check right
		if checkCell.x+1 < len(grid[0]) {
			rightNeighbour := grid[checkCell.y][checkCell.x+1]
			if isValidStep(checkCell, rightNeighbour) {
				if rightNeighbour == endCell {
					return strconv.Itoa(checkCell.distance + 1)
				}
				rightNeighbour.distance = checkCell.distance + 1
				rightNeighbour.visited = true
				grid[rightNeighbour.y][rightNeighbour.x] = rightNeighbour
				possibilePathQueue = append(possibilePathQueue, rightNeighbour)
			}
		}
	}
	return strconv.Itoa(-1)
}

func isValidStep(from Cell, to Cell) bool {
	if !to.visited &&
		(from.value >= to.value-1) {
		return true
	}
	return false
}

func getAdjacentCells(cell Cell, grid [][]Cell) []Cell {
	adjacents := make([]Cell, 0)
	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1
	x := cell.x
	y := cell.y
	// up
	if y > 0 {
		adjacents = append(adjacents, grid[y-1][x])
	}
	// down
	if y < maxY {
		adjacents = append(adjacents, grid[y+1][x])
	}
	// left
	if x > 0 {
		adjacents = append(adjacents, grid[y][x-1])
	}
	// right
	if x < maxX {
		adjacents = append(adjacents, grid[y][x+1])
	}
	return adjacents
}
