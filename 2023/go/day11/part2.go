package day11

import (
	"adventUtil"
	"slices"
	"strconv"
	"strings"
)

func trackEmptyVertical(input [][]string) []int {
	empty_vert := []int{}

	for y := 0; y < len(input[0]); y++ {
		is_empty := true
		for x := range input {
			if input[x][y] != "." {
				is_empty = false
			}
		}
		if is_empty {
			empty_vert = append(empty_vert, y)
		}
	}

	return empty_vert
}

func trackEmptyHorizontals(input [][]string) []int {
	empty_horiz := []int{}

	for x := 0; x < len(input); x++ {
		is_empty := true
		for _, val := range input[x] {
			if val != "." {
				is_empty = false
			}
		}
		if is_empty {
			empty_horiz = append(empty_horiz, x)
		}
	}

	return empty_horiz
}

func Part2Solution(input []string) string {
	multiply_factor := 1000000
	input_split := [][]string{}
	for _, val := range input {
		input_split = append(input_split, strings.Split(val, ""))
	}

	coords := []adventUtil.Coordinate{}
	for x := range input_split {
		for y := range input_split[x] {
			if input_split[x][y] == "#" {
				coords = append(coords, adventUtil.Coordinate{X: x, Y: y})
			}
		}
	}

	empty_horiz := trackEmptyHorizontals(input_split)
	empty_vert := trackEmptyVertical(input_split)

	distance_sum := 0
	for i, uni := range coords {
		for _, uni2 := range coords[i+1:] {
			distance := 0
			start_x := adventUtil.MinInt(uni.X, uni2.X)
			end_x := adventUtil.MaxInt(uni.X, uni2.X)
			for x := start_x + 1; x <= end_x; x++ {
				if slices.Contains(empty_horiz, x) {
					distance += multiply_factor
				} else {
					distance++
				}
			}

			start_y := adventUtil.MinInt(uni.Y, uni2.Y)
			end_y := adventUtil.MaxInt(uni.Y, uni2.Y)
			for y := start_y + 1; y <= end_y; y++ {
				if slices.Contains(empty_vert, y) {
					distance += multiply_factor
				} else {
					distance++
				}
			}
			distance_sum += distance
		}
	}

	return strconv.Itoa(distance_sum)
}
