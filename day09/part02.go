package day09

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Snake struct {
	head Point
	body []Point
}

func Part2(input []string) string {
	result := 0

	snake := Snake{
		head: Point{0, 0},
		body: make([]Point, 9),
	}

	positionsVisited := make(map[string]int)
	positionsVisited["0,0"] = 1 // start

	for _, headMove := range input {
		move := strings.Split(headMove, " ")
		directionMoved, stepsMoved := move[0], move[1]
		steps, _ := strconv.Atoi(stepsMoved)
		for i := 0; i < steps; i++ {
			// move head
			if directionMoved == "U" {
				snake.head.y += 1
			} else if directionMoved == "D" {
				snake.head.y -= 1
			} else if directionMoved == "R" {
				snake.head.x += 1
			} else if directionMoved == "L" {
				snake.head.x -= 1
			}
			for i, cell := range snake.body {
				var follow Point
				if i == 0 {
					// check head
					follow = snake.head
				} else {
					// check previous
					follow = snake.body[i-1]
				}
				if !areTouching(follow.x, follow.y, cell.x, cell.y) {
					if cell.x == follow.x {
						if cell.y < follow.y {
							snake.body[i].y += 1
						} else if cell.y > follow.y {
							snake.body[i].y -= 1
						}
					} else if cell.y == follow.y {
						if cell.x < follow.x {
							snake.body[i].x += 1
						} else if cell.x > follow.x {
							snake.body[i].x -= 1
						}
					} else {
						// diagonal move needed
						if follow.x > cell.x {
							snake.body[i].x += 1
							if follow.y > cell.y {
								snake.body[i].y += 1
							} else {
								snake.body[i].y -= 1
							}
						} else if follow.x < cell.x {
							snake.body[i].x -= 1
							if follow.y > cell.y {
								snake.body[i].y += 1
							} else {
								snake.body[i].y -= 1
							}
						}
					}

					tail := snake.body[len(snake.body)-1]
					key := fmt.Sprintf("%d,%d", tail.x, tail.y)
					_, alreadyVisited := positionsVisited[key]
					if alreadyVisited {
						positionsVisited[key] += 1
					} else {
						positionsVisited[key] = 1
					}
				}
			}
		}
	}

	result = len(positionsVisited)

	return strconv.Itoa(result)
}
