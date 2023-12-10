package day10

import (
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

var grid [][]rune
var visited map[Point]bool
var start Point

func main() {
	Day10task2()
}

func Day10task2() {
	// Your maze string here
	var str string = `your_maze_string`
	lines := strings.Split(str, "\n")

	grid = make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = []rune(line)
		for x, c := range grid[y] {
			if c == 'S' {
				start = Point{x, y}
			}
		}
	}

	visited = make(map[Point]bool)
	longestPath := make([]Point, 0)

	// Check in all four directions
	directions := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range directions {
		visited = make(map[Point]bool)
		path, _ := CheckPos(Point{start.x + dir.x, start.y + dir.y}, []Point{})
		if len(path) > len(longestPath) {
			longestPath = path
		}
	}

	count := countInsideSpaces(longestPath)
	fmt.Println("part2: ", count)
}

func CheckPos(p Point, path []Point) ([]Point, int) {
	if visited[p] || p.x < 0 || p.y < 0 || p.x >= len(grid[0]) || p.y >= len(grid) || grid[p.y][p.x] == '.' {
		return path, len(path)
	}

	visited[p] = true
	newPath := append([]Point(nil), path...)
	newPath = append(newPath, p)

	var maxLength int
	var maxPath []Point

	switch grid[p.y][p.x] {
	case '|', 'L', 'F':
		path1, length1 := CheckPos(Point{p.x, p.y - 1}, newPath)
		path2, length2 := CheckPos(Point{p.x, p.y + 1}, newPath)
		if length1 > length2 {
			maxPath, maxLength = path1, length1
		} else {
			maxPath, maxLength = path2, length2
		}
	case '-', 'J', '7':
		path1, length1 := CheckPos(Point{p.x - 1, p.y}, newPath)
		path2, length2 := CheckPos(Point{p.x + 1, p.y}, newPath)
		if length1 > length2 {
			maxPath, maxLength = path1, length1
		} else {
			maxPath, maxLength = path2, length2
		}
	case 'S':
		return newPath, len(newPath)
	}

	return maxPath, maxLength
}

var maze [][]int
var visited [][]bool

func isValid(x, y int) bool {
	if x < 0 || y < 0 || x >= len(maze) || y >= len(maze[0]) {
		return false
	}
	if visited[x][y] || maze[x][y] == 1 {
		return false
	}
	return true
}

func floodFill(x, y int) int {
	if !isValid(x, y) {
		return 0
	}
	visited[x][y] = true
	return 1 + floodFill(x+1, y) + floodFill(x-1, y) + floodFill(x, y+1) + floodFill(x, y-1)
}

func countInsideSpaces(path []Point) int {
	visited = make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}
	count := 0
	for _, point := range path {
		count += floodFill(point.x, point.y)
	}
	return count
}
