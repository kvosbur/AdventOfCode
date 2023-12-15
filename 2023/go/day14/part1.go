package day14

import (
	"strconv"
)

func shiftRockNorth(input [][]rune, start_x int, y int) {
	for x := start_x - 1; x >= 0; x-- {
		if input[x][y] == '.' {
			input[x][y] = 'O'
			input[x+1][y] = '.'
		} else {
			return
		}
	}
}

func scoreNorth(input [][]rune) int {
	row_value := 1
	sum := 0
	for x := len(input) - 1; x >= 0; x-- {
		for _, val := range input[x] {
			if val == 'O' {
				sum += row_value
			}
		}
		row_value++
	}
	return sum
}

func shiftAllNorth(input [][]rune) {
	for x := 1; x < len(input); x++ {
		for y := range input[x] {
			if input[x][y] == 'O' && input[x-1][y] == '.' {
				shiftRockNorth(input, x, y)
			}
		}
	}
}

func Part1Solution(input []string) string {
	input_split := [][]rune{}
	for _, val := range input {
		temp := []rune{}
		for _, char := range val {
			temp = append(temp, char)
		}
		input_split = append(input_split, temp)
	}

	shiftAllNorth(input_split)

	value := scoreNorth(input_split)
	return strconv.Itoa(value)
}
