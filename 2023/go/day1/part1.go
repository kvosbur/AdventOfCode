package day1

import (
	"slices"
	"strconv"
)

func Part1Solution(input []string) string {
	digits := []int{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	sum := 0

	for _, line := range input {
		first_digit := -1
		last_digit := -1
		for _, char := range line {
			intChar := int(char)
			charIsDigit := slices.Contains(digits, intChar)
			if first_digit == -1 && charIsDigit {
				first_digit = intChar - 48
			}
			if charIsDigit {
				last_digit = intChar - 48
			}
		}
		value := first_digit*10 + last_digit
		sum += value
	}

	return strconv.Itoa(sum)
}
