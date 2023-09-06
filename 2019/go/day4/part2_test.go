package day4

import "testing"

func TestInterestingInterval(t *testing.T) {
	input := []string{
		"110-120",
	}

	output := Part2Solution(input)

	expected := "8"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestEdge1Interval(t *testing.T) {
	input := []string{
		"123444-123445",
	}

	output := Part2Solution(input)

	expected := "1"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestEdge2Interval(t *testing.T) {
	input := []string{
		"112233-112233",
	}

	output := Part2Solution(input)

	expected := "1"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestEdge3Interval(t *testing.T) {
	input := []string{
		"111122-111122",
	}

	output := Part2Solution(input)

	expected := "1"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestEdge4Interval(t *testing.T) {
	input := []string{
		"11112-11112",
	}

	output := Part2Solution(input)

	expected := "0"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}
