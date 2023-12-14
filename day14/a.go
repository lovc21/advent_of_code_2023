package day14

import (
	"fmt"
	"strings"
)

func Day14task1() {
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
	fmt.Println(matrix)

	intMatrix := make([][]int, len(matrix))
	for i, row := range matrix {
		intMatrix[i] = make([]int, len(row))
		for j, r := range row {
			intMatrix[i][j] = int(r)
		}
	}
	fmt.Println(intMatrix)

	sortedMatrix := sortColumns(intMatrix)
	fmt.Println(sortedMatrix)
	fmt.Println("len(sortedMatrix)", len(sortedMatrix))
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

func sortColumns(matrix [][]int) [][]int {
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
