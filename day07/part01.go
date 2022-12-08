package day07

import (
	"strconv"
	"strings"
)

func Part1(input []string) string {
	result := 0

	for _, dirSize := range buildDirSizeMap(input) {
		if dirSize <= 100000 {
			result += dirSize
		}
	}

	return strconv.Itoa(result)
}

func buildDirSizeMap(input []string) map[string]int {

	dirStack := []string{}
	dirSizes := make(map[string]int)

	for _, bash := range input {
		if isListCommand(bash) {
			continue
		} else if isChangeDirCommand(bash) {
			dirName := getDirFromCd(bash)
			if dirName == ".." {
				dirStack = dirStack[0 : len(dirStack)-1]
			} else {
				dirStack = append(dirStack, dirName)
			}
		} else {
			size, _ := parseLsOutput(bash)
			if size != "dir" {
				sizeValue, _ := strconv.Atoi(size)
				currentDir := dirStack[len(dirStack)-1]
				_, exists := dirSizes[currentDir]
				if exists {
					dirSizes[currentDir] = dirSizes[currentDir] + sizeValue
				} else {
					dirSizes[currentDir] = sizeValue
				}
				for i := 0; i < len(dirStack)-1; i++ {
					parentDir := dirStack[i]
					_, exists := dirSizes[parentDir]
					if exists {
						dirSizes[parentDir] = dirSizes[parentDir] + sizeValue
					} else {
						dirSizes[parentDir] = sizeValue
					}
				}
			}
		}
	}
	return dirSizes
}

func isListCommand(bash string) bool {
	return strings.HasPrefix(bash, "$ ls")
}

func isChangeDirCommand(bash string) bool {
	return strings.HasPrefix(bash, "$ cd")
}

func getDirFromCd(command string) string {
	return strings.Replace(command, "$ cd ", "", 1)
}

func parseLsOutput(output string) (string, string) {
	splits := strings.Split(output, " ")
	return splits[0], splits[1]
}
