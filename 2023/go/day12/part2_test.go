package day12_test

import (
	"day12"
	"strings"
	"testing"
)

func TestBasic2(t *testing.T) {
	inputs := []testCase{
		{
			input: strings.Split(`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`, "\n"),
			solution: "525152"},
		{input: []string{"???.### 1,1,3"}, solution: "1"},
		{input: []string{".??..??...?##. 1,1,3"}, solution: "16384"},
		{input: []string{"?#?#?#?#?#?#?#? 1,3,1,6"}, solution: "1"},
		{input: []string{"????.#...#... 4,1,1"}, solution: "16"},
		{input: []string{"????.######..#####. 1,6,5"}, solution: "2500"},
		{input: []string{"?###???????? 3,2,1"}, solution: "506250"},
	}
	// 4 * (5 + 4 + 3)
	for _, tc := range inputs {
		output := day12.Part2Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
