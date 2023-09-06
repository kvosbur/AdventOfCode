package day4

import (
	"strconv"
	"strings"
)

func Part2Solution(input []string) string {
	numbers := strings.Split(input[0], "-")
	low, _ := strconv.Atoi(numbers[0])
	high, _ := strconv.Atoi(numbers[1])

	countMatch := 0
	for i := low; i <= high; i++ {
		// need to verify increasing / adjacent pair
		prevDigit := i % 10
		newVal := i / 10
		correct := true
		m := make(map[int]int)
		m[prevDigit] = 1
		for {
			nextDigit := newVal % 10
			val, ok := m[nextDigit]
			if !ok {
				val = 0
			}
			m[nextDigit] = val + 1
			if nextDigit > prevDigit {
				correct = false
				break
			}
			newVal = newVal / 10
			prevDigit = nextDigit
			if newVal == 0 {
				break
			}
		}

		hasPair := false
		for _, val := range m {
			if val == 2 {
				hasPair = true
			}
		}
		if hasPair && correct {
			countMatch += 1
		}

	}
	return strconv.Itoa(countMatch)
}

// 763
