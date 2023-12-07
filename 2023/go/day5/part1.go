package day5

import (
	"adventUtil"
	"slices"
	"strconv"
	"strings"
)

var digits = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func isDigit(char rune) bool {
	return slices.Contains(digits, char)
}

func minInt(input []int) int {
	minimum := input[0]
	for _, num := range input {
		if num < minimum {
			minimum = num
		}
	}
	return minimum
}

type rangeMapping struct {
	destination_start int
	source_start      int
	distance          int
}

func mapSeed(seed int, mappings []rangeMapping) int {
	for _, mapping := range mappings {
		if seed >= mapping.source_start &&
			seed < mapping.source_start+mapping.distance {
			distance := seed - mapping.source_start
			return mapping.destination_start + distance
		}
	}
	return seed
}

func mapSeeds(seeds []int, mappings []rangeMapping) []int {
	new_seeds := []int{}
	for _, seed := range seeds {
		new_seeds = append(new_seeds, mapSeed(seed, mappings))
	}
	return new_seeds
}

/*
seed 2 fertalizer -> water should be 49

*/

func Part1Solution(input []string) string {
	seed_split := strings.Split(input[0], ":")
	seeds := adventUtil.ConvertStringsToInts(strings.Split(seed_split[1], " "))
	// fmt.Println("next seeds", seeds)

	mappings := []rangeMapping{}
	for index := 2; index < len(input); index++ {
		if input[index] == "" {
			// received all mapping
			// fmt.Println("next mappings", mappings)
			seeds = mapSeeds(seeds, mappings)
			// fmt.Println("next seeds", seeds)
		} else if !isDigit(rune(input[index][0])) {
			// start of the next mapping
			mappings = []rangeMapping{}
		} else {
			// recived range Mapping
			nums := adventUtil.ConvertStringsToInts(strings.Split(input[index], " "))
			mappings = append(mappings, rangeMapping{nums[0], nums[1], nums[2]})
		}
	}

	seeds = mapSeeds(seeds, mappings)
	// fmt.Println("next seeds", seeds)

	return strconv.Itoa(minInt(seeds))
}
