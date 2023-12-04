package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func Day1task2() {

	var str string = `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	// split stri in to lines
	lines := strings.Split(str, "\n")

	//check if number is in line
	newNumberArr := []int{}
	for _, line := range lines {
		stringBuild := ""
		firstNum, lastNum := -1, -1
		for _, ch := range line {
			stringBuild = stringBuild + string(ch)
			if s, err := strconv.Atoi(string(ch)); err == nil {
				if firstNum == -1 {
					firstNum = s
					stringBuild = ""
				}
				stringBuild = ""
				lastNum = s
			}
			if strings.Contains(stringBuild, "one") {
				if firstNum == -1 {
					firstNum = 1
					stringBuild = ""
				}
				lastNum = 1
				stringBuild = ""
			} else if strings.Contains(stringBuild, "two") {
				if firstNum == -1 {
					firstNum = 2
					stringBuild = ""
				}
				lastNum = 2
				stringBuild = ""
			} else if strings.Contains(stringBuild, "three") {
				if firstNum == -1 {
					firstNum = 3
					stringBuild = ""
				}
				lastNum = 3
				stringBuild = ""
			} else if strings.Contains(stringBuild, "four") {
				if firstNum == -1 {
					firstNum = 4
					stringBuild = ""
				}
				lastNum = 4
				stringBuild = ""
			} else if strings.Contains(stringBuild, "five") {
				if firstNum == -1 {
					firstNum = 5
					stringBuild = ""
				}
				lastNum = 5
				stringBuild = ""
			} else if strings.Contains(stringBuild, "six") {
				if firstNum == -1 {
					firstNum = 6
					stringBuild = ""
				}
				lastNum = 6
				stringBuild = ""
			} else if strings.Contains(stringBuild, "seven") {
				if firstNum == -1 {
					firstNum = 7
					stringBuild = ""
				}
				lastNum = 7
				stringBuild = ""
			} else if strings.Contains(stringBuild, "eight") {
				if firstNum == -1 {
					firstNum = 8
					stringBuild = ""
				}
				lastNum = 8
				stringBuild = ""
			} else if strings.Contains(stringBuild, "nine") {
				if firstNum == -1 {
					firstNum = 9
					stringBuild = ""
				}
				stringBuild = ""
				lastNum = 9
			}

		}
		if firstNum != -1 {
			newNumberArr = append(newNumberArr, firstNum)
		}
		if lastNum != -1 {
			newNumberArr = append(newNumberArr, lastNum)
		}
	}
	finalNumberArr := []int{}

	for i := 0; i < len(newNumberArr); i = i + 2 {
		if i+1 < len(newNumberArr) {
			combine, _ := strconv.Atoi(strconv.Itoa(newNumberArr[i]) + strconv.Itoa(newNumberArr[i+1]))
			finalNumberArr = append(finalNumberArr, combine)
		}
	}

	final := 0

	for i := 0; i < len(finalNumberArr); i = i + 1 {
		final = final + finalNumberArr[i]
	}

	fmt.Println(final)
	fmt.Println(finalNumberArr)
	//fmt.Println(newNumberArr)
	//fmt.Println(lines)
}
