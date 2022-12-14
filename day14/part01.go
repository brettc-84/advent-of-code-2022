package day14

import (
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func Part1(input []string) string {
	result := 0

	world := make([]Point, 0)

	maxY := 0

	for _, path := range input {
		pathPoints := strings.Split(path, " -> ")
		for i := 0; i < len(pathPoints)-1; i++ {
			fromXy := strings.Split(pathPoints[i], ",")
			fromX, _ := strconv.Atoi(fromXy[0])
			fromY, _ := strconv.Atoi(fromXy[1])

			toXy := strings.Split(pathPoints[i+1], ",")
			toX, _ := strconv.Atoi(toXy[0])
			toY, _ := strconv.Atoi(toXy[1])

			if toY > maxY {
				maxY = toY
			}

			if fromX == toX {
				// vertical line
				if fromY < toY {
					for yy := fromY; yy <= toY; yy++ {
						rock := Point{x: fromX, y: yy}
						world = appendUnique(world, rock)
					}
				} else {
					for yy := fromY; yy >= toY; yy-- {
						rock := Point{x: fromX, y: yy}
						world = appendUnique(world, rock)
					}
				}
			} else {
				// horizontal line
				if fromX < toX {
					for xx := fromX; xx <= toX; xx++ {
						rock := Point{x: xx, y: fromY}
						world = appendUnique(world, rock)
					}
				} else {
					for xx := fromX; xx >= toX; xx-- {
						rock := Point{x: xx, y: fromY}
						world = appendUnique(world, rock)
					}
				}
			}
		}
	}

	sandFalling := true

	for sandFalling {
		sandBit := Point{500, 0}

		canMove := true
		for canMove {
			// in bounds?
			if sandBit.y > maxY {
				sandFalling = false
				break
			}

			// try move down
			down := Point{x: sandBit.x, y: sandBit.y + 1}
			if whatIsAtPointInWorld(down, world) == 0 {
				// move down
				sandBit.y = sandBit.y + 1
				continue
			}
			// try move down-left
			downLeft := Point{x: sandBit.x - 1, y: sandBit.y + 1}
			if whatIsAtPointInWorld(downLeft, world) == 0 {
				// move down-left
				sandBit.x -= 1
				sandBit.y += 1
				continue
			}

			// try move down-right
			downRight := Point{x: sandBit.x + 1, y: sandBit.y + 1}
			if whatIsAtPointInWorld(downRight, world) == 0 {
				// move down-right
				sandBit.x += 1
				sandBit.y += 1
				continue
			}
			world = append(world, sandBit)
			result += 1

			// can't move
			canMove = false
		}
	}

	return strconv.Itoa(result)
}

func whatIsAtPointInWorld(point Point, world []Point) int {
	// 0 = nothing
	// 1 = something
	for _, coord := range world {
		if point == coord {
			return 1
		}
	}
	return 0
}

func appendUnique(points []Point, toAdd Point) []Point {
	for _, point := range points {
		if point == toAdd {
			return points
		}
	}
	return append(points, toAdd)
}
