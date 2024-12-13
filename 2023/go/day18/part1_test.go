package day17_test

import (
	"day17"
	"strings"
	"testing"
)

type testCase struct {
	input    []string
	solution string
}

func TestBasic1(t *testing.T) {
	inputs := []testCase{
		{
			input: strings.Split(`R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`, "\n"),
			solution: "62"},
	}
	for _, tc := range inputs {
		output := day17.Part1Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
