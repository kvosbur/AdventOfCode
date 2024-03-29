package day14_test

import (
	"day14"
	"strings"
	"testing"
)

func TestBasic2(t *testing.T) {
	inputs := []testCase{
		{
			input: strings.Split(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`, "\n"),
			solution: "63"},
	}
	for _, tc := range inputs {
		output := day14.Part2BruteSolution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
