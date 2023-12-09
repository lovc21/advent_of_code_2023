package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func Day9task1() {
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

		var lastNumbers []int
		for {
			fmt.Println("Processing numbers:", numbers)
			lastNumbers = append(lastNumbers, numbers[len(numbers)-1])
			if len(numbers) <= 1 {
				break
			}

			var allSame bool
			numbers, allSame = calculateDifferences(numbers)
			if allSame {
				lastNumbers = append(lastNumbers, numbers[len(numbers)-1])
				break
			}
		}
		sum := 0
		for _, number := range lastNumbers {
			sum += number
		}
		totalsum += sum
	}
	fmt.Println(totalsum)
}

func calculateDifferences(numbers []int) ([]int, bool) {
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
