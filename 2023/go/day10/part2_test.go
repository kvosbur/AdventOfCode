package day10_test

import (
	"day10"
	"strings"
	"testing"
)

func TestBasic2(t *testing.T) {
	inputs := []testCase{
		{
			input: strings.Split(
				`..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||L.|.
.L--JL--J.
..........`, "\n"),
			solution: "4"},
		{input: strings.Split(
			`.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`, "\n"),
			solution: "8"},
	}

	for _, tc := range inputs {
		output := day10.Part2Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
