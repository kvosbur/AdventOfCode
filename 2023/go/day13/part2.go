package day13

import (
	"fmt"
	"strconv"
	"strings"
)

func printPattern(pattern [][]string) {
	for _, line := range pattern {
		fmt.Println(strings.Join(line, ""))
	}
	fmt.Println("")
}

func newEvaluatePattern(pattern [][]string, ignore_horiz int, ignore_vert int) int {
	sum := 0
	has_vertical, vert_value := findVerticalSplit(pattern, ignore_vert)
	if has_vertical {
		sum += vert_value
	}
	has_row, horiz_value := findHorizontalSplit(pattern, ignore_horiz)
	if has_row {
		sum += (100 * horiz_value)
	}
	return sum
}

func copyPattern(pattern [][]string) [][]string {
	new_pattern := [][]string{}
	for x := range pattern {
		new_pattern = append(new_pattern, []string{})
		new_pattern[x] = append(new_pattern[x], pattern[x]...)
	}
	return new_pattern
}

func findDifferentWorkingPattern(pattern [][]string) int {
	initial_sum := 0
	has_vertical, vert_value := findVerticalSplit(pattern, -1)
	if has_vertical {
		initial_sum += vert_value
	}
	has_row, horiz_value := findHorizontalSplit(pattern, -1)
	if has_row {
		initial_sum += (100 * horiz_value)
	}
	for x := range pattern {
		for y := range pattern[x] {
			new_pattern := copyPattern(pattern)
			if new_pattern[x][y] == "." {
				new_pattern[x][y] = "#"
			} else {
				new_pattern[x][y] = "."
			}
			next_val := newEvaluatePattern(new_pattern, horiz_value-1, vert_value-1)
			if next_val != initial_sum && next_val != 0 {
				return next_val
			}
		}
	}

	return -1
}

func Part2Solution(input []string) string {
	pattern := [][]string{}

	sum := 0
	for _, line := range input {
		if line == "" {
			sum += findDifferentWorkingPattern(pattern)
			pattern = [][]string{}
		} else {
			pattern = append(pattern, strings.Split(line, ""))
		}
	}
	sum += findDifferentWorkingPattern(pattern)
	return strconv.Itoa(sum)
}
