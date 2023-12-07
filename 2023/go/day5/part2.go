package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type seedRange struct {
	start int
	end   int
}

// return values: unmapped seed ranges, mapped seed ranges (should no longer consider for mapping)
func mapSeedRangeForMapping(sr seedRange, mapping rangeMapping) ([]seedRange, []seedRange) {
	source_end := mapping.source_start + mapping.distance        // exclusive
	dest_end := mapping.destination_start + mapping.distance - 1 // inclusive
	// the seed range is before the given mapping so it is a 1:1
	if sr.end < mapping.source_start {
		return []seedRange{
			seedRange{sr.start, sr.end},
		}, []seedRange{}
	}

	// seed range overlaps start but not end
	if sr.end >= mapping.source_start && sr.end < source_end &&
		sr.start < mapping.source_start {
		return []seedRange{
				seedRange{sr.start, mapping.source_start - 1}, // before mapping
			},
			[]seedRange{
				seedRange{mapping.destination_start, mapping.destination_start + sr.end - mapping.source_start}, // overlap beginning of mapping
			}
	}

	// seed range is contained by mapping
	if sr.start >= mapping.source_start && sr.end < source_end {
		return []seedRange{},
			[]seedRange{
				seedRange{
					mapping.destination_start + sr.start - mapping.source_start,
					mapping.destination_start + sr.end - mapping.source_start,
				},
			}
	}

	// seed range is overflowing end of mapping
	if sr.start < source_end && sr.start >= mapping.source_start && sr.end >= source_end {
		return []seedRange{
				seedRange{source_end, sr.end}, // not mapped overflow
			},
			[]seedRange{
				seedRange{mapping.destination_start + sr.start - mapping.source_start, dest_end},
			}
	}

	// mapping is contained by seed range
	if sr.start < mapping.source_start && sr.end >= source_end {
		return []seedRange{
				seedRange{sr.start, mapping.source_start - 1}, // the before mapping
				seedRange{source_end, sr.end},                 // the after mapping
			},
			[]seedRange{
				seedRange{mapping.destination_start, dest_end}, // the full mapping range
			}
	}

	// seed range is entirely above mapping
	if sr.start >= source_end {
		return []seedRange{
			seedRange{sr.start, sr.end},
		}, []seedRange{}
	}
	fmt.Println("UHOH")
	fmt.Printf("%+v |  %+v\n", sr, mapping)
	os.Exit(1)
	return []seedRange{}, []seedRange{}
}

func mapSeedRange(sr seedRange, mappings []rangeMapping) []seedRange {
	new_seed_ranges := []seedRange{sr}
	end_seed_ranges := []seedRange{}
	for _, mapping := range mappings {
		leftover_new_seed_ranges := []seedRange{}
		for _, next_sr := range new_seed_ranges {
			continuing, finished := mapSeedRangeForMapping(next_sr, mapping)
			leftover_new_seed_ranges = append(leftover_new_seed_ranges, continuing...)
			end_seed_ranges = append(end_seed_ranges, finished...)
		}
		new_seed_ranges = leftover_new_seed_ranges
	}
	return append(new_seed_ranges, end_seed_ranges...)
}

func mapSeedRanges(seedRanges []seedRange, mappings []rangeMapping) []seedRange {
	new_seed_ranges := []seedRange{}
	for _, sr := range seedRanges {
		new_seed_ranges = append(new_seed_ranges, mapSeedRange(sr, mappings)...)
	}
	return new_seed_ranges
}

func Part2Solution(input []string) string {
	seed_split := strings.Split(input[0], ":")
	seeds := convertStringsToInts(strings.Split(seed_split[1], " "))
	seedStructs := []seedRange{}
	for i := 0; i < len(seeds); i += 2 {
		seedStructs = append(seedStructs, seedRange{seeds[i], seeds[i] + seeds[i+1] - 1})
	}
	// fmt.Println("next seeds", seedStructs)

	mappings := []rangeMapping{}
	for index := 2; index < len(input); index++ {
		if input[index] == "" {
			// received all mapping
			// fmt.Println("next mappings", mappings)
			seedStructs = mapSeedRanges(seedStructs, mappings)
			// fmt.Println("next seeds", index, len(seedStructs))
		} else if !isDigit(rune(input[index][0])) {
			// start of the next mapping
			mappings = []rangeMapping{}
		} else {
			// recived range Mapping
			nums := convertStringsToInts(strings.Split(input[index], " "))
			mappings = append(mappings, rangeMapping{nums[0], nums[1], nums[2]})
		}
	}

	seedStructs = mapSeedRanges(seedStructs, mappings)
	// fmt.Println("next seeds", len(seedStructs))
	minimum := seedStructs[0].start
	for _, ss := range seedStructs {
		if ss.start < minimum {
			minimum = ss.start
		}
	}

	return strconv.Itoa(minimum)
}
