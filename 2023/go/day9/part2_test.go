package day9_test

import (
	"day9"
	"testing"
)

func TestBasic2(t *testing.T) {
	inputs := []testCase{
		{
			input: []string{
				"0 3 6 9 12 15",
				"1 3 6 10 15 21",
				"10 13 16 21 30 45",
			},
			solution: "2"},
	}

	for _, tc := range inputs {
		output := day9.Part2Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
