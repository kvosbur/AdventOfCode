package day4

import (
	"strconv"
	"strings"
)

func Part2Solution(input []string) string {
	card_counts := map[int]int{}
	for i := range input {
		card_counts[i] = 1
	}

	for i, card_line := range input {
		weight := card_counts[i]
		split_card_line := strings.Split(card_line, ":")
		numbers_split := strings.Split(split_card_line[1], "|")
		winning_numbers := strings.Split(strings.Trim(numbers_split[0], " "), " ")
		my_numbers := strings.Split(strings.Trim(numbers_split[1], " "), " ")
		win_count := 0
		for _, my_num := range my_numbers {
			if my_num == "" {
				continue
			}
			for _, win_num := range winning_numbers {
				if win_num == "" {
					continue
				}
				if my_num == win_num {
					win_count += 1
				}
			}
		}
		for j := i + 1; j <= i+win_count; j++ {
			card_counts[j] += weight
		}
	}

	sum := 0
	for _, val := range card_counts {
		sum += val
	}

	return strconv.Itoa(sum)
}
