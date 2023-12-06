package day4

import (
	"strconv"
	"strings"
)

func Part1Solution(input []string) string {
	score := 0
	for _, card_line := range input {
		split_card_line := strings.Split(card_line, ":")
		numbers_split := strings.Split(split_card_line[1], "|")
		winning_numbers := strings.Split(strings.Trim(numbers_split[0], " "), " ")
		my_numbers := strings.Split(strings.Trim(numbers_split[1], " "), " ")
		line_score := 0
		for _, my_num := range my_numbers {
			if my_num == "" {
				continue
			}
			for _, win_num := range winning_numbers {
				if win_num == "" {
					continue
				}
				if my_num == win_num {
					if line_score == 0 {
						line_score = 1
					} else {
						line_score = line_score * 2
					}
				}
			}

		}
		score += line_score
	}

	return strconv.Itoa(score)
}
