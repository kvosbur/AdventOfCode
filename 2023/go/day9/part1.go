package day9

import (
	"adventUtil"
	"strconv"
	"strings"
)

func isEndOfDiffSlice(diff_slice []int) bool {
	if len(diff_slice) == 1 {
		return true
	}
	for _, val := range diff_slice {
		if val != 0 {
			return false
		}
	}
	return true
}

func getDiffSlices(start_ints []int) [][]int {
	result := [][]int{start_ints}
	result_index := 0
	for !isEndOfDiffSlice(result[result_index]) {
		current_diff_slice := result[result_index]
		next_diff_slice := []int{}
		for i := 1; i < len(current_diff_slice); i++ {
			diff := current_diff_slice[i] - current_diff_slice[i-1]
			next_diff_slice = append(next_diff_slice, diff)
		}
		result = append(result, next_diff_slice)
		result_index++
	}
	return result
}

func extrapolateNextValue(diff_slices [][]int) int {
	current_val := 0
	for i := len(diff_slices) - 1; i >= 0; i-- {
		current_slice := diff_slices[i]
		current_val = current_val + current_slice[len(current_slice)-1]
	}
	return current_val
}

func getNextValueInHistory(start_line []int) int {
	// iteratively make diff arrays
	diff_slices := getDiffSlices(start_line)

	// use diff arrays to get next value
	return extrapolateNextValue(diff_slices)
}

func Part1Solution(input []string) string {
	result := 0
	for _, line := range input {
		int_line := adventUtil.ConvertStringsToInts(strings.Split(line, " "))
		extrapolated := getNextValueInHistory(int_line)
		result += extrapolated
	}
	return strconv.Itoa(result)
}
