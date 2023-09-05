package day3

import "testing"

func TestPart2BasicIntersection(t *testing.T) {
	input := []string{
		"R5,U6",
		"U5,R7",
	}

	output := Part2Solution(input)

	expected := "20"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestPart2LowerIntersection(t *testing.T) {
	input := []string{
		"R8,U5,L5,D3",
		"U7,R6,D4,L4",
	}

	output := Part2Solution(input)

	expected := "30"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestPart2Complex1Intersection(t *testing.T) {
	input := []string{
		"R75,D30,R83,U83,L12,D49,R71,U7,L72",
		"U62,R66,U55,R34,D71,R55,D58,R83",
	}

	output := Part2Solution(input)

	expected := "610"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestPart2Complex2Intersection(t *testing.T) {
	input := []string{
		"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
		"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
	}

	output := Part2Solution(input)

	expected := "410"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}
