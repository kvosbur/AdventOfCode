package day14

import (
	"fmt"
	"strconv"
)

func shiftRockNorth(input [][]rune, start_x int, y int) {
	x := start_x - 1
	input[start_x][y] = '.'
	for x >= 0 {
		if input[x][y] != '.' {
			break
		}
		x--
	}
	input[x+1][y] = 'O'
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
		for y, val := range input[x] {
			if val == 'O' {
				shiftRockNorth(input, x, y)
			}
		}
	}
}

func printInput(split_input [][]rune) {
	for _, row := range split_input {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println("")
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

	printInput(input_split)

	value := scoreNorth(input_split)
	return strconv.Itoa(value)
}
