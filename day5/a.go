package day5

import (
	"fmt"
	"strconv"
	"strings"
)

func Day5task3() {

	var str string = `seeds: 79 14 55 13

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

	parts := strings.Split(str, "\n\n")
	almanac := make(map[string][]int)
	// make a map for the input values
	for _, part := range parts {
		newpart := strings.Split(part, ":")

		if newpart[0] == "seeds" {
			for _, num := range strings.Fields(strings.TrimSpace(newpart[1])) {
				if num, err := strconv.Atoi(string(num)); err == nil {
					almanac["seeds"] = append(almanac["seeds"], num)
				}
			}
		} else if newpart[0] == "seed-to-soil map" {
			for _, num := range strings.Fields(strings.TrimSpace(newpart[1])) {
				if num, err := strconv.Atoi(string(num)); err == nil {
					almanac["seed-to-soil"] = append(almanac["seed-to-soil"], num)
				}
			}
		} else if newpart[0] == "soil-to-fertilizer map" {
			for _, num := range strings.Fields(strings.TrimSpace(newpart[1])) {
				if num, err := strconv.Atoi(string(num)); err == nil {
					almanac["soil-to-fertilizer"] = append(almanac["soil-to-fertilizer"], num)
				}
			}
		} else if newpart[0] == "fertilizer-to-water map" {
			for _, num := range strings.Fields(strings.TrimSpace(newpart[1])) {
				if num, err := strconv.Atoi(string(num)); err == nil {
					almanac["fertilizer-to-water"] = append(almanac["fertilizer-to-water"], num)
				}
			}
		} else if newpart[0] == "water-to-light map" {
			for _, num := range strings.Fields(strings.TrimSpace(newpart[1])) {
				if num, err := strconv.Atoi(string(num)); err == nil {
					almanac["water-to-light"] = append(almanac["water-to-light"], num)
				}
			}
		} else if newpart[0] == "light-to-temperature map" {
			for _, num := range strings.Fields(strings.TrimSpace(newpart[1])) {
				if num, err := strconv.Atoi(string(num)); err == nil {
					almanac["light-to-temperature"] = append(almanac["light-to-temperature"], num)
				}
			}
		} else if newpart[0] == "temperature-to-humidity map" {
			for _, num := range strings.Fields(strings.TrimSpace(newpart[1])) {
				if num, err := strconv.Atoi(string(num)); err == nil {
					almanac["temperature-to-humidity"] = append(almanac["temperature-to-humidity"], num)
				}
			}
		} else if newpart[0] == "humidity-to-location map" {
			for _, num := range strings.Fields(strings.TrimSpace(newpart[1])) {
				if num, err := strconv.Atoi(string(num)); err == nil {
					almanac["humidity-to-location"] = append(almanac["humidity-to-location"], num)
				}
			}
		}
	}

	fmt.Println("almanac done")

	seedToLocation := make(map[int]int)

	for _, seed := range almanac["seeds"] {
		fmt.Println("seed", seed)

		currentValue := seed
		order := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

		for _, key := range order {
			fmt.Println("key", key)
			mapping := Mappings(almanac[key])
			currentValue = mapping[currentValue]
		}

		// Store the final location for each seed
		seedToLocation[seed] = currentValue
	}

	// Now find the seed with the lowest location number
	minLocation := 100
	minSeed := -1
	for seed, location := range seedToLocation {
		if location < minLocation {
			minLocation = location
			minSeed = seed
		}
	}

	if minSeed != -1 {
		fmt.Printf("Seed with lowest location number: Seed %d, Location %d\n", minSeed, minLocation)
	} else {
		fmt.Println("No seeds found.")
	}

	fmt.Println(almanac)

}

func Mappings(values []int) map[int]int {
	mapping := make(map[int]int)

	// make a map for the input values 1234567889 -> 1234567889
	for i := 0; i < 100; i++ {
		mapping[i] = i
		fmt.Println("initial mapping", mapping)
	}

	// apply the mappings to the map
	for i := 0; i < len(values); i += 3 {
		destStart := values[i]
		srcStart := values[i+1]
		length := values[i+2]

		for j := 0; j < length; j++ {
			mapping[srcStart+j] = destStart + j
		}
		fmt.Println("mapping", mapping)
	}

	fmt.Println(mapping)

	return mapping
}
