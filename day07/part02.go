package day07

import (
	"strconv"
)

const minSpace = 30000000
const totalSpace = 70000000

func Part2(input []string) string {
	allDirSizes := buildDirSizeMap(input)

	return strconv.Itoa(findSizeOfSmallestFolderDelete(allDirSizes))
}

func findSizeOfSmallestFolderDelete(dirs map[string]int) int {
	rootSize := dirs["/"]
	unused := totalSpace - rootSize
	spaceRequired := minSpace - unused

	smallest := rootSize

	for _, dirSize := range dirs {
		if dirSize >= spaceRequired && dirSize < smallest {
			smallest = dirSize
		}
	}
	return smallest
}
