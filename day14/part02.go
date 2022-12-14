package day14

import (
	"strconv"
)

func Part2(input []string) string {
	result := 0

	world, maxY := buildWorld(input)

	sandFalling := true

	floorY := maxY + 2

	for x := 100; x < 900; x++ {
		floor := Point{x: x, y: floorY}
		world[floor.ToString()] = true
	}

	for sandFalling {
		sandBit := Point{500, 0}

		canMove := true
		for canMove {

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
			world[sandBit.ToString()] = true
			result += 1
			if (sandBit == Point{x: 500, y: 0}) {
				sandFalling = false
				break
			}

			// can't move
			canMove = false
		}
	}

	return strconv.Itoa(result)
}
