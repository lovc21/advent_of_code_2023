package day10

import (
	"fmt"
	"strings"

	"modernc.org/mathutil"
)

type Point struct {
	x, y int
}

var grid = make([][]rune, 0)
var visited = make(map[Point]bool)
var start = Point{-1, -1}

func Day10task1() {
	var str string = `
	..F7.
	.FJ|.
	SJ.L7
	|F--J
	LJ...
`
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	grid = make([][]rune, len(lines)-2)
	for y, line := range lines[1 : len(lines)-1] {
		grid[y] = []rune(line)
		for x, c := range grid[y] {
			if c == 'S' {
				start = Point{x, y}
			}
		}
	}

	p := start
	r := checkPos(Point{p.x + 1, p.y}, 0)
	visited = make(map[Point]bool)
	l := checkPos(Point{p.x - 1, p.y}, 0)
	visited = make(map[Point]bool)
	d := checkPos(Point{p.x, p.y + 1}, 0)
	visited = make(map[Point]bool)
	u := checkPos(Point{p.x, p.y - 1}, 0)
	fmt.Println(mathutil.MaxVal(r, l, d, u) / 2) // Divide by 2 because each step is counted twice (going and coming back)
}

func checkPos(p Point, step int) int {
	if visited[p] || p.x < 0 || p.y < 0 || p.x >= len(grid[0]) || p.y >= len(grid) || grid[p.y][p.x] == '.' {
		return step
	}
	step++
	visited[p] = true

	switch grid[p.y][p.x] {
	case '|':
		d := checkPos(Point{p.x, p.y + 1}, step)
		u := checkPos(Point{p.x, p.y - 1}, step)
		return mathutil.Max(d, u)
	case '-':
		r := checkPos(Point{p.x + 1, p.y}, step)
		l := checkPos(Point{p.x - 1, p.y}, step)
		return mathutil.Max(r, l)
	case 'F':
		r := checkPos(Point{p.x + 1, p.y}, step)
		d := checkPos(Point{p.x, p.y + 1}, step)
		return mathutil.Max(r, d)
	case '7':
		l := checkPos(Point{p.x - 1, p.y}, step)
		d := checkPos(Point{p.x, p.y + 1}, step)
		return mathutil.Max(l, d)
	case 'J':
		l := checkPos(Point{p.x - 1, p.y}, step)
		u := checkPos(Point{p.x, p.y - 1}, step)
		return mathutil.Max(l, u)
	case 'L':
		r := checkPos(Point{p.x + 1, p.y}, step)
		u := checkPos(Point{p.x, p.y - 1}, step)
		return mathutil.Max(r, u)
	case 'S':
		fmt.Println("found", step)
		break
	}

	return step
}
