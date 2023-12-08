package day8

import (
	"fmt"
	"strings"
)

func Day8task1() {
	var str string = `RL

	AAA = (BBB, CCC)
	BBB = (DDD, EEE)
	CCC = (ZZZ, GGG)
	DDD = (DDD, DDD)
	EEE = (EEE, EEE)
	GGG = (GGG, GGG)
	ZZZ = (ZZZ, ZZZ)`

	puzzleMap := parsePuzzle(str)
	instructions := strings.TrimSpace(strings.Split(str, "\n")[0])

	fmt.Println("Processing puzzle with instructions: ", instructions)
	processInstructions(instructions, "AAA", puzzleMap, 0, 0, make(map[string]bool))

}

func parsePuzzle(str string) map[string][2]string {
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

func processInstructions(instructions, current string, puzzleMap map[string][2]string, count, instructionIndex int, history map[string]bool) bool {
	fmt.Println("Processing instructions:", instructions, "from", current)

	// Base case: stop recursion when "ZZZ" is found
	if current == "ZZZ" {
		fmt.Println("ZZZ found at step", count)
		return true
	}

	// Check if current key exists in the map
	next, exists := puzzleMap[current]
	if !exists {
		fmt.Println("Key not found in puzzle:", current)
		return false
	}

	historyKey := fmt.Sprintf("%s-%d", current, instructionIndex)

	// Check if this path has been visited before
	if _, visited := history[historyKey]; visited {
		fmt.Println("Loop detected, backtracking...")
		return false // Backtrack on loop detection
	}

	history[historyKey] = true

	// Get the current instruction based on instructionIndex
	instruction := instructions[instructionIndex%len(instructions)]
	var nextKey string
	if instruction == 'R' {
		nextKey = next[1]
	} else if instruction == 'L' {
		nextKey = next[0]
	} else {
		fmt.Println("Invalid instruction:", string(instruction))
		return false
	}

	fmt.Printf("Moving %c from %s to %s\n", instruction, current, nextKey)

	// Process the next instruction
	return processInstructions(instructions, nextKey, puzzleMap, count+1, (instructionIndex+1)%len(instructions), history)
}
