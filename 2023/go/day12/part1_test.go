package day12_test

import (
	"day12"
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
			input: strings.Split(`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`, "\n"),
			solution: "21"},
		{input: []string{"???.### 1,1,3"}, solution: "1"},
		{input: []string{".??..??...?##. 1,1,3"}, solution: "4"},
		{input: []string{"?#?#?#?#?#?#?#? 1,3,1,6"}, solution: "1"},
		{input: []string{"????.#...#... 4,1,1"}, solution: "1"},
		{input: []string{"????.######..#####. 1,6,5"}, solution: "4"},
		{input: []string{"?###???????? 3,2,1"}, solution: "10"},
		{input: []string{"??##.?#?.?#?# 4,3,3"}, solution: "1"},
		{input: []string{"?????.??##?????????. 2,6,2"}, solution: "48"},
		{input: []string{"?.???????###.????? 1,2,2,4,3"}, solution: "3"},
		{input: []string{"?.#.??##???### 1,3,4"}, solution: "2"},
	}
	// 4 * (5 + 4 + 3)
	for _, tc := range inputs {
		output := day12.Part1Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
