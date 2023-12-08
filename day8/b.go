package day8

import (
	"fmt"
	"regexp"
	"strings"
)

func Day8task2() {
	str := `LR

    11A = (11B, XXX)
    11B = (XXX, 11Z)
    11Z = (11B, XXX)
    22A = (22B, XXX)
    22B = (22C, 22C)
    22C = (22Z, 22Z)
    22Z = (22B, 22B)
    XXX = (XXX, XXX)`

	lines := strings.Split(str, "\n")
	instructions := lines[0]
	r, _ := regexp.Compile("([0-9A-Z]+) = \\(([0-9A-Z]+), ([0-9A-Z]+)\\)")

	puzzleMap := make(map[string][]string)
	for _, line := range lines[2:] {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		matches := r.FindStringSubmatch(line)
		if len(matches) >= 4 {
			par, left, right := matches[1], matches[2], matches[3]
			puzzleMap[par] = []string{left, right}
		}
	}

	var current []string
	for key, _ := range puzzleMap {
		if strings.HasSuffix(key, "A") { // Assuming the start nodes end with 'A'
			current = append(current, key)
		}
	}

	var lengths []int
	for _, node := range current {
		lengths = append(lengths, loop_n(node, instructions, puzzleMap))
	}

	fmt.Println("Lengths:", lengths)
}

func loop_n(current string, instructions string, puzzleMap map[string][]string) int {
	count := 0
	for current != "Z" { // Assuming "Z" is the termination condition
		instruction := instructions[count%len(instructions)]
		if instruction == 'R' {
			current = puzzleMap[current][1]
		} else {
			current = puzzleMap[current][0]
		}
		count++
	}
	return count
}
