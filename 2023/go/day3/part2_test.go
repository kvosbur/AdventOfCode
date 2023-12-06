package day3_test

import (
	"day3"
	"testing"
)

func TestBasic2(t *testing.T) {
	inputs := []testCase{
		{
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			solution: "467835"},
	}

	for _, tc := range inputs {
		output := day3.Part2Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
