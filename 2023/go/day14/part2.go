package day14

import (
	"fmt"
	"strconv"
)

func shiftRockSouth(input [][]rune, start_x int, y int) {
	x := start_x + 1
	input[start_x][y] = '.'
	for x < len(input) {
		if input[x][y] != '.' {
			break
		}
		x++
	}
	input[x-1][y] = 'O'
}

func shiftAllSouth(input [][]rune) {
	for x := len(input) - 2; x >= 0; x-- {
		for y, val := range input[x] {
			if val == 'O' {
				shiftRockSouth(input, x, y)
			}
		}
	}
}

func shiftRockWest(input [][]rune, x int, start_y int) {
	y := start_y - 1
	row := input[x]
	row[start_y] = '.'
	for y >= 0 {
		if row[y] != '.' {
			break
		}
		y--
	}
	row[y+1] = 'O'
}

func shiftAllWest(input [][]rune) {
	for y := 1; y < len(input[0]); y++ {
		for x := range input {
			if input[x][y] == 'O' {
				shiftRockWest(input, x, y)
			}
		}
	}
}

func shiftRockEast(input [][]rune, x int, start_y int) {
	y := start_y + 1
	row := input[x]
	row[start_y] = '.'
	for y < len(row) {
		if row[y] != '.' {
			break
		}
		y++
	}
	row[y-1] = 'O'
}

func shiftAllEast(input [][]rune) {
	for y := len(input[0]) - 2; y >= 0; y-- {
		for x := range input {
			if input[x][y] == 'O' {
				shiftRockEast(input, x, y)
			}
		}
	}
}

func Part2Solution(input []string) string {
	input_split := [][]rune{}
	cycle_count := 1000000000
	for _, val := range input {
		temp := []rune{}
		for _, char := range val {
			temp = append(temp, char)
		}
		input_split = append(input_split, temp)
	}
	fmt.Println("++++++++++++++")

	// day11.PrintInput(input_split)
	for i := 0; i < cycle_count; i++ {
		shiftAllNorth(input_split)
		shiftAllWest(input_split)
		shiftAllSouth(input_split)
		shiftAllEast(input_split)
	}

	value := scoreNorth(input_split)
	return strconv.Itoa(value)
}
