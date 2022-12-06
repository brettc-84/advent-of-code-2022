package day03

import "strconv"

func unique(runeSlice []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}
	for _, entry := range runeSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Part2(input []string) string {
	total := 0
	for i := 0; i < len(input)-2; i += 3 {
		elf1 := input[i]
		elf2 := input[i+1]
		elf3 := input[i+2]

		var firstCommon []rune
		for _, item1 := range elf1 {
			for _, item2 := range elf2 {
				if item1 == item2 {
					firstCommon = append(firstCommon, rune(item1))
				}
			}
		}

		var finalCommon []rune
		for _, f1 := range firstCommon {
			for _, item3 := range elf3 {
				if f1 == item3 {
					finalCommon = append(finalCommon, rune(f1))
				}
			}
		}
		uniqueValues := unique(finalCommon)
		total += int(getItemValue(byte(uniqueValues[0])))
	}
	return strconv.Itoa(total)
}
