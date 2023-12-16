package day15

import (
	"fmt"
	"strings"
)

func Day15task1() {
	var str string = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

	strSplit := strings.Split(str, ",")
	toltalsum := 0
	for _, itme := range strSplit {
		currentValue := 0
		for _, char := range itme {
			ascii := int(char)
			currentValue += ascii
			currentValue *= 17
			currentValue %= 256
		}
		toltalsum += currentValue
		fmt.Println("Current Value for", itme, ":", currentValue)
	}

	fmt.Println(strSplit)
	fmt.Println(toltalsum)
	fmt.Println("Hello, playground day15 is working ")
}
