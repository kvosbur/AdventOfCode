package day1

import "testing"

func TestPart2BasicOneLine(t *testing.T) {
	input := []string{
		"12",
	}

	output := Day1Part2Solution(input)
	expected := "2"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestPart2LargeValue(t *testing.T) {
	input := []string{
		"1969",
	}

	output := Day1Part2Solution(input)
	expected := "966"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}
