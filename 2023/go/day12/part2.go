package day12

import (
	"adventUtil"
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

func countArrangementsNoTracking(info_line string, info_index int, group_index int, groups []int) int {
	current_group := groups[group_index]
	// fmt.Println("countArrangements", info_line[info_index:], groups[group_index], "groups_index", group_index, "info index", info_index)
	if info_index+current_group > len(info_line) {
		// fmt.Println("early return")
		return 0
	}
	count := 0
	last_questionmark_count := -1
	for i := info_index; i+current_group <= len(info_line); i++ {
		group_matches := checkSubstringCanBeInMaintaince(info_line, i, i+current_group)
		// fmt.Println("group matches", group_matches, info_line[i:i+current_group], info_index)

		if group_matches {
			if group_index == len(groups)-1 {
				if !strings.Contains(info_line[i+current_group:], "#") {
					// printSolution(info_line, groups, append(tracking, i))
					count++
				}
			} else {
				current_question_mark_count := strings.Count(info_line[i:i+current_group], "?")
				// !(last_questionmark_count == 0 && current_question_mark_count == 0)
				if i+current_group < len(info_line) && info_line[i+current_group] != '#' && !(last_questionmark_count == 0 && current_question_mark_count == 0) {
					count += countArrangementsNoTracking(info_line, i+current_group+1, group_index+1, groups)
					// if current_question_mark_count == 0 {
					// 	// fmt.Println("early return", info_line[i:])
					// 	return count
					// }
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

func Part2Solution(input []string) string {
	fold_factor := 5
	total_sum := 0
	c1 := make(chan int)
	for index, line := range input {

		go func(i int, l string) {
			split_line := strings.Split(l, " ")
			info := split_line[0]
			groups := adventUtil.ConvertStringsToInts(strings.Split(split_line[1], ","))

			bigger_info := multiplyString(info, fold_factor)
			bigger_groups := multiplyIntSlice(groups, fold_factor)
			temp := countArrangementsNoTracking(bigger_info, 0, 0, bigger_groups)
			fmt.Println(info, groups, temp)
			c1 <- temp
		}(index, line)

	}

	i := 0
	for range input {
		next_val := <-c1
		total_sum += next_val
		fmt.Println("(", i, "/", len(input), ")", time.Now())
		i++
	}
	return strconv.Itoa(total_sum)
}

//41.827s
// 31.117

// ( 948 / 1000 ) 2023-12-12 23:28:53.763695 -0500 EST m=+903.233555796
