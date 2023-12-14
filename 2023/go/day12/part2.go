package day12

import (
	"adventUtil"
	"fmt"
	"strconv"
	"strings"
)

var memorization = map[string]int{}

func multiplyString(input string, factor int) string {
	temp := []string{}
	for i := 0; i < factor; i++ {
		temp = append(temp, input)
	}
	return strings.Join(temp, "?")
}

func multiplyIntSlice(input []int, factor int) []int {
	temp := []int{}
	for i := 0; i < factor; i++ {
		temp = append(temp, input...)
	}
	return temp
}

func getCacheKey(index int, group_index int) string {
	return strconv.Itoa(index) + "," + strconv.Itoa(group_index)
}

func countArrangementsNoTracking(info_line string, info_index int, group_index int, groups []int) int {
	val, cache_hit := memorization[getCacheKey(info_index, group_index)]
	if cache_hit {
		return val
	}
	current_group := groups[group_index]
	if info_index+current_group > len(info_line) {
		return 0
	}
	count := 0
	last_questionmark_count := -1
	for i := info_index; i+current_group <= len(info_line); i++ {
		group_matches := checkSubstringCanBeInMaintaince(info_line, i, i+current_group)

		if group_matches {
			if group_index == len(groups)-1 {
				if !strings.Contains(info_line[i+current_group:], "#") {
					count++
				}
			} else {
				current_question_mark_count := strings.Count(info_line[i:i+current_group], "?")
				if i+current_group < len(info_line) && info_line[i+current_group] != '#' && !(last_questionmark_count == 0 && current_question_mark_count == 0) {
					arrangement_count := countArrangementsNoTracking(info_line, i+current_group+1, group_index+1, groups)
					count += arrangement_count

					last_questionmark_count = current_question_mark_count
				}
			}
		}
		if info_line[i] == '#' {
			memorization[getCacheKey(info_index, group_index)] = count
			return count
		}
	}
	memorization[getCacheKey(info_index, group_index)] = count
	return count
}

func Part2Solution(input []string) string {
	fold_factor := 5
	total_sum := 0
	for i, line := range input {
		split_line := strings.Split(line, " ")
		info := split_line[0]
		groups := adventUtil.ConvertStringsToInts(strings.Split(split_line[1], ","))

		bigger_info := multiplyString(info, fold_factor)
		bigger_groups := multiplyIntSlice(groups, fold_factor)
		temp := countArrangementsNoTracking(bigger_info, 0, 0, bigger_groups)
		for k := range memorization {
			delete(memorization, k)
		}
		fmt.Println("(", i, "/", 1000, ")", info, groups, temp)
		total_sum += temp
	}

	return strconv.Itoa(total_sum)
}

//41.827s
// 31.117

// ( 948 / 1000 ) 2023-12-12 23:28:53.763695 -0500 EST m=+903.233555796
