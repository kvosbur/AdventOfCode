package day10_test

import (
	"day10"
	"testing"
)

type testCase struct {
	input    []string
	solution string
}

func TestBasic1(t *testing.T) {
	inputs := []testCase{
		{
			input: []string{
				"7-F7-",
				".FJ|7",
				"SJLL7",
				"|F--J",
				"LJ.LJ",
			},
			solution: "8"},
		{
			input: []string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			},
			solution: "4"},
	}

	for _, tc := range inputs {
		output := day10.Part1Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
