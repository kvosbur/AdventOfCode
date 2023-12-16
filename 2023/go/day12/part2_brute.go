package day12

import (
	"adventUtil"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func multiplyStringBrute(input string, factor int) string {
	temp := []string{}
	for i := 0; i < factor; i++ {
		temp = append(temp, input)
	}
	return strings.Join(temp, "?")
}

func multiplyIntSliceBrute(input []int, factor int) []int {
	temp := []int{}
	for i := 0; i < factor; i++ {
		temp = append(temp, input...)
	}
	return temp
}

func checkSubstringCanBeInMaintainceBrute(info_line string, start int, end int) bool {
	for i := start; i < end; i++ {
		if info_line[i] == '.' {
			return false
		}
	}
	return true
}

func countArrangementsNoMem(info_line string, info_index int, group_index int, groups []int) int {
	current_group := groups[group_index]
	if info_index+current_group > len(info_line) {
		return 0
	}
	count := 0
	last_questionmark_count := -1
	for i := info_index; i+current_group <= len(info_line); i++ {
		group_matches := checkSubstringCanBeInMaintainceBrute(info_line, i, i+current_group)

		if group_matches {
			if group_index == len(groups)-1 {
				if !strings.Contains(info_line[i+current_group:], "#") {
					count++
				}
			} else {
				current_question_mark_count := strings.Count(info_line[i:i+current_group], "?")
				if i+current_group < len(info_line) && info_line[i+current_group] != '#' && !(last_questionmark_count == 0 && current_question_mark_count == 0) {
					arrangement_count := countArrangementsNoMem(info_line, i+current_group+1, group_index+1, groups)
					count += arrangement_count

					last_questionmark_count = current_question_mark_count
				}
			}
		}
		if info_line[i] == '#' {
			return count
		}
	}
	return count
}

func Part2BruteSolution(input []string) string {
	fold_factor := 5
	total_sum := 0
	for i, line := range input {
		split_line := strings.Split(line, " ")
		info := split_line[0]
		groups := adventUtil.ConvertStringsToInts(strings.Split(split_line[1], ","))

		bigger_info := multiplyStringBrute(info, fold_factor)
		bigger_groups := multiplyIntSliceBrute(groups, fold_factor)
		temp := countArrangementsNoMem(bigger_info, 0, 0, bigger_groups)
		fmt.Println("(", i, "/", 1000, ")", info, groups, temp, time.Now())
		total_sum += temp
	}

	return strconv.Itoa(total_sum)
}

//41.827s
// 31.117

// ( 948 / 1000 ) 2023-12-12 23:28:53.763695 -0500 EST m=+903.233555796
