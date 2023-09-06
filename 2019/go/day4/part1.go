package day4

import (
	"strconv"
	"strings"
)

func Part1Solution(input []string) string {
	numbers := strings.Split(input[0], "-")
	low, _ := strconv.Atoi(numbers[0])
	high, _ := strconv.Atoi(numbers[1])

	countMatch := 0
	for i := low; i <= high; i++ {
		// need to verify increasing / adjacent pair
		adjacentPairFound := false
		prevDigit := i % 10
		newVal := i / 10
		for {
			nextDigit := newVal % 10
			if nextDigit > prevDigit {
				break
			}
			if nextDigit == prevDigit {
				adjacentPairFound = true
			}
			newVal = newVal / 10
			prevDigit = nextDigit
			if newVal == 0 {
				if adjacentPairFound {
					countMatch += 1
				}
				break
			}
		}
	}
	return strconv.Itoa(countMatch)
}

// correct answer is 1178
