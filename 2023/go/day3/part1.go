package day3

import (
	"slices"
	"strconv"
)

var digits = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func isSymbol(character rune) bool {
	return !slices.Contains(digits, character) && character != '.'
}

func hasAdjacentSymbol(input []string, x int, y int) bool {
	// top left
	if x > 0 && y > 0 && isSymbol(rune(input[x-1][y-1])) {
		return true
	}
	// top
	if x > 0 && isSymbol(rune(input[x-1][y])) {
		return true
	}
	// top right
	if x > 0 && y < len(input[0])-1 && isSymbol(rune(input[x-1][y+1])) {
		return true
	}
	// right
	if y < len(input[0])-1 && isSymbol(rune(input[x][y+1])) {
		return true
	}
	// bottom right
	if x < len(input)-1 && y < len(input[0])-1 && isSymbol(rune(input[x+1][y+1])) {
		return true
	}
	// bottom
	if x < len(input)-1 && isSymbol(rune(input[x+1][y])) {
		return true
	}
	// bottom left
	if x < len(input)-1 && y > 0 && isSymbol(rune(input[x+1][y-1])) {
		return true
	}
	// left
	if y > 0 && isSymbol(rune(input[x][y-1])) {
		return true
	}

	return false
}

func Part1Solution(input []string) string {

	sum := 0
	for x, line := range input {
		current_num := ""
		num_should_be_counted := false
		for y, character := range line {
			if slices.Contains(digits, character) {
				current_num += string(character)
				if hasAdjacentSymbol(input, x, y) {
					num_should_be_counted = true
				}
			} else if len(current_num) > 0 {
				if num_should_be_counted {
					num, _ := strconv.Atoi(current_num)
					sum += num
				}
				current_num = ""
				num_should_be_counted = false
			}
		}
		if len(current_num) > 0 {
			if num_should_be_counted {
				num, _ := strconv.Atoi(current_num)
				sum += num
			}
		}
	}

	return strconv.Itoa(sum)
}
