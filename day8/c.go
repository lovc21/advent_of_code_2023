package day8

import (
	"fmt"
	"strings"
)

func mytask() {
	var str string = `RL

	AAA = (BBB, CCC)
	BBB = (DDD, EEE)
	CCC = (ZZZ, GGG)
	DDD = (DDD, DDD)
	EEE = (EEE, EEE)
	GGG = (GGG, GGG)
	ZZZ = (ZZZ, ZZZ)`

	count := 0
	input := strings.Split(str, "\n")
	instructions := input[0]
	puzle := input[2:] // AAA = (BBB, CCC) it start here

	fmt.Println("Hello, playground day3 is working ")
	fmt.Println(instructions) // RL
	for _, v := range puzle {
		getreult := strings.Split(v, "=")
		getanwser := strings.TrimSpace(getreult[0])
		getpath := strings.TrimSpace(getreult[1])
		if strings.Contains(getanwser, "AAA") {
			for _, v := range instructions {
				if string(v) == "R" {
					getright := strings.Split(getpath, ",")
					getright[0] = strings.Trim(getright[0], "(")
					fmt.Println(getright[0])
					fmt.Println("R")
					count++
					processInstruction(instructions, getright[0], count)
					// call the funcion again using recursion
				} else if string(v) == "L" {
					getleft := strings.Split(getpath, ",")
					getleft[1] = strings.Trim(getleft[1], ")")
					fmt.Println(getleft[1])
					fmt.Println("L")
					count++
					processInstruction(instructions, getleft[1], count)
					// call the funcion again using recursion
				}
			}
		}
	}
}

func processInstruction(instructions, getanwser string, count int) {
	// Base case: stop recursion when "ZZZ" is found
	if strings.Contains(instructions, "ZZZ") {
		fmt.Println("ZZZ found")
		fmt.Println(count)
		return
	}

	// Recursive case: continue recursion
	for i, v := range instructions {
		if string(v) == "R" {
			getright := strings.Split(getanwser, ",")
			getright[0] = strings.Trim(getright[0], "(")
			fmt.Println(getright[0])
			fmt.Println("R")
			count++
			processInstruction(instructions, getright[0], count)
			// call the funcion again using recursion
		} else if string(v) == "L" {
			getleft := strings.Split(getanwser, ",")
			getleft[1] = strings.Trim(getleft[1], ")")
			fmt.Println(getleft[1])
			fmt.Println("L")
			count++
			processInstruction(instructions, getleft[1], count)
			// call the funcion again using recursion
		} else if i == len(instructions)-1 {
			// add RL to the end of instructions
			instructions = instructions + "RL"
		}
	}
}
