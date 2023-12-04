package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2task1() {

	var str string = `
	Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
	`
	// max allowed numbers
	var red int = 12
	var green int = 13
	var blue int = 14

	// split lines
	lines := strings.Split(strings.TrimSpace(str), "\n")

	// total sum of allowed games
	totalSum := 0
	for _, line := range lines {
		idAndGame := strings.Split(strings.TrimSpace(line), ":")
		numberId := strings.Split(strings.TrimSpace(idAndGame[0]), " ")
		idsum, _ := strconv.Atoi(numberId[1])
		gameAllowed := true
		for _, words := range idAndGame[1:] {
			oneGame := strings.Split(strings.TrimSpace(words), ";")
			for _, word := range oneGame {
				colorNumbers := strings.Split(strings.TrimSpace(word), ",")
				for _, colorNumber := range colorNumbers {
					parts := strings.Split(strings.TrimSpace(colorNumber), " ")
					number, _ := strconv.Atoi(parts[0])
					color := parts[1]
					if (color == "red" && number > red) || (color == "green" && number > green) || (color == "blue" && number > blue) {
						gameAllowed = false
						break
					}
				}
				if !gameAllowed {
					break
				}
			}
			if !gameAllowed {
				break
			}
		}
		if gameAllowed {
			totalSum += idsum
		}
	}

	fmt.Println(totalSum)

}
