package day09

import (
	"strconv"
	"strings"
)

func Part1(input []string) string {
	result := 0

	snake := Snake{
		head: Point{0, 0},
		body: make([]Point, 1),
	}

	positionsVisited := make(map[string]int)
	positionsVisited["0,0"] = 1 // start

	for _, headMove := range input {
		move := strings.Split(headMove, " ")
		directionMoved, stepsMoved := move[0], move[1]
		steps, _ := strconv.Atoi(stepsMoved)
		for i := 0; i < steps; i++ {
			snake.head.move(directionMoved)
			snake.moveSnake()
		}

	}

	result = len(positionsVisited)

	return strconv.Itoa(result)
}

func (point *Point) move(direction string) {
	switch direction {
	case "U":
		point.y += 1
	case "D":
		point.y -= 1
	case "R":
		point.x += 1
	case "L":
		point.x -= 1
	case "UR":
		point.x += 1
		point.y += 1
	case "DR":
		point.x += 1
		point.y -= 1
	case "UL":
		point.x -= 1
		point.y += 1
	case "DL":
		point.x -= 1
		point.y -= 1
	}
}

func (snake *Snake) moveSnake() {
	for i, cell := range snake.body {
		var follow Point
		if i == 0 {
			follow = snake.head
		} else {
			follow = snake.body[i-1]
		}
		if !areTouching(follow.x, follow.y, cell.x, cell.y) {

		}
	}
}

func areTouching(headX, headY, tailX, tailY int) bool {
	if (tailX == headX && tailY == headY) ||
		(headX == tailX+1 && headY == tailY) ||
		(headX == tailX-1 && headY == tailY) ||
		(headY == tailY+1 && headX == tailX) ||
		(headY == tailY-1 && headX == tailX) ||
		(headX == tailX+1 && headY == tailY+1) ||
		(headX == tailX-1 && headY == tailY-1) ||
		(headX == tailX+1 && headY == tailY-1) ||
		(headX == tailX-1 && headY == tailY+1) {
		return true
	}
	return false
}
