package day09

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input []string) string {
	result := 0

	headX, tailX, headY, tailY := 0, 0, 0, 0
	positionsVisited := make(map[string]int)
	positionsVisited["0,0"] = 1 // start

	for _, headMove := range input {
		move := strings.Split(headMove, " ")
		directionMoved, stepsMoved := move[0], move[1]
		steps, _ := strconv.Atoi(stepsMoved)
		for i := 0; i < steps; i++ {
			if directionMoved == "L" {
				headX -= 1
			} else if directionMoved == "R" {
				headX += 1
			} else if directionMoved == "D" {
				headY -= 1
			} else if directionMoved == "U" {
				headY += 1
			}
			if !areTouching(headX, headY, tailX, tailY) {
				if tailX == headX || tailY == headY {
					switch directionMoved {
					case "R":
						tailX += 1
					case "U":
						tailY += 1
					case "L":
						tailX -= 1
					case "D":
						tailY -= 1
					}
				} else {
					// diagonal move needed
					if headX > tailX {
						tailX += 1
						if headY > tailY {
							tailY += 1
						} else {
							tailY -= 1
						}
					} else if headX < tailX {
						tailX -= 1
						if headY > tailY {
							tailY += 1
						} else {
							tailY -= 1
						}
					}
				}
				key := fmt.Sprintf("%d,%d", tailX, tailY)
				_, alreadyVisited := positionsVisited[key]
				if alreadyVisited {
					positionsVisited[key] += 1
				} else {
					positionsVisited[key] = 1
				}
			}
		}
	}

	result = len(positionsVisited)

	return strconv.Itoa(result)
}

// func moveTail(headX int, headY int, tailX int, tailY int) (int int) {
// 	if tailX == headX || tailY == headY {
// 		// vertical move
// 		switch directionMoved {
// 		case "R":
// 			tailX += 1
// 		case "U":
// 			tailY += 1
// 		case "L":
// 			tailX -= 1
// 		case "D":
// 			headY -= 1
// 		}
// 	} else {
// 		// diagonal move needed
// 		if headX > tailX {
// 			tailX += 1
// 			if headY > tailY {
// 				tailY += 1
// 			} else {
// 				tailY -= 1
// 			}
// 		} else if headX < tailX {
// 			tailX -= 1
// 			if headY > tailY {
// 				tailY += 1
// 			} else {
// 				tailY -= 1
// 			}
// 		}
// 	}
// }

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
