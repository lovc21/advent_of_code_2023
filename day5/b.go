package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lovc21/advent_of_code_2024/util"
)

func Day5task2(input string) {
	lines := strings.Split(input, "\n")

	// Extracting the seeds and creating initial ranges
	seedStrs := strings.Fields(strings.Split(lines[0], ":")[1])
	var seedRanges [][2]int
	for i := 0; i < len(seedStrs); i += 2 {
		start, _ := strconv.Atoi(seedStrs[i])
		length, _ := strconv.Atoi(seedStrs[i+1])
		seedRanges = append(seedRanges, [2]int{start, start + length - 1})
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
		var newSeedRanges [][2]int
		for _, r := range seedRanges {
			for _, mapping := range m {
				offset := mapping[0] - mapping[1]
				mapLow := mapping[1]
				mapHigh := mapping[1] + mapping[2] - 1

				srLow := r[0]
				srHigh := r[1]

				if srLow <= mapHigh && mapLow <= srHigh {
					// Middle part affect
					low := util.IntMin(srLow, mapLow) + offset
					high := util.IntMax(srHigh, mapHigh) + offset
					newSeedRanges = append(newSeedRanges, [2]int{low, high})
				}

				// Left part reduction
				if srLow < mapLow {
					newSeedRanges = append(newSeedRanges, [2]int{srLow, util.IntMin(srHigh, mapLow-1)})
				}

				// Right part append
				if mapHigh < srHigh {
					newSeedRanges = append(newSeedRanges, [2]int{util.IntMin(mapHigh+1, srLow), srHigh})
				}
			}
		}
		seedRanges = newSeedRanges
	}

	// Find the minimum start value
	minStart := int(^uint(0) >> 1) // Max int value
	for _, rangePair := range seedRanges {
		start := rangePair[0]
		if start < minStart {
			minStart = start
		}
	}
	fmt.Println("The lowest location number is:", minStart)
}
