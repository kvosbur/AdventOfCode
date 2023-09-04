package day1

import "testing"

func TestBasicOneLine(t *testing.T) {
	input := []string{
		"12",
	}

	output := Day1Part1Solution(input)
	expected := "2"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestTruncation(t *testing.T) {
	input := []string{
		"14",
	}

	output := Day1Part1Solution(input)
	expected := "2"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestLargeValue(t *testing.T) {
	input := []string{
		"100756",
	}

	output := Day1Part1Solution(input)
	expected := "33583"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestMultipleValues(t *testing.T) {
	input := []string{
		"12",
		"14",
	}

	output := Day1Part1Solution(input)
	expected := "4"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}
