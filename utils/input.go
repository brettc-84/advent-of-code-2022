package utils

import (
	"os"
	"strings"
)

// From https://gobyexample.com/reading-files
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadMultiLine(filePath string) []string {
	dat, err := os.ReadFile(filePath)
	check(err)
	return strings.Split(string(dat), "\n")
}

func ReadSingleLine(filePath string) string {
	return ReadMultiLine(filePath)[0]
}
