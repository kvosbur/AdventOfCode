package day12

import (
	"adventUtil"
	"fmt"
	"strconv"
	"strings"
)

func checkSubstringCanBeInMaintaince(info_line string, start int, end int) bool {
	// sub := info_line[start:end]
	// match, _ := regexp.Match("^[#\\?]*$", []byte(sub))
	// return match

	for i := start; i < end; i++ {
		if info_line[i] == '.' {
			return false
		}
	}
	return true
}

func getQuestionLocationsInRange(info_line string, start int, end int) []int {
	result := []int{}
	sub := info_line[start:end]
	for i, char := range sub {
		if char == '?' {
			result = append(result, start+i)
		}
	}
	return result
}

func intSlicesAreEqual(first []int, second []int) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func printSolution(info_line string, groups []int, tracking []int) {
	split_line := strings.Split(info_line, "")
	for i := range groups {
		for j := 0; j < groups[i]; j++ {
			split_line[tracking[i]+j] = "#"
		}
	}
	fmt.Println("Solution", strings.Join(split_line, ""))
}

func countArrangements(info_line string, info_index int, group_index int, groups []int, tracking []int) int {
	current_group := groups[group_index]
	// fmt.Println("countArrangements", info_line[info_index:], groups[group_index], "groups_index", group_index, "info index", info_index)
	if info_index+current_group > len(info_line) {
		// fmt.Println("early return")
		return 0
	}
	count := 0
	last_match_question_locations := []int{-1}
	for i := info_index; i+current_group <= len(info_line); i++ {
		group_matches := checkSubstringCanBeInMaintaince(info_line, i, i+current_group)
		// fmt.Println("group matches", group_matches, info_line[i:i+current_group], info_index)

		if group_matches {
			if group_index == len(groups)-1 {
				if i+current_group == len(info_line) || !strings.Contains(info_line[i+current_group:], "#") {
					// printSolution(info_line, groups, append(tracking, i))
					count++
				}
			} else {
				next_match_question_locations := getQuestionLocationsInRange(info_line, i, i+current_group)
				questions_changed := !intSlicesAreEqual(last_match_question_locations, next_match_question_locations)

				if i+current_group < len(info_line) && info_line[i+current_group] != '#' && questions_changed {
					count += countArrangements(info_line, i+current_group+1, group_index+1, groups, append(tracking, i))
					if strings.Count(info_line[i:i+current_group], "?") == 0 {
						// fmt.Println("early return", info_line[i:])
						return count
					}
					last_match_question_locations = next_match_question_locations
				}
			}
		}
		if info_line[i] == '#' {
			return count
		}
	}
	return count
}

func Part1Solution(input []string) string {
	total_sum := 0
	for _, line := range input {
		split_line := strings.Split(line, " ")
		info := split_line[0]
		groups := adventUtil.ConvertStringsToInts(strings.Split(split_line[1], ","))

		temp := countArrangementsNoTracking(info, 0, 0, groups)
		// temp := countArrangements(info, 0, 0, groups, []int{})
		fmt.Println(info, groups, temp)

		total_sum += temp
	}
	return strconv.Itoa(total_sum)
}
