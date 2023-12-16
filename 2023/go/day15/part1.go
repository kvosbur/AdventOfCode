package day15

import (
	"strconv"
	"strings"
)

func hash(input string) int {
	value := 0
	for _, char := range input {
		value += int(char)
		value *= 17
		value %= 256
	}
	return value
}

func Part1Solution(input []string) string {
	split_line := strings.Split(input[0], ",")
	sum := 0
	for _, seq := range split_line {
		h := hash(seq)
		sum += h
	}
	return strconv.Itoa(sum)
}
