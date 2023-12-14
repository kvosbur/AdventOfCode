package day13

import (
	"strconv"
	"strings"
)

func columnsAreEqual(pattern [][]string, first_col int, second_col int) bool {
	for _, line := range pattern {
		if line[first_col] != line[second_col] {
			return false
		}
	}
	return true
}

func findVerticalSplit(pattern [][]string, ignore_start int) (bool, int) {
	for start := 0; start < len(pattern[0])-1; start++ {
		if start == ignore_start {
			continue
		}
		left_column := start
		right_column := start + 1
		columns_all_match := true
		for left_column >= 0 && right_column < len(pattern[0]) {
			if !columnsAreEqual(pattern, left_column, right_column) {
				columns_all_match = false
			}

			left_column--
			right_column++
		}

		if columns_all_match {
			return true, start + 1
		}
	}
	return false, -5
}

func rowsAreEqual(pattern [][]string, first_row int, second_row int) bool {
	for y := range pattern[0] {
		if pattern[first_row][y] != pattern[second_row][y] {
			return false
		}
	}
	return true
}

func findHorizontalSplit(pattern [][]string, ignore_start int) (bool, int) {
	for start := 0; start < len(pattern)-1; start++ {
		if start == ignore_start {
			continue
		}
		high_row := start
		low_row := start + 1
		rows_all_match := true
		for high_row >= 0 && low_row < len(pattern) {
			if !rowsAreEqual(pattern, high_row, low_row) {
				rows_all_match = false
			}

			high_row--
			low_row++
		}

		if rows_all_match {
			return true, start + 1
		}
	}
	return false, -5
}

func evaluatePattern(pattern [][]string) int {
	sum := 0
	has_vertical, vert_value := findVerticalSplit(pattern, -1)
	if has_vertical {
		sum += vert_value
	}
	has_row, horiz_value := findHorizontalSplit(pattern, -1)
	if has_row {
		sum += (100 * horiz_value)
	}
	return sum
}

func Part1Solution(input []string) string {
	pattern := [][]string{}

	sum := 0
	for _, line := range input {
		if line == "" {
			sum += evaluatePattern(pattern)
			pattern = [][]string{}
		} else {
			pattern = append(pattern, strings.Split(line, ""))
		}
	}
	sum += evaluatePattern(pattern)
	return strconv.Itoa(sum)
}
