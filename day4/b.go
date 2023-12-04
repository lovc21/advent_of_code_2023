package day4

import (
	"fmt"
	"strings"
)

func Day4task2() {
	var str string = `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

	lines := strings.Split(str, "\n")

	var updatedLines []string

	toltalsum := 0
	for index, line := range lines {
		if line == "" {
			continue
		}

		updatedLines = append(updatedLines, line) // Add original line
		processLine(&updatedLines, lines, index, line)

	}

	lenoflines := len(updatedLines)
	fmt.Println(lenoflines)
	fmt.Println(lines)
	fmt.Println(toltalsum)
	fmt.Println("Hello, playground day4 is working ")
}

func processLine(updatedLines *[]string, lines []string, index int, line string) {
	removeGame := strings.Split(line, ":")
	removeCard := strings.Split(removeGame[1], "|")
	numbersBefore := strings.Fields(strings.TrimSpace(removeCard[0]))
	numbersAfter := strings.Fields(strings.TrimSpace(removeCard[1]))

	matches := make(map[string]bool)
	for _, num1 := range numbersBefore {
		for _, num2 := range numbersAfter {
			if num1 == num2 {
				matches[num1] = true
			}
		}
	}

	count := len(matches)
	for i := 1; i <= count; i++ {
		if index+i < len(lines) && lines[index+i] != "" {
			*updatedLines = append(*updatedLines, lines[index+i])
			processLine(updatedLines, lines, index+i, lines[index+i])
		}
	}
}
