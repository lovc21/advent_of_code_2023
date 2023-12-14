package day13

import (
	"fmt"
	"strings"
)

func Day13task1() {
	var str string = `#.##..##.
	..#.##.#.
	##......#
	##......#
	..#.##.#.
	..##..##.
	#.#.##.#.

	#...##..#
	#....#..#
	..##..###
	#####.##.
	#####.##.
	..##..###
	#....#..#`

	toltalsum := 0

	// Normalize line endings and trim the string
	str = strings.ReplaceAll(str, "\r\n", "\n")
	str = strings.TrimSpace(str)

	// Split the string on double newlines
	parts := strings.Split(str, "\n\n")

	if len(parts) != 2 {
		fmt.Println("Unexpected number of parts, expected 2.")
		return
	}

	// Process the first part for vertical storage
	firstPart := strings.Split(strings.TrimSpace(parts[0]), "\n")
	var maxWidth int
	for _, line := range firstPart {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	MapColumns := make([][]rune, maxWidth)
	for x := range MapColumns {
		MapColumns[x] = make([]rune, len(firstPart))
		for y := range MapColumns[x] {
			MapColumns[x][y] = ' ' // Initialize with spaces
		}
	}

	for y, line := range firstPart {
		for x, char := range strings.TrimSpace(line) {
			if x < maxWidth {
				MapColumns[x][y] = char
			}
		}
	}

	fmt.Println("Vertical storage of second part:")
	startIndexColumns, endIndex := findMiddleRows(MapColumns)
	fmt.Printf("Middle row(s) are at index(es): %d and %d\n", startIndexColumns+1, endIndex+1)

	// Process the second part for horizontal storage
	secondPart := strings.Split(strings.TrimSpace(parts[1]), "\n")
	MapRows := make([][]rune, len(secondPart))
	for y, line := range secondPart {
		MapRows[y] = []rune(strings.TrimSpace(line))
	}

	// Find the middle row(s)
	fmt.Println("Horizontal storage of second part:")
	startIndexRows, endIndex := findMiddleRows(MapRows)
	fmt.Printf("Middle row(s) are at index(es): %d and %d\n", startIndexRows+1, endIndex+1)

	startIndexRows = (startIndexRows + 1) * 100
	toltalsum = startIndexRows + 5
	fmt.Println(toltalsum)
}

// Function to check if two rows are the same
func areRowsSame(row1, row2 []rune) bool {
	for i := range row1 {
		if row1[i] != row2[i] {
			return false
		}
	}
	return true
}

// Function to find the middle row(s) after discarding non-matching outer rows
func findMiddleRows(MapRows [][]rune) (int, int) {
	start := 0
	end := len(MapRows) - 1

	for start <= end {
		if !areRowsSame(MapRows[start], MapRows[end]) {
			start++
			end--
		} else {
			// Check if pointers have met or crossed
			if start == end || start+1 == end {
				break
			}
			start++
			end--
		}
	}

	return start, end
}
