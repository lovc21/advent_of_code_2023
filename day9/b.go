package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func Day9task2() {
	str := `0 3 6 9 12 15
	1 3 6 10 15 21
	10 13 16 21 30 45`
	totalsum := 0
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")

		numbers := make([]int, 0, len(words)-1)

		for _, word := range words {
			number, err := strconv.Atoi(strings.TrimSpace(word))
			if err != nil {
				fmt.Println("Error converting word to integer:", err)
				break
			}

			numbers = append(numbers, number)
		}

		var firstNumber []int
		for {
			fmt.Println("Processing numbers:", numbers)
			firstNumber = append(firstNumber, numbers[0])
			if len(numbers) <= 1 {
				break
			}

			var allSame bool
			numbers, allSame = CalculateDifferences(numbers)
			if allSame {
				firstNumber = append(firstNumber, numbers[0])
				break
			}
		}
		fmt.Println(firstNumber)
		var result int
		if len(firstNumber) > 0 {
			for i, number := range firstNumber {
				if i%2 == 0 {
					result += number
				} else {
					result -= number
				}
			}
		}
		totalsum += result
	}
	fmt.Println(totalsum)
}

func CalculateDifferences(numbers []int) ([]int, bool) {
	diffs := make([]int, len(numbers)-1)
	allSame := true
	firstDiff := 0
	if len(numbers) > 1 {
		firstDiff = numbers[1] - numbers[0]
	}
	for i := 1; i < len(numbers); i++ {
		diffs[i-1] = numbers[i] - numbers[i-1]
		if diffs[i-1] != firstDiff {
			allSame = false
		}
	}
	return diffs, allSame
}
