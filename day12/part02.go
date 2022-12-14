package day12

import (
	"math"
	"strconv"
)

func Part2(input []string) string {
	// result := 0

	grid := make([][]Cell, len(input))
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
				coord.value = 'a'
			} else if cell == 'E' {
				coord.value = 'z'
				endCell = coord
			}
			grid[y][x] = coord
		}
	}

	// get valid starting points
	startingPoints := make([]Cell, 0)
	for _, row := range grid {
		for _, cell := range row {
			if cell.value == 'a' && hasValidAdjacentCells(cell, grid) {
				startingPoints = append(startingPoints, cell)
			}
		}
	}

	// find path

	min := math.MaxInt

	for _, try := range startingPoints {
		thisGrid := make([][]Cell, len(grid))
		for i := range grid {
			thisGrid[i] = make([]Cell, len(grid[i]))
			copy(thisGrid[i], grid[i])
		}
		distance := calcDistance(try, endCell, thisGrid)
		if distance != -1 {
			if distance < min {
				min = distance
			}
		}
	}
	return strconv.Itoa(min)
}

func calcDistance(from Cell, destination Cell, grid [][]Cell) int {
	possibilePathQueue := make([]Cell, 0)
	possibilePathQueue = append(possibilePathQueue, from)
	for len(possibilePathQueue) > 0 {
		checkCell := possibilePathQueue[0]
		possibilePathQueue = possibilePathQueue[1:]
		if checkCell == destination {
			return checkCell.distance
		}

		// check up
		if checkCell.y-1 >= 0 {
			upNeighbour := grid[checkCell.y-1][checkCell.x]
			if isValidStep(checkCell, upNeighbour) {
				if upNeighbour == destination {
					return checkCell.distance + 1
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
				if downNeighbour == destination {
					return checkCell.distance + 1
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
				if leftNeighbour == destination {
					return checkCell.distance + 1
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
				if rightNeighbour == destination {
					return checkCell.distance + 1
				}
				rightNeighbour.distance = checkCell.distance + 1
				rightNeighbour.visited = true
				grid[rightNeighbour.y][rightNeighbour.x] = rightNeighbour
				possibilePathQueue = append(possibilePathQueue, rightNeighbour)
			}
		}
	}
	return -1
}

func hasValidAdjacentCells(checkCell Cell, grid [][]Cell) bool {
	// check up
	if checkCell.y-1 >= 0 {
		upNeighbour := grid[checkCell.y-1][checkCell.x]
		if isValidStep(checkCell, upNeighbour) {
			return true
		}
	}
	// check down
	if checkCell.y+1 < len(grid) {
		downNeighbour := grid[checkCell.y+1][checkCell.x]
		if isValidStep(checkCell, downNeighbour) {
			return true
		}
	}
	// check left
	if checkCell.x-1 >= 0 {
		leftNeighbour := grid[checkCell.y][checkCell.x-1]
		if isValidStep(checkCell, leftNeighbour) {
			return true
		}
	}
	// check right
	if checkCell.x+1 < len(grid[0]) {
		rightNeighbour := grid[checkCell.y][checkCell.x+1]
		if isValidStep(checkCell, rightNeighbour) {
			return true
		}
	}
	return false
}
