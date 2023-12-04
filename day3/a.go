package day3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day3task1() {
	var str string = `
	467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..`

	// split stri in to lines
	lines := strings.Split(str, "\n")
	//toltalsum := 0
	nubmer := ""
	nubmerStringArr := []string{}
	symbols := []string{"*", "#", "+", "$", "-", "!", "(", ")", ":", ";", "=", "?", "@", "[", "]", "^", "_", "{", "}", "|", "~"}
	numberOfLines := len(lines)
	linelen := len(lines[0])

	ARR2 := [numberOfLines][linelen]string{}

	for _, line := range lines {
		//check if number is in line
		numeric := regexp.MustCompile(`\d`).MatchString(line)
		if numeric {
			for i, ch := range line {
				// check if ther is a number and add it to the array
				if _, err := strconv.Atoi(string(ch)); err == nil {
					nubmer = nubmer + string(ch)
				}

				// check if ther is a symbol next to the number and add it to the array
				if string(ch) == "*" || string(ch) == "$" || string(ch) == "+" {
					if nubmer != "" {
						nubmerStringArr = append(nubmerStringArr, nubmer)
						nubmer = ""
					}
				}

				// check if ther is a symbol before to the number and add it to the array
				if string(ch) == "*" || string(ch) == "$" || string(ch) == "+" {
					if nubmer == "" {
						nubmerStringArr = append(nubmerStringArr, nubmer)
						nubmer = ""
					}
				}

				// add number to array
				if nubmer != "" {
					if string(ch) == "." {
						// check if line above has an simbol for the number len + 1
						if i > 0 && containsSymbol(lines[i-1], symbols, len(nubmer), i) {
							// check if line below has an simbol for the number len + 1
							if i < len(lines)-1 && containsSymbol(lines[i+1], symbols, len(nubmer), i) {
								nubmerStringArr = append(nubmerStringArr, nubmer)
								nubmer = ""
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(nubmerStringArr)
}

func containsSymbol(line string, symbols []string, lent int, index int) bool {
	// if i >= 0 and i < len(a) and j >= 0 and j < len(a[0]
	for i := index - lent - 1; i < index; i++ {
		for _, symbol := range symbols {
			if i < 0 || i >= len(line) || i < 0 || i >= len(line) {
				continue
			}
			ch := line[i]
			if string(ch) == symbol {
				return true
			}
		}
	}
	return false
}
