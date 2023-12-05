package day1_test

import (
	"day1"
	"testing"
)

func TestBasic2(t *testing.T) {
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
		output := day1.Part2Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}
}

func TestSpelledOut1(t *testing.T) {
	inputs := []testCase{
		{input: []string{"two1nine"}, solution: "29"},
		{input: []string{"eightwothree"}, solution: "83"},
		{input: []string{"abcone2threexyz"}, solution: "13"},
		{input: []string{"xtwone3four"}, solution: "24"},
		{input: []string{"4nineeightseven2"}, solution: "42"},
		{input: []string{"zoneight234"}, solution: "14"},
		{input: []string{"7pqrstsixteen"}, solution: "76"},
		{
			input: []string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			solution: "281"},
	}

	for _, tc := range inputs {
		output := day1.Part2Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}
}
