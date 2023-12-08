package day5

import (
	"adventUtil"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type rangeMappingImprov struct {
	destination_start int
	source_start      int
	source_end        int
}

func mapSeedImprov(seed int, mappings []rangeMappingImprov) int {
	for _, mapping := range mappings {
		if seed >= mapping.source_start &&
			seed < mapping.source_end {
			distance := seed - mapping.source_start
			return mapping.destination_start + distance
		}
	}
	return seed
}

func Part2BruteSolution(input []string) string {
	fmt.Println("Start", time.Now())
	seed_split := strings.Split(input[0], ":")
	seeds := adventUtil.ConvertStringsToInts(strings.Split(seed_split[1], " "))
	// fmt.Println("next seeds", seeds)
	seedStructs := []seedRange{}
	for i := 0; i < len(seeds); i += 2 {
		seedStructs = append(seedStructs, seedRange{seeds[i], seeds[i] + seeds[i+1] - 1})
	}
	sum := 0
	for _, ss := range seedStructs {
		sum += ss.end - ss.start
	}
	all_mappings := [][]rangeMappingImprov{}

	// setup mappings
	mappings := []rangeMappingImprov{}
	for index := 2; index < len(input); index++ {
		if input[index] == "" {
			// received all mapping
			// fmt.Println("next mappings", mappings)
			all_mappings = append(all_mappings, mappings)
			// fmt.Println("next seeds", seeds)
		} else if !isDigit(rune(input[index][0])) {
			// start of the next mapping
			mappings = []rangeMappingImprov{}
		} else {
			// recived range Mapping
			nums := adventUtil.ConvertStringsToInts(strings.Split(input[index], " "))
			mappings = append(mappings, rangeMappingImprov{nums[0], nums[1], nums[1] + nums[2]})
		}
	}
	all_mappings = append(all_mappings, mappings)

	found_min := mapSeedImprov(seedStructs[0].start, mappings)
	iterated := 0
	for _, ss := range seedStructs {
		for seed := ss.start; seed <= ss.end; seed++ {
			next_seed := seed
			for _, mappings := range all_mappings {
				next_seed = mapSeedImprov(next_seed, mappings)
			}
			if next_seed < found_min {
				fmt.Println(next_seed)
				found_min = next_seed
			}

			iterated += 1
			if iterated%10000000 == 0 {
				fmt.Println(float64(iterated)/float64(sum), iterated, "time:", time.Now())
			}
		}
	}

	fmt.Println("END time:", time.Now())
	return strconv.Itoa(found_min)
	// 324.810267865 slight optimization
}
