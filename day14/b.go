package day14

import (
	"fmt"
	"strings"
)

func Day14task2() {
	var str string = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

	lines := strings.Split(str, "\n")

	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}

	intMatrix := make([][]int, len(matrix))
	for i, row := range matrix {
		intMatrix[i] = make([]int, len(row))
		for j, r := range row {
			intMatrix[i][j] = int(r)
		}
	}

	// do tihs loping 1000 times
	numberOfCycles := 1000
	var sortedMatrix [][]int = intMatrix
	for i := 0; i < numberOfCycles; i++ {
		sortedMatrix = SortColumnsAscending(sortedMatrix)
		sortedMatrix = sortRowsRight(sortedMatrix)
		sortedMatrix = sortColumnsDescending(sortedMatrix)
		sortedMatrix = sortRowsLeft(sortedMatrix)
	}

	count := 0
	countRows := 0
	for i := 0; i < len(sortedMatrix); i++ {
		for j := 0; j < len(sortedMatrix[0]); j++ {
			if sortedMatrix[i][j] == 79 {
				count++
			}
		}
		countRows = countRows + (count * (len(sortedMatrix[0]) - i))
		count = 0
	}
	fmt.Println("countRows", countRows)
	fmt.Println("Hello, playground day14 is working ")

}

func SortColumnsAscending(matrix [][]int) [][]int {
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			if matrix[i][j] == 35 {
				continue
			}

			for k := i + 1; k < len(matrix); k++ {
				if matrix[k][j] == 35 {
					break
				}

				if matrix[k][j] > matrix[i][j] {
					matrix[i][j], matrix[k][j] = matrix[k][j], matrix[i][j]
				}
			}
		}
	}
	return matrix
}

func sortColumnsDescending(matrix [][]int) [][]int {
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			if matrix[i][j] == 35 {
				continue
			}

			lowestPosition := i
			for k := i + 1; k < len(matrix); k++ {
				if matrix[k][j] == 35 {
					break
				}
				if matrix[k][j] < matrix[i][j] {
					lowestPosition = k
				}
			}
			if lowestPosition != i {
				matrix[i][j], matrix[lowestPosition][j] = matrix[lowestPosition][j], matrix[i][j]
			}
		}
	}
	return matrix
}

func sortRowsRight(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 35 {
				continue
			}

			furthestPosition := j
			for k := j + 1; k < len(matrix[i]); k++ {
				if matrix[i][k] == 35 {
					break
				}
				if matrix[i][k] < matrix[i][j] {
					furthestPosition = k
				}
			}

			if furthestPosition != j {
				matrix[i][j], matrix[i][furthestPosition] = matrix[i][furthestPosition], matrix[i][j]
			}
		}
	}
	return matrix
}

func sortRowsLeft(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			if matrix[i][j] == 35 {
				continue
			}

			furthestPosition := j
			for k := j - 1; k >= 0; k-- {
				if matrix[i][k] == 35 {
					break
				}
				if matrix[i][k] < matrix[i][j] {
					furthestPosition = k
				}
			}
			if furthestPosition != j {
				matrix[i][j], matrix[i][furthestPosition] = matrix[i][furthestPosition], matrix[i][j]
			}
		}
	}
	return matrix
}
