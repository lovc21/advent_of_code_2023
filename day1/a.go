package main

import (
	"fmt"
	"strings"
)

func wordToDigit(word string) (int, bool) {
	numberMap := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4,
		"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
	value, exists := numberMap[word]
	return value, exists
}

func extractDigits(line string) (int, int, bool) {
	words := strings.FieldsFunc(line, func(r rune) bool {
		return !('a' <= r && r <= 'z')
	})

	if len(words) == 0 {
		return 0, 0, false
	}

	firstDigit, ok := wordToDigit(words[0])
	if !ok {
		return 0, 0, false
	}

	lastDigit, ok := wordToDigit(words[len(words)-1])
	if !ok {
		return 0, 0, false
	}

	return firstDigit, lastDigit, true
}

func main() {
	inputLines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	totalSum := 0
	for _, line := range inputLines {
		firstDigit, lastDigit, found := extractDigits(line)
		if found {
			calibrationValue := firstDigit*10 + lastDigit
			totalSum += calibrationValue
		}
	}

	fmt.Println("Total Sum of Calibration Values:", totalSum)
}
