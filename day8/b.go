package day8

import (
	"fmt"
	"strings"
)

func Day8task2() {
	var str string = `LR

	11A = (11B, XXX)
	11B = (XXX, 11Z)
	11Z = (11B, XXX)
	22A = (22B, XXX)
	22B = (22C, 22C)
	22C = (22Z, 22Z)
	22Z = (22B, 22B)
	XXX = (XXX, XXX)`

	puzzleMap := parsePuzzletask2(str)
	instructions := strings.TrimSpace(strings.Split(str, "\n")[0])

	fmt.Println("Processing puzzle with instructions: ", instructions)
	recursionTask2(instructions, "AAA", "AAA", puzzleMap, 0, 0, make(map[string]bool))

}

func parsePuzzletask2(str string) map[string][2]string {
	puzzleMap := make(map[string][2]string)
	lines := strings.Split(str, "\n")

	for _, line := range lines[1:] {
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue // skip invalid lines
		}
		key := strings.TrimSpace(parts[0])
		values := strings.Split(strings.Trim(parts[1], "() "), ", ")
		if len(values) != 2 {
			continue // skip invalid entries
		}
		puzzleMap[key] = [2]string{values[0], values[1]}
	}

	return puzzleMap
}

func recursionTask2(instructions, current1 string, current2 string, puzzleMap map[string][2]string, count, instructionIndex int, history map[string]bool) bool {
	return true
}
