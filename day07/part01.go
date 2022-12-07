package day07

import (
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Folder struct {
	name   string
	files  []File
	dirs   map[string]*Folder
	parent *Folder
	size   int
}

func Part1(input []string) string {
	result := 0

	root := Folder{
		name:  "/",
		dirs:  make(map[string]*Folder),
		files: make([]File, 0),
	}
	var currentDir = &root

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
	calcSize(&root)

	for _, dir := range root.dirs {
		if dir.size <= 100000 {
			result += dir.size
		}
		result += calculateResult(dir)
	}

	return strconv.Itoa(result)
}

func calcSize(start *Folder) int {
	if len(start.dirs) == 0 {
		return start.size
	}
	totalSize := 0
	for _, dir := range start.dirs {
		totalSize += calcSize(dir)
	}
	start.size += totalSize
	return start.size
}

func calculateResult(start *Folder) int {
	size := 0
	for _, dir := range start.dirs {
		if dir.size <= 100000 {
			size += dir.size
		}
		size += calculateResult(dir)
	}
	return size
}
