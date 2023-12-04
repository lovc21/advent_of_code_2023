package day4

import (
	"fmt"
	"math"
	"strings"
)

func Day4task1() {
	var str string = `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

	lines := strings.Split(str, "\n")
	toltalsum := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		remmuveGame := strings.Split(line, ":")
		remmuveCard := strings.Split(remmuveGame[1], "|")
		numbersBefore := strings.Fields(strings.TrimSpace(remmuveCard[0]))
		numbersAfter := strings.Fields(strings.TrimSpace(remmuveCard[1]))

		matches := make(map[string]bool)
		for _, num1 := range numbersBefore {
			for _, num2 := range numbersAfter {
				if num1 == num2 {
					matches[num1] = true
				}
			}
		}

		count := len(matches)
		if count > 0 {
			toltalsum = toltalsum + int(math.Pow(2, float64(count-1)))
		}
	}

	fmt.Println(toltalsum)
	fmt.Println("Hello, playground day4 is working ")
}
