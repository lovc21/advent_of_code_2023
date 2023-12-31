package day11

import (
	"fmt"
	"strings"
)

func Day11task1() {
	var str string = `
	...#......
	.......#..
	#.........
	..........
	......#...
	.#........
	.........#
	..........
	.......#..
	#...#.....`

	lines := strings.Split(str, "\n")
	grid := make([][]rune, len(lines)-1)
	markers := make(map[int]Point)
	nextMarkerID := 0 // Start assigning galaxy markers from 0

	// Parse the grid and dynamically assign markers
	for y, line := range lines {
		if line == "" {
			continue
		}
		grid[y-1] = make([]rune, len(line))
		for x, cell := range line {
			grid[y-1][x] = cell
			if cell == '#' { // Assuming galaxies are marked with '#'
				markers[nextMarkerID] = Point{x, y - 1}
				nextMarkerID++
			}
		}
	}

	totalDistance := 0
	for i := 0; i < nextMarkerID; i++ {
		for j := i + 1; j < nextMarkerID; j++ {
			dist := BFS(grid, markers[i], markers[j])
			if dist != -1 {
				totalDistance += dist
			}
		}
	}

	fmt.Println("Total distance:", totalDistance)
}

type Point struct {
	x, y int
}

type QueueNode struct {
	point Point
	dist  int
}

var rowNum = []int{-1, 0, 0, 1, -1, -1, 1, 1}
var colNum = []int{0, -1, 1, 0, -1, 1, -1, 1}

// Check if a cell is valid for movement
func isValid(mat [][]rune, visited [][]bool, row, col int) bool {
	return row >= 0 && row < len(mat) && col >= 0 && col < len(mat[0]) && !visited[row][col]
}

// BFS to find shortest path in a grid
func BFS(mat [][]rune, src, dest Point) int {
	if src == dest {
		return 0
	}

	visited := make([][]bool, len(mat))
	for i := range visited {
		visited[i] = make([]bool, len(mat[0]))
	}

	q := []QueueNode{{src, 0}}
	visited[src.y][src.x] = true

	for len(q) > 0 {
		curr := q[0]
		point := curr.point

		q = q[1:]

		for i := 0; i < 4; i++ {
			row := point.y + rowNum[i]
			col := point.x + colNum[i]

			if isValid(mat, visited, row, col) {
				visited[row][col] = true
				adjCell := QueueNode{Point{col, row}, curr.dist + 1}

				if adjCell.point == dest {
					return adjCell.dist
				}

				q = append(q, adjCell)
			}
		}
	}

	return -1
}
