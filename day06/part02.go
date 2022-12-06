package day06

func Part2(input string) int {
	start := 0
	for i := 14; i < len(input); i++ {
		mark := input[start:i]
		if hasUnique(mark) {
			return i
		}
		start += 1
	}
	return -1
}
