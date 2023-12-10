package day10

import (
	"fmt"
	"strconv"
	"strings"
)

var northOpenPipes = "|LJ"
var eastOpenPipes = "LF-"
var southOpenPipes = "|7F"
var westOpenPipes = "-J7"

type Direction int

const (
	North Direction = 1
	East  Direction = 2
	South Direction = 3
	West  Direction = 4
)

type coordinates struct {
	x      int
	y      int
	letter string
	step   int
}

// direction is as follows: 1: north, 2: east, 3: south, 4: west
func checkIfPipeOpenInDirection(pipe string, direction Direction) bool {
	if pipe == "S" {
		return true
	}
	if direction == North {
		return strings.Contains(northOpenPipes, pipe)
	} else if direction == East {
		return strings.Contains(eastOpenPipes, pipe)
	} else if direction == South {
		return strings.Contains(southOpenPipes, pipe)
	} else if direction == West {
		return strings.Contains(westOpenPipes, pipe)
	}
	return false
}

func findStart(input [][]string) coordinates {
	for x, line := range input {
		for y, val := range line {
			if val == "S" {
				return coordinates{x, y, "S", 0}
			}
		}
	}
	return coordinates{-1, -1, "S", 0}
}

func getConnectedCoords(start_coords coordinates, input [][]string) []coordinates {
	results := []coordinates{}
	up_x := start_coords.x - 1
	down_x := start_coords.x + 1
	right_y := start_coords.y + 1
	left_y := start_coords.y - 1
	// check spot above
	if start_coords.x > 0 && checkIfPipeOpenInDirection(input[up_x][start_coords.y], South) &&
		checkIfPipeOpenInDirection(start_coords.letter, North) {
		results = append(results, coordinates{up_x, start_coords.y, input[up_x][start_coords.y], start_coords.step + 1})
	}
	// check spot to right
	if start_coords.y < len(input[0])-1 && checkIfPipeOpenInDirection(input[start_coords.x][right_y], West) &&
		checkIfPipeOpenInDirection(start_coords.letter, East) {
		results = append(results, coordinates{start_coords.x, right_y, input[start_coords.x][right_y], start_coords.step + 1})
	}
	// check spot to bottom
	if start_coords.x < len(input)-1 && checkIfPipeOpenInDirection(input[down_x][start_coords.y], North) &&
		checkIfPipeOpenInDirection(start_coords.letter, South) {
		results = append(results, coordinates{down_x, start_coords.y, input[down_x][start_coords.y], start_coords.step + 1})
	}
	// check spot to left
	if start_coords.y > 0 && checkIfPipeOpenInDirection(input[start_coords.x][left_y], East) &&
		checkIfPipeOpenInDirection(start_coords.letter, West) {
		results = append(results, coordinates{start_coords.x, left_y, input[start_coords.x][left_y], start_coords.step + 1})
	}
	return results
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

	next_coords := []coordinates{findStart(input_split)}
	var next coordinates
	max_val := -1
	// get starting connections to start
	for len(next_coords) > 0 {
		next, next_coords = next_coords[0], next_coords[1:]
		next_coords = append(next_coords, getConnectedCoords(next, input_split)...)
		input_split[next.x][next.y] = strconv.Itoa(next.step)
		if next.step > max_val {
			max_val = next.step
		}
	}

	return strconv.Itoa(max_val)
}
