package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var str string = `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	// split stri in to lines
	lines := strings.Split(str, "\n")

	//check if number is in line and is first or last
	leterNumberArr := []int{}
	newNumberArr := []int{}
	for _, line := range lines {
		firstNum, lastNum := -1, -1
		if strings.Contains(line, "one") {
			leterNumberArr = append(leterNumberArr, 1)
			if firstNum == -1 {
				firstNum = 1
			}
			lastNum = 1
		} else if strings.Contains(line, "two") {
			leterNumberArr = append(leterNumberArr, 2)
			if firstNum == -1 {
				firstNum = 2
			}
			lastNum = 2
		} else if strings.Contains(line, "three") {
			leterNumberArr = append(leterNumberArr, 3)
			if firstNum == -1 {
				firstNum = 3
			}
			lastNum = 3
		} else if strings.Contains(line, "four") {
			leterNumberArr = append(leterNumberArr, 4)
			if firstNum == -1 {
				firstNum = 4
			}
			lastNum = 4
		} else if strings.Contains(line, "five") {
			leterNumberArr = append(leterNumberArr, 5)
			if firstNum == -1 {
				firstNum = 5
			}
			lastNum = 5
		} else if strings.Contains(line, "six") {
			leterNumberArr = append(leterNumberArr, 6)
			if firstNum == -1 {
				firstNum = 6
			}
			lastNum = 6
		} else if strings.Contains(line, "seven") {
			leterNumberArr = append(leterNumberArr, 7)
			if firstNum == -1 {
				firstNum = 7
			}
			lastNum = 7
		} else if strings.Contains(line, "eight") {
			leterNumberArr = append(leterNumberArr, 8)
			if firstNum == -1 {
				firstNum = 8
			}
			lastNum = 8
		} else if strings.Contains(line, "nine") {
			leterNumberArr = append(leterNumberArr, 9)
			if firstNum == -1 {
				firstNum = 9
			}
			lastNum = 9
		}
		for _, ch := range line {
			if s, err := strconv.Atoi(string(ch)); err == nil {
				if firstNum == -1 {
					firstNum = s
				}
				lastNum = s
			}
			if firstNum != -1 {
				newNumberArr = append(newNumberArr, firstNum)
			}
			if lastNum != -1 {
				newNumberArr = append(newNumberArr, lastNum)
			}
		}
	}
	//combine the numbers
	finalNumberArr := []int{}
	for i := 0; i < len(newNumberArr); i = i + 2 {
		if i+1 < len(newNumberArr) {
			combine, _ := strconv.Atoi(strconv.Itoa(newNumberArr[i]) + strconv.Itoa(newNumberArr[i+1]))
			finalNumberArr = append(finalNumberArr, combine)
		}
	}

	//calculate the result
	final := 0
	for i := 0; i < len(finalNumberArr); i = i + 1 {
		final = final + finalNumberArr[i]
	}

	fmt.Println(leterNumberArr)
	fmt.Println(final)
	fmt.Println(finalNumberArr)
	fmt.Println(newNumberArr)
	fmt.Println(lines)

}
