package day10

import (
	"adventUtil"
	"fmt"
	"strconv"
	"strings"
)

var blankSpots = "=."

func addVerticalLine(input [][]string, y int) [][]string {
	new_input := [][]string{}
	fmt.Println("y", y)
	// assume that you need to check left for comparison with current y
	for _, line := range input {
		bridging_char := "="
		if checkIfPipeOpenInDirection(line[y-1], East) {
			bridging_char = "-"
		}
		new_line := []string{}
		new_line = append(new_line, line[:y]...)
		new_line = append(new_line, bridging_char)
		new_line = append(new_line, line[y:]...)
		new_input = append(new_input, new_line)
	}
	return new_input
}

func addHorizontalSpaceIfNecessary(input [][]string) [][]string {
	new_input := input
	for _, line := range input {
		for y := 1; y < len(line); y++ {
			if !strings.Contains(blankSpots, line[y]) && !strings.Contains(blankSpots, line[y-1]) && !checkIfPipeOpenInDirection(line[y-1], East) {
				// need to add a vertical line here
				new_input = addVerticalLine(new_input, y)
				fmt.Println("----------------")
				printInput(new_input)
				return addHorizontalSpaceIfNecessary(new_input)
			}
		}
	}
	return new_input
}

func addHorizontalLine(input [][]string, x int) [][]string {
	new_input := [][]string{}
	fmt.Println("x", x)
	// assume that you need to check up for comparison with current x
	new_line := []string{}
	for y := range input[x] {
		bridging_char := "="
		if checkIfPipeOpenInDirection(input[x-1][y], South) {
			bridging_char = "|"
		}
		new_line = append(new_line, bridging_char)
	}
	new_input = append(new_input, input[:x]...)
	new_input = append(new_input, new_line)
	new_input = append(new_input, input[x:]...)
	return new_input
}

func addVerticalSpaceIfNecessary(input [][]string) [][]string {
	new_input := input

	for y := range input[0] {
		for x := 1; x < len(input); x++ {
			if !strings.Contains(blankSpots, input[x][y]) && !strings.Contains(blankSpots, input[x-1][y]) && !checkIfPipeOpenInDirection(input[x-1][y], South) {
				// need to add a horizontal line here
				new_input = addHorizontalLine(new_input, x)
				fmt.Println("----------------")
				printInput(new_input)
				return addVerticalSpaceIfNecessary(new_input)
			}
		}
	}
	return new_input
}

func makeBoolTracker(input [][]string) [][]bool {
	bool_tracker := [][]bool{}
	for range input {
		line := []bool{}
		for range input[0] {
			line = append(line, false)
		}
		bool_tracker = append(bool_tracker, line)
	}
	return bool_tracker
}

// return values: (count of items within, bool whether boundary touched, updated tracker slices)
func fanOut(coord adventUtil.Coordinate, input [][]string, bool_tracker [][]bool) (int, bool, [][]bool, [][]string, []adventUtil.Coordinate) {
	new_input := input
	coordinates := []adventUtil.Coordinate{coord}
	bool_tracker[coord.X][coord.Y] = true

	border_touched := false
	count := 0
	seen_coords := []adventUtil.Coordinate{}
	for len(coordinates) > 0 {
		next_coord := coordinates[0]
		coordinates = coordinates[1:]
		seen_coords = append(seen_coords, next_coord)
		if new_input[next_coord.X][next_coord.Y] == "." {
			count++
		}
		// new_input[next_coord.X][next_coord.Y] = strconv.Itoa(ident % 10)

		border_touched = border_touched || next_coord.X == 0 || next_coord.Y == 0 || next_coord.X == len(input)-1 || next_coord.Y == len(input[0])-1
		// add point above
		if next_coord.X > 0 && strings.Contains(blankSpots, input[next_coord.X-1][next_coord.Y]) && !bool_tracker[next_coord.X-1][next_coord.Y] {
			coordinates = append(coordinates, adventUtil.Coordinate{X: next_coord.X - 1, Y: next_coord.Y})
			bool_tracker[next_coord.X-1][next_coord.Y] = true
		}
		// add point right
		if next_coord.Y < len(input[0])-1 && strings.Contains(blankSpots, input[next_coord.X][next_coord.Y+1]) && !bool_tracker[next_coord.X][next_coord.Y+1] {
			coordinates = append(coordinates, adventUtil.Coordinate{X: next_coord.X, Y: next_coord.Y + 1})
			bool_tracker[next_coord.X][next_coord.Y+1] = true
		}
		// add point below
		if next_coord.X < len(input)-1 && strings.Contains(blankSpots, input[next_coord.X+1][next_coord.Y]) && !bool_tracker[next_coord.X+1][next_coord.Y] {
			coordinates = append(coordinates, adventUtil.Coordinate{X: next_coord.X + 1, Y: next_coord.Y})
			bool_tracker[next_coord.X+1][next_coord.Y] = true
		}
		// add point left
		if next_coord.Y > 0 && strings.Contains(blankSpots, input[next_coord.X][next_coord.Y-1]) && !bool_tracker[next_coord.X][next_coord.Y-1] {
			coordinates = append(coordinates, adventUtil.Coordinate{X: next_coord.X, Y: next_coord.Y - 1})
			bool_tracker[next_coord.X][next_coord.Y-1] = true
		}
	}
	return count, border_touched, bool_tracker, new_input, seen_coords
}

func searchGroups(input [][]string, bool_tracker [][]bool) (int, [][]string) {
	var count int
	var border_touched bool
	var seen_coords []adventUtil.Coordinate
	new_input := input
	total_count := 0
	for x := range input {
		for y := range input[x] {
			if !bool_tracker[x][y] && strings.Contains(blankSpots, input[x][y]) {
				count, border_touched, bool_tracker, new_input, seen_coords = fanOut(adventUtil.Coordinate{X: x, Y: y}, new_input, bool_tracker)
				if !border_touched {
					for _, coord := range seen_coords {
						// fmt.Println(new_input[coord.X][coord.Y])
						if new_input[coord.X][coord.Y] == "." {
							new_input[coord.X][coord.Y] = "+"
						}
					}

					total_count += count
				}
			}
		}
	}
	return total_count, new_input
}

func getConnectedCoordsImproved(start_coords coordinates, input [][]string, bool_tracker [][]bool) []coordinates {
	results := []coordinates{}
	up_x := start_coords.x - 1
	down_x := start_coords.x + 1
	right_y := start_coords.y + 1
	left_y := start_coords.y - 1
	// check spot above
	if start_coords.x > 0 && !bool_tracker[up_x][start_coords.y] && checkIfPipeOpenInDirection(input[up_x][start_coords.y], South) &&
		checkIfPipeOpenInDirection(start_coords.letter, North) {
		results = append(results, coordinates{up_x, start_coords.y, input[up_x][start_coords.y], start_coords.step + 1})
	}
	// check spot to right
	if start_coords.y < len(input[0])-1 && !bool_tracker[start_coords.x][right_y] && checkIfPipeOpenInDirection(input[start_coords.x][right_y], West) &&
		checkIfPipeOpenInDirection(start_coords.letter, East) {
		results = append(results, coordinates{start_coords.x, right_y, input[start_coords.x][right_y], start_coords.step + 1})
	}
	// check spot to bottom
	if start_coords.x < len(input)-1 && !bool_tracker[down_x][start_coords.y] && checkIfPipeOpenInDirection(input[down_x][start_coords.y], North) &&
		checkIfPipeOpenInDirection(start_coords.letter, South) {
		results = append(results, coordinates{down_x, start_coords.y, input[down_x][start_coords.y], start_coords.step + 1})
	}
	// check spot to left
	if start_coords.y > 0 && !bool_tracker[start_coords.x][left_y] && checkIfPipeOpenInDirection(input[start_coords.x][left_y], East) &&
		checkIfPipeOpenInDirection(start_coords.letter, West) {
		results = append(results, coordinates{start_coords.x, left_y, input[start_coords.x][left_y], start_coords.step + 1})
	}
	return results
}

func replaceNonCircleCharacters(input [][]string) [][]string {
	new_input := input
	bool_tracker := makeBoolTracker(new_input)

	next_coords := []coordinates{findStart(new_input)}
	var next coordinates
	max_val := -1
	// figure out circular pipe
	for len(next_coords) > 0 {
		next, next_coords = next_coords[0], next_coords[1:]
		// fmt.Println(next, len(next_coords))
		next_coords = append(next_coords, getConnectedCoordsImproved(next, new_input, bool_tracker)...)
		bool_tracker[next.x][next.y] = true
		if next.step > max_val {
			max_val = next.step
		}
	}
	fmt.Println("bool tracker")

	// iterate and mark all non-circle as "."
	for x := range bool_tracker {
		for y, val := range bool_tracker[x] {
			if !val {
				new_input[x][y] = "."
			}
		}
	}
	return new_input
}

func Part2Solution(input []string) string {
	input_split := [][]string{}
	for _, val := range input {
		input_split = append(input_split, strings.Split(val, ""))
	}

	fmt.Println("before split")
	input_split = replaceNonCircleCharacters(input_split)

	printInput(input_split)

	input_split = addHorizontalSpaceIfNecessary(input_split)
	input_split = addVerticalSpaceIfNecessary(input_split)

	bool_tracker := makeBoolTracker(input_split)

	fmt.Println("====================")
	printInput(input_split)

	solution, input_split := searchGroups(input_split, bool_tracker)
	fmt.Println(solution)
	fmt.Println("====================")
	printInput(input_split)

	return strconv.Itoa(solution)
}
