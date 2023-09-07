package day6

import "testing"

func TestBasicPart2(t *testing.T) {
	input := []string{
		"COM)A",
		"A)SAN",
		"A)YOU",
	}

	output := Part2Solution(input)

	expected := "2"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestPart2Example(t *testing.T) {
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
		"K)YOU",
		"I)SAN",
	}

	output := Part2Solution(input)

	expected := "4"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}
