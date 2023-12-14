package day12

import (
	"fmt"
	"strconv"
	"strings"
)

func Day12task1() {
	totalsum := 0

	var str string = `???.### 1,1,3
	.??..??...?##. 1,1,3
	?#?#?#?#?#?#?#? 1,3,1,6
	????.#...#... 4,1,1
	????.######..#####. 1,6,5
	?###???????? 3,2,1`

	lines := strings.Split(str, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		splitForSprings := strings.Split(line, " ")
		//fmt.Println(splitForSprings[0])
		cfg := splitForSprings[0]
		// Create a tuple tip of ints
		numbers := strings.Split(splitForSprings[1], ",")

		nums := make([]int, len(numbers))

		for i, num := range numbers {
			nums[i], _ = strconv.Atoi(num)
		}

		totalsum += countTheBrokenSprings(cfg, nums)
		fmt.Println(totalsum)
		fmt.Println(nums)
	}

	fmt.Println(totalsum)
}

func countTheBrokenSprings(cfg string, nums []int) int {
	if cfg == "" {
		if len(nums) == 0 {
			return 1 // No more cfg and nums, valid arrangement
		}
		return 0 // No more cfg but nums left, invalid arrangement
	}

	if len(nums) == 0 {
		if strings.Contains(cfg, "#") {
			return 0 // No more nums but cfg contains '#', invalid arrangement
		}
		return 1 // No more nums and cfg doesn't contain '#', valid arrangement
	}

	result := 0
	firstChar := cfg[0]

	// Case 1: Current character is '.' or '?', move to next character
	if firstChar == '.' || firstChar == '?' {
		result += countTheBrokenSprings(cfg[1:], nums)
	}

	// Case 2: Current character is '#' or '?', and the segment meets the condition
	if firstChar == '#' || firstChar == '?' {
		if len(nums) > 0 && nums[0] <= len(cfg) && !strings.Contains(cfg[:nums[0]], ".") {
			if nums[0] == len(cfg) || (nums[0] < len(cfg) && cfg[nums[0]] != '#') {
				result += countTheBrokenSprings(cfg[nums[0]:], nums[1:])
			}
		}
	}

	return result
}
