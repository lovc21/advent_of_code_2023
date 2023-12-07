package day5

import (
	"fmt"
	"strconv"
	"strings"
)

func Day5task1() {
	var input = `seeds: 79 14 55 13

	seed-to-soil map:
	50 98 2
	52 50 48
	
	soil-to-fertilizer map:
	0 15 37
	37 52 2
	39 0 15
	
	fertilizer-to-water map:
	49 53 8
	0 11 42
	42 0 7
	57 7 4
	
	water-to-light map:
	88 18 7
	18 25 70
	
	light-to-temperature map:
	45 77 23
	81 45 19
	68 64 13
	
	temperature-to-humidity map:
	0 69 1
	1 0 69
	
	humidity-to-location map:
	60 56 37
	56 93 4`

	lines := strings.Split(input, "\n")

	// Extracting the seeds
	seedStrs := strings.Fields(strings.Split(lines[0], ":")[1])
	var seeds []int
	for _, s := range seedStrs {
		seed, _ := strconv.Atoi(s)
		seeds = append(seeds, seed)
	}

	var maps [][][]int
	var currentMap [][]int

	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.Contains(line, "map:") {
			if currentMap != nil {
				maps = append(maps, currentMap)
			}
			currentMap = [][]int{}
		} else {
			parts := strings.Fields(line)
			var parsedParts []int
			for _, part := range parts {
				num, _ := strconv.Atoi(part)
				parsedParts = append(parsedParts, num)
			}
			currentMap = append(currentMap, parsedParts)
		}
	}
	maps = append(maps, currentMap) // Append the last map

	for _, m := range maps {
		for i := range seeds {
			for _, mapping := range m {
				offset := mapping[0] - mapping[1]
				if mapping[1] <= seeds[i] && seeds[i] <= mapping[1]+mapping[2] {
					seeds[i] += offset
					break
				}
			}
		}
	}

	// Find the minimum seed value
	minSeed := seeds[0]
	for _, s := range seeds[1:] {
		if s < minSeed {
			minSeed = s
		}
	}

	fmt.Println("The lowest location number is:", minSeed)
}
