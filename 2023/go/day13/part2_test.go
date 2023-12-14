package day13_test

import (
	"day13"
	"strings"
	"testing"
)

func TestBasic2(t *testing.T) {
	inputs := []testCase{
		{
			input: strings.Split(`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`, "\n"),
			solution: "400"},
	}
	for _, tc := range inputs {
		output := day13.Part2Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
