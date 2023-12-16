package day16_test

import (
	"day16"
	"strings"
	"testing"
)

func TestBasic2(t *testing.T) {
	inputs := []testCase{
		{
			input: strings.Split(`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`, "\n"),
			solution: "51"},
	}
	for _, tc := range inputs {
		output := day16.Part2Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
