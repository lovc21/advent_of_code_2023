package day17

import (
	"container/heap"
	"fmt"
	"strings"
)

func Day17task2() {
	str := `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

	lines := strings.Split(str, "\n")
	grid := [][]int{}

	for _, line := range lines {
		line = strings.TrimSpace(line) // Remove leading and trailing whitespace
		row := []int{}
		for _, char := range line {
			row = append(row, int(char-'0'))
		}
		grid = append(grid, row)
	}

	seen := make(map[[5]int]bool)
	pq := &minPath{path{0, 0, 0, 0, 0, 0}}
	heap.Init(pq)

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for pq.Len() > 0 {
		p := heap.Pop(pq).(path)
		r, c, dr, dc, n := p.r, p.c, p.dr, p.dc, p.n

		if r == len(grid)-1 && c == len(grid[0])-1 && n >= 4 {
			fmt.Println(p.hl)
			break
		}

		if seen[[5]int{r, c, dr, dc, n}] {
			continue
		}
		seen[[5]int{r, c, dr, dc, n}] = true

		if n < 10 && (dr != 0 || dc != 0) {
			nr := r + dr
			nc := c + dc
			if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) {
				heap.Push(pq, path{p.hl + grid[nr][nc], nr, nc, dr, dc, n + 1})
			}
		}

		if n >= 4 || (dr == 0 && dc == 0) {
			for _, direction := range directions {
				ndr, ndc := direction[0], direction[1]
				if (ndr != dr || ndc != dc) && (ndr != -dr || ndc != -dc) {
					nr := r + ndr
					nc := c + ndc
					if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) {
						heap.Push(pq, path{p.hl + grid[nr][nc], nr, nc, ndr, ndc, 1})
					}
				}
			}
		}
	}

	fmt.Println("Hello, playground day17 is working ")
}
