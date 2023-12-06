package day6

import (
	"fmt"
	"strconv"
	"strings"
)

func Day6task1() {
	var input string = `Time:        49     87     78     95
	Distance:   356   1378   1502   1882`

	data := make(map[string][]int)
	lines := strings.Split(input, "\n")
	getdata := strings.Split(lines[0], ":")

	data["Time"] = append(data["Time"])
	data["Distance"] = append(data["Distance"])
	for _, num := range strings.Fields(strings.TrimSpace(getdata[1])) {
		if num, err := strconv.Atoi(string(num)); err == nil {
			data["Time"] = append(data["Time"], num)
		}
	}

	getdata = strings.Split(lines[1], ":")
	for _, num := range strings.Fields(strings.TrimSpace(getdata[1])) {
		if num, err := strconv.Atoi(string(num)); err == nil {
			data["Distance"] = append(data["Distance"], num)
		}
	}

	fmt.Println(data)

	totalSum := 1
	for index, time := range data["Time"] {
		totalSumLocal := 0
		for i := 0; i < time; i++ {
			totalTime := (time - i) * i
			if totalTime > data["Distance"][index] {
				fmt.Println(data["Distance"][index])
				totalSumLocal++
			}
		}
		fmt.Println(totalSumLocal)
		totalSum *= totalSumLocal
	}

	fmt.Println(totalSum)

	fmt.Println("Hello, playground day6 is working ")

}
