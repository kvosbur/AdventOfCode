package day6

import "testing"

func TestBasic(t *testing.T) {
	input := []string{
		"A)B",
	}

	output := Part1Solution(input)

	expected := "1"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestSumation(t *testing.T) {
	input := []string{
		"A)B",
		"B)C",
	}

	output := Part1Solution(input)

	expected := "3"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestGivenExample(t *testing.T) {
	input := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}

	output := Part1Solution(input)

	expected := "42"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}
