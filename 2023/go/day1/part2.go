package day1

import (
	"slices"
	"strconv"
)

func getDigit(line string, index int) int {
	digits := []int{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	digit_str := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	if slices.Contains(digits, int(line[index])) {
		return int(line[index]) - 48
	}

	for i, spelled_out := range digit_str {
		if index+len(spelled_out) > len(line) {
			continue
		}
		sub := line[index : index+len(spelled_out)]
		if sub == spelled_out {
			return i + 1
		}
	}
	return -1
}

func Part2Solution(input []string) string {

	sum := 0

	for _, line := range input {
		first_digit := -1
		last_digit := -1
		for i, _ := range line {
			digit := getDigit(line, i)
			charIsDigit := digit != -1
			if first_digit == -1 && charIsDigit {
				first_digit = digit
			}
			if charIsDigit {
				last_digit = digit
			}
		}
		value := first_digit*10 + last_digit
		sum += value
	}

	return strconv.Itoa(sum)
}
