package day07

import (
	"strconv"
	"strings"
)

const minSpace = 30000000
const totalSpace = 70000000

func Part2(input []string) string {
	root := Folder{
		name:  "/",
		dirs:  make(map[string]*Folder),
		files: make([]File, 0),
	}
	var currentDir = &root
	allDirSizes := make(map[string]int)

	// build nodes
	for _, bash := range input {

		if strings.HasPrefix(bash, "$ cd") {
			// cd command
			gotoDir := strings.Replace(bash, "$ cd ", "", 1)
			if gotoDir == "/" {
				*currentDir = root
			} else if gotoDir == ".." {
				currentDir = currentDir.parent
			} else {
				nextDir, exists := currentDir.dirs[gotoDir]
				if !exists {
					newDir := Folder{
						name:   gotoDir,
						parent: currentDir,
						dirs:   make(map[string]*Folder),
						files:  make([]File, 0),
					}
					currentDir.dirs[gotoDir] = &newDir
					*currentDir = newDir
				} else {
					currentDir = nextDir
				}
			}
		} else if strings.HasPrefix(bash, "$ ls") {
			// ls command
			continue
		} else {
			// result of ls
			lsResult := strings.Split(bash, " ")
			if lsResult[0] == "dir" {
				// directory
				dirName := lsResult[1]
				_, exists := currentDir.dirs[dirName]
				if !exists {
					childFolder := Folder{
						name:   dirName,
						parent: currentDir,
						dirs:   make(map[string]*Folder),
						files:  make([]File, 0),
					}
					currentDir.dirs[dirName] = &childFolder
				}
			} else {
				size, fileName := lsResult[0], lsResult[1]
				sizeValue, _ := strconv.Atoi(size)
				file := File{name: fileName, size: sizeValue}
				currentDir.files = append(currentDir.files, file)
				currentDir.size += sizeValue
			}
		}
	}
	// calculate sizes
	calcSizeWithMap(&root, allDirSizes)

	return strconv.Itoa(findSizeOfSmallestFolderDelete(allDirSizes, root.size))
}

func calcSizeWithMap(start *Folder, fm map[string]int) int {
	if len(start.dirs) == 0 {
		fm[start.name] = start.size
		return start.size
	}
	totalSize := 0
	for _, dir := range start.dirs {
		fm[dir.name] = dir.size
		totalSize += calcSizeWithMap(dir, fm)
	}
	start.size += totalSize
	fm[start.name] = start.size
	return start.size
}

func findSizeOfSmallestFolderDelete(dirs map[string]int, rootSize int) int {
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
