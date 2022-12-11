package day11

import (
	"strconv"
	"strings"
)

func Part2(input []string) string {
	result := 0

	monkeys := make([]Monkey, 0)
	cd := 1

	for _, line := range input {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Monkey") {
			monkeys = append(monkeys, Monkey{inspections: 0})
			continue
		}
		thisMonkey := &monkeys[len(monkeys)-1]
		if strings.HasPrefix(line, "Starting items") {
			thisMonkey.items = make([]int, 0)
			itemsList := strings.Split(strings.Replace(line, "Starting items: ", "", 1), ",")
			for _, item := range itemsList {
				itemLevel, _ := strconv.Atoi(strings.TrimSpace(item))
				thisMonkey.items = append(thisMonkey.items, itemLevel)
			}
			continue
		}
		if strings.HasPrefix(line, "Operation") {
			operationText := strings.Replace(line, "Operation: new = ", "", 1)
			if strings.Contains(operationText, "+") {
				addend, _ := strconv.Atoi(strings.Split(operationText, " + ")[1])
				thisMonkey.operation = func(old int) int {
					return old + addend
				}
			} else if strings.Contains(operationText, "*") {
				factors := strings.Split(operationText, " * ")
				if factors[1] == "old" {
					thisMonkey.operation = func(old int) int {
						return old * old
					}
				} else {
					factorValue, _ := strconv.Atoi(factors[1])
					thisMonkey.operation = func(old int) int {
						return old * factorValue
					}
				}
			}
			continue
		}
		if strings.HasPrefix(line, "Test") {
			divisor, _ := strconv.Atoi(strings.Replace(line, "Test: divisible by ", "", 1))
			cd *= divisor
			thisMonkey.test = Test{
				check: func(old int) bool {
					return old%divisor == 0
				},
			}
			continue
		}
		if strings.HasPrefix(line, "If true") {
			thisMonkey.test.true, _ = strconv.Atoi(strings.Replace(line, "If true: throw to monkey ", "", 1))
			continue
		}

		if strings.HasPrefix(line, "If false") {
			thisMonkey.test.false, _ = strconv.Atoi(strings.Replace(line, "If false: throw to monkey ", "", 1))
			continue
		}
	}

	// start game
	rounds := 10000
	max1, max2 := 0, 0
	for i := 0; i < rounds; i++ {
		for m, monkey := range monkeys {
			for _, item := range monkey.items {
				newWorryLevel := monkey.operation(item)
				boredLevel := newWorryLevel % cd
				checkResult := monkey.test.check(boredLevel)
				monkeys[m].inspections += 1
				if monkeys[m].inspections > max1 {
					max1 = monkeys[m].inspections
				} else if monkeys[m].inspections > max2 {
					max2 = monkeys[m].inspections
				}
				if checkResult {
					monkeys[monkey.test.true].items = append(monkeys[monkey.test.true].items, boredLevel)
				} else {
					monkeys[monkey.test.false].items = append(monkeys[monkey.test.false].items, boredLevel)
				}
			}
			monkeys[m].items = make([]int, 0)
		}
	}

	result = max1 * max2
	return strconv.Itoa(result)
}
