package day9

import (
	"adventUtil"
	"strconv"
	"strings"
)

func extrapolatePrevValue(diff_slices [][]int) int {
	current_val := 0
	for i := len(diff_slices) - 1; i >= 0; i-- {
		current_slice := diff_slices[i]
		current_val = current_slice[0] - current_val
	}
	return current_val
}

func getPrevValueInHistory(start_line []int) int {
	// iteratively make diff arrays
	diff_slices := getDiffSlices(start_line)

	// use diff arrays to get next value
	return extrapolatePrevValue(diff_slices)
}

func Part2Solution(input []string) string {
	result := 0
	for _, line := range input {
		int_line := adventUtil.ConvertStringsToInts(strings.Split(line, " "))
		extrapolated := getPrevValueInHistory(int_line)
		result += extrapolated
	}
	return strconv.Itoa(result)
}
