package day2_test

import (
	"day2"
	"testing"
)

type testCase struct {
	input    []string
	solution string
}

func TestBasic1(t *testing.T) {
	inputs := []testCase{
		{input: []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"}, solution: "1"},
		{input: []string{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"}, solution: "2"},
		{input: []string{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"}, solution: "0"},
		{input: []string{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"}, solution: "0"},
		{input: []string{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}, solution: "5"},
		{
			input: []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			solution: "8"},
	}

	for _, tc := range inputs {
		output := day2.Part1Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
