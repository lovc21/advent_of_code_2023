package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2task2() {

	var str string = `
	Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
	`
	gameSum := 0
	maxRed := 0
	maxGreen := 0
	maxBlue := 0
	// split lines
	lines := strings.Split(strings.TrimSpace(str), "\n")

	// total sum of allowed games
	totalSum := 0
	for _, line := range lines {
		idAndGame := strings.Split(strings.TrimSpace(line), ":")
		maxRed = 0
		maxGreen = 0
		maxBlue = 0
		for _, words := range idAndGame[1:] {
			oneGame := strings.Split(strings.TrimSpace(words), ";")
			for _, word := range oneGame {
				colorNumbers := strings.Split(strings.TrimSpace(word), ",")
				for _, colorNumber := range colorNumbers {
					parts := strings.Split(strings.TrimSpace(colorNumber), " ")
					number, _ := strconv.Atoi(parts[0])
					color := parts[1]
					if color == "red" {
						if number > maxRed {
							maxRed = number
						}
					} else if color == "green" {
						if number > maxGreen {
							maxGreen = number
						}
					} else if color == "blue" {
						if number > maxBlue {
							maxBlue = number
						}
					}
				}
			}
		}
		// sum of allowed games for each id
		gameSum = maxRed * maxGreen * maxBlue
		fmt.Println(gameSum)
		totalSum += gameSum
	}
	fmt.Println(totalSum)
}
