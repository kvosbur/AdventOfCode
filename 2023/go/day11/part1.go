package day11

import (
	"adventUtil"
	"fmt"
	"strconv"
	"strings"
)

func addVerticalLine(input [][]string, y int) [][]string {
	new_input := [][]string{}
	for _, line := range input {
		new_line := []string{}
		new_line = append(new_line, line[:y]...)
		new_line = append(new_line, ".")
		new_line = append(new_line, line[y:]...)
		new_input = append(new_input, new_line)
	}
	return new_input
}

func addVerticalLineIfNecessary(input [][]string) [][]string {
	new_input := input

	for y := 0; y < len(new_input[0]); y++ {
		is_empty := true
		for x := range new_input {
			if new_input[x][y] != "." {
				is_empty = false
			}
		}
		if is_empty {
			new_input = addVerticalLine(new_input, y)
			y++
		}
	}

	return new_input
}

func addHorizontalLine(input [][]string, x int) [][]string {
	new_input := [][]string{}
	new_line := []string{}
	for range input[x] {
		new_line = append(new_line, ".")
	}
	new_input = append(new_input, input[:x]...)
	new_input = append(new_input, new_line)
	new_input = append(new_input, input[x:]...)
	return new_input
}

func addHorizontalLineIfNecessary(input [][]string) [][]string {
	new_input := input

	for x := 0; x < len(new_input); x++ {
		is_empty := true
		for _, val := range new_input[x] {
			if val != "." {
				is_empty = false
			}
		}
		if is_empty {
			new_input = addHorizontalLine(new_input, x)
			x++
		}
	}

	return new_input
}

func printInput(split_input [][]string) {
	for _, val := range split_input {
		fmt.Println(val)
	}
}

func Part1Solution(input []string) string {
	input_split := [][]string{}
	for _, val := range input {
		input_split = append(input_split, strings.Split(val, ""))
	}

	input_split = addVerticalLineIfNecessary(input_split)
	input_split = addHorizontalLineIfNecessary(input_split)

	coords := []adventUtil.Coordinate{}
	for x := range input_split {
		for y := range input_split[x] {
			if input_split[x][y] == "#" {
				coords = append(coords, adventUtil.Coordinate{X: x, Y: y})
			}
		}
	}

	distance_sum := 0
	for i, uni := range coords {
		for _, uni2 := range coords[i+1:] {
			distance_sum += adventUtil.AbsoluteInt(uni.X-uni2.X) + adventUtil.AbsoluteInt(uni.Y-uni2.Y)
		}
	}

	return strconv.Itoa(distance_sum)
}
