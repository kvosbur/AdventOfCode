package day4

import "testing"

func TestSmallInterval(t *testing.T) {
	input := []string{
		"10-30",
	}

	output := Part1Solution(input)

	expected := "2"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}
