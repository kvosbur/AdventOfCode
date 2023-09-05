package day3

import "testing"

func TestBasicIntersection(t *testing.T) {
	input := []string{
		"R5,U6",
		"U5,R7",
	}

	output := Part1Solution(input)

	expected := "10"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestLowerIntersection(t *testing.T) {
	input := []string{
		"R8,U5,L5,D3",
		"U7,R6,D4,L4",
	}

	output := Part1Solution(input)

	expected := "6"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestComplex1Intersection(t *testing.T) {
	input := []string{
		"R75,D30,R83,U83,L12,D49,R71,U7,L72",
		"U62,R66,U55,R34,D71,R55,D58,R83",
	}

	output := Part1Solution(input)

	expected := "159"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestComplex2Intersection(t *testing.T) {
	input := []string{
		"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
		"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
	}

	output := Part1Solution(input)

	expected := "135"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}
