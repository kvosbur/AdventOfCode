package day6_test

import (
	"day6"
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
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			solution: "288"},
	}

	for _, tc := range inputs {
		output := day6.Part1Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
