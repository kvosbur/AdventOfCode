package day2

import "testing"

func TestHalt(t *testing.T) {
	input := []string{
		"99,1,1,1,0",
	}

	output := Part1Solution(input)

	expected := "99"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestAdd(t *testing.T) {
	input := []string{
		"1,1,1,0,99",
	}

	output := Part1Solution(input)

	expected := "2"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestMultiply(t *testing.T) {
	input := []string{
		"2,1,4,0,99",
	}

	output := Part1Solution(input)

	expected := "99"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestSimple(t *testing.T) {
	input := []string{
		"1,1,1,4,99,5,6,0,99",
	}

	output := Part1Solution(input)

	expected := "30"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}
