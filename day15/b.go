package day15

import (
	"fmt"
	"strconv"
	"strings"
)

func Day15task2() {
	var str string = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

	strSplit := strings.Split(str, ",")

	boxes := make([][][2]string, 256)

	for _, item := range strSplit {
		label := item[:2]
		box := HASH(label)
		if strings.ContainsAny(item, "=") {
			focalLength := item[3:]
			addLens(box, label, focalLength, &boxes)
		}
		if strings.ContainsAny(item, "-") {
			removeLens(box, label, &boxes)
		}
	}
	fmt.Println(boxes)

	totalFocus := 0
	for i, box := range boxes {
		for j, lens := range box {
			focalLength, _ := strconv.Atoi(lens[1])
			focusPower := (i + 1) * (j + 1) * focalLength
			totalFocus += focusPower
		}
	}

	fmt.Println(totalFocus)
	fmt.Println("Hello, playground day15 is working ")
}

func isInArray(val string, array []string) bool {
	for _, item := range array {
		if item == val {
			return true
		}
	}
	return false
}

func HASH(str string) int {
	currentValue := 0
	for _, char := range str {
		ascii := int(char)
		currentValue += ascii
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func addLens(box int, label string, focalLength string, boxes *[][][2]string) {
	lens := [2]string{label, focalLength}
	for i, l := range (*boxes)[box] {
		if l[0] == label {
			(*boxes)[box][i] = lens
			return
		}
	}
	(*boxes)[box] = append((*boxes)[box], lens)
}

func removeLens(box int, label string, boxes *[][][2]string) {
	for i, lens := range (*boxes)[box] {
		if lens[0] == label {
			(*boxes)[box] = append((*boxes)[box][:i], (*boxes)[box][i+1:]...)
			return
		}
	}
}
