package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func day1task1() {

	var str string = `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	// split stri in to lines
	lines := strings.Split(str, "\n")

	//check if number is in line
	newNumberArr := []int{}
	for _, line := range lines {
		firstNum, lastNum := -1, -1
		for _, ch := range line {

			if s, err := strconv.Atoi(string(ch)); err == nil {
				if firstNum == -1 {
					firstNum = s
				}
				lastNum = s
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
	//fmt.Println(finalNumberArr)
	//fmt.Println(newNumberArr)
	//fmt.Println(lines)
}
