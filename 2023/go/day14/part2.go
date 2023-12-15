package day14

import (
	"fmt"
	"strconv"
	"strings"
)

func shiftRockSouth(input [][]string, start_x int, y int) {
	for x := start_x + 1; x < len(input); x++ {
		if input[x][y] == "." {
			input[x][y] = "O"
			input[x-1][y] = "."
		} else {
			return
		}
	}
}

func shiftAllSouth(input [][]string) {
	for x := len(input) - 2; x >= 0; x-- {
		for y := range input[x] {
			if input[x][y] == "O" && input[x+1][y] == "." {
				shiftRockSouth(input, x, y)
			}
		}
	}
}

func shiftRockWest(input [][]string, x int, start_y int) {
	for y := start_y - 1; y >= 0; y-- {
		if input[x][y] == "." {
			input[x][y] = "O"
			input[x][y+1] = "."
		} else {
			return
		}
	}
}

func shiftAllWest(input [][]string) {
	for y := 1; y < len(input[0]); y++ {
		for x := range input {
			if input[x][y] == "O" && input[x][y-1] == "." {
				shiftRockWest(input, x, y)
			}
		}
	}
}

func shiftRockEast(input [][]string, x int, start_y int) {
	for y := start_y + 1; y < len(input[0]); y++ {
		if input[x][y] == "." {
			input[x][y] = "O"
			input[x][y-1] = "."
		} else {
			return
		}
	}
}

func shiftAllEast(input [][]string) {
	for y := len(input[0]) - 2; y >= 0; y-- {
		for x := range input {
			if input[x][y] == "O" && input[x][y+1] == "." {
				shiftRockEast(input, x, y)
			}
		}
	}
}

func Part2Solution(input []string) string {
	input_split := [][]string{}
	cycle_count := 10000
	for _, val := range input {
		input_split = append(input_split, strings.Split(val, ""))
	}
	fmt.Println("++++++++++++++")
	seen_values := map[int][]int{}

	// day11.PrintInput(input_split)
	for i := 0; i < cycle_count; i++ {
		shiftAllNorth(input_split)
		shiftAllWest(input_split)
		shiftAllSouth(input_split)
		shiftAllEast(input_split)
		value := scoreNorth(input_split)
		map_res, found := seen_values[value]
		if !found {
			// fmt.Println("new value:", i, value)
			seen_values[value] = []int{i}
		} else {
			map_res = append(map_res, i)
			seen_values[value] = map_res
		}
	}

	fmt.Println("---------------")
	for key, val := range seen_values {
		if len(val) < 3 {
			continue
		}
		fmt.Println("-------------------")
		fmt.Println(key)
		new_arr := []int{val[0]}
		for x := 1; x < len(val); x++ {
			new_arr = append(new_arr, val[x]-val[x-1])
		}
		fmt.Println(new_arr)
	}
	// just figured out solution by analyzing the patterns above.

	value := scoreNorth(input_split)
	return strconv.Itoa(value)
}
