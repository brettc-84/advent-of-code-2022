package day03

import (
	"strconv"
	"unicode"
)

func getItemValue(char byte) rune {
	if unicode.IsLower(rune(char)) {
		return rune(char) - 96
	}
	return rune(char) - 38
}

func findDuplicate(rucksack string) (string, int) {
	for i1 := 0; i1 < len(rucksack)/2; i1++ {
		for i2 := len(rucksack) / 2; i2 < len(rucksack); i2++ {
			if rucksack[i1] == rucksack[i2] {

				// fmt.Println("---", getItemValue(rucksack[i1]))
				return string(rucksack[i1]), int(getItemValue(rucksack[i1]))
			}
		}
	}
	return "", 0
}

func Part1(input []string) string {
	result := 0
	for _, rucksack := range input {
		_, value := findDuplicate(rucksack)
		result += value
	}
	return strconv.Itoa(result)
}
