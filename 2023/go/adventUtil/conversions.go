package adventUtil

import (
	"strconv"
	"strings"
)

func ConvertStringsToInts(input []string) []int {
	dest := []int{}
	for _, str := range input {
		if str == "" {
			continue
		}
		num, _ := strconv.Atoi(strings.Trim(str, " "))
		dest = append(dest, num)
	}
	return dest
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func AbsoluteInt(input int) int {
	if input < 0 {
		return input * -1
	}
	return input
}
