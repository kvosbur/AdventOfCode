package day15_test

import (
	"day15"
	"testing"
)

type testCase struct {
	input    []string
	solution string
}

func TestBasic1(t *testing.T) {
	inputs := []testCase{
		{
			input:    []string{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"},
			solution: "1320"},
	}
	for _, tc := range inputs {
		output := day15.Part1Solution(tc.input)
		if output != tc.solution {
			t.Error("Output does not match expected, received:", output, "expected:", tc.solution)
		}
	}

}
