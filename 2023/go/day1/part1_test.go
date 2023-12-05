package day1_test

import (
	"day1"
	"testing"
)

type testCase struct {
	input    []string
	solution string
}

func TestBasic1(t *testing.T) {
	inputs := []testCase{
		{input: []string{"1abc2"}, solution: "12"},
		{input: []string{"pqr3stu8vwx"}, solution: "38"},
		{input: []string{"a1b2c3d4e5f"}, solution: "15"},
		{input: []string{"treb7uchet"}, solution: "77"},
		{
			input: []string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			solution: "142"},
	}

	for _, tc := range inputs {
		output := day1.Part1Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
