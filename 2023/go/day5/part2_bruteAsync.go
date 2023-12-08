package day5

import (
	"adventUtil"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Part2BruteAsyncSolution(input []string) string {
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

	c1 := make(chan int)
	for _, ss := range seedStructs {
		go func(ss2 seedRange) {
			found_min := mapSeedImprov(ss2.start, mappings)
			for seed := ss2.start; seed <= ss2.end; seed++ {
				next_seed := seed
				for _, mappings := range all_mappings {
					next_seed = mapSeedImprov(next_seed, mappings)
				}
				if next_seed < found_min {
					found_min = next_seed
				}
			}
			c1 <- found_min
		}(ss)
	}

	new_minimum := 281453544
	for range seedStructs {
		next_val := <-c1
		fmt.Println("next val received:", next_val)
		if next_val < new_minimum {
			new_minimum = next_val
			fmt.Println("++++ Is Better")
		}
	}

	fmt.Println("END time:", time.Now())
	return strconv.Itoa(new_minimum)
	// 53.405 parallelized
}
