package day5

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

func TestAddNewInstruction(t *testing.T) {
	input := []string{
		"101,2,5,0,99,1",
	}

	output := Part1Solution(input)

	expected := "3"
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestEqualPositionModeTrue(t *testing.T) {
	input := []int{
		3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8,
	}
	c := makeComputer(input, 8)
	c.runComputer()
	output := c.outputs[0]
	expected := 1
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestEqualPositionModeFalse(t *testing.T) {
	input := []int{
		3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8,
	}
	c := makeComputer(input, 6)
	c.runComputer()
	output := c.outputs[0]
	expected := 0
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestLessThanPositionModeTrue(t *testing.T) {
	input := []int{
		3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8,
	}
	c := makeComputer(input, 6)
	c.runComputer()
	output := c.outputs[0]
	expected := 1
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestLessThanPositionModeFalse(t *testing.T) {
	input := []int{
		3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8,
	}
	c := makeComputer(input, 8)
	c.runComputer()
	output := c.outputs[0]
	expected := 0
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestEqualImmediateModeTrue(t *testing.T) {
	input := []int{
		3, 3, 1108, -1, 8, 3, 4, 3, 99,
	}
	c := makeComputer(input, 8)
	c.runComputer()
	output := c.outputs[0]
	expected := 1
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestEqualImmediateModeFalse(t *testing.T) {
	input := []int{
		3, 3, 1108, -1, 8, 3, 4, 3, 99,
	}
	c := makeComputer(input, 9)
	c.runComputer()
	output := c.outputs[0]
	expected := 0
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestLessThanImmediateModeFalse(t *testing.T) {
	input := []int{
		3, 3, 1107, -1, 8, 3, 4, 3, 99,
	}
	c := makeComputer(input, 9)
	c.runComputer()
	output := c.outputs[0]
	expected := 0
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestLessThanImmediateModeTrue(t *testing.T) {
	input := []int{
		3, 3, 1107, -1, 8, 3, 4, 3, 99,
	}
	c := makeComputer(input, 5)
	c.runComputer()
	output := c.outputs[0]
	expected := 1
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestJumpPosition1(t *testing.T) {
	input := []int{
		3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9,
	}
	c := makeComputer(input, 58)
	c.runComputer()
	output := c.outputs[0]
	expected := 1
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestJumpPosition1Negate(t *testing.T) {
	input := []int{
		3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9,
	}
	c := makeComputer(input, 0)
	c.runComputer()
	output := c.outputs[0]
	expected := 0
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestJumpImmediate(t *testing.T) {
	input := []int{
		3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1,
	}
	c := makeComputer(input, 58)
	c.runComputer()
	output := c.outputs[0]
	expected := 1
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestJumpImmediateNegate(t *testing.T) {
	input := []int{
		3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1,
	}
	c := makeComputer(input, 0)
	c.runComputer()
	output := c.outputs[0]
	expected := 0
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestLargeExampleLess(t *testing.T) {
	input := []int{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	}
	c := makeComputer(input, 6)
	c.runComputer()
	output := c.outputs[0]
	expected := 999
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestLargeExampleEqual(t *testing.T) {
	input := []int{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	}
	c := makeComputer(input, 8)
	c.runComputer()
	output := c.outputs[0]
	expected := 1000
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}

func TestLargeExampleGreater(t *testing.T) {
	input := []int{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	}
	c := makeComputer(input, 10)
	c.runComputer()
	output := c.outputs[0]
	expected := 1001
	if output != expected {
		t.Error("Output does not match expected, received:", output, "expected:", expected)
	}
}
