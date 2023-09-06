package day5

import (
	"fmt"
	"os"
	"strconv"
)

type computer struct {
	currentInstructionPointer int
	memory                    []int
	executing                 bool
	input                     int
	outputs                   []int
}

func makeComputer(initialMemory []int, input int) computer {
	return computer{
		currentInstructionPointer: 0,
		memory:                    initialMemory,
		executing:                 true,
		input:                     input,
	}
}

func (c *computer) getBaseOpCode() int {
	return c.memory[c.currentInstructionPointer] % 100
}

func (c *computer) doInstruction() {
	opcode := c.getBaseOpCode()
	switch opcode {
	case 99:
		c.executing = false
		c.currentInstructionPointer += 1
	case 1:
		c.doAdd()
	case 2:
		c.doMultiply()
	case 3:
		c.doSaveInput()
	case 4:
		c.doSendOutput()
	case 5:
		c.doJumpIfTrue()
	case 6:
		c.doJumpIfFalse()
	case 7:
		c.doLessThan()
	case 8:
		c.doEquals()
	default:
		fmt.Println("unknown opcode", opcode)
		os.Exit(1)
	}
}

func (c *computer) runComputer() {
	for {
		c.doInstruction()
		if !c.executing {
			break
		}
	}
}

func (c *computer) getMemory(position int) int {
	return c.memory[position]
}

func (c *computer) getParameterValue(p int) int {
	i := c.memory[c.currentInstructionPointer]
	paramValue := c.memory[c.currentInstructionPointer+p]
	iStr := strconv.Itoa(i)

	if len(iStr) <= 2 {
		return c.memory[paramValue]
	}

	iStr = iStr[:len(iStr)-2]
	if (p - 1) >= len(iStr) {
		// if string not long enough then it is 0 which is param mode
		return c.memory[paramValue]
	}
	mode := iStr[len(iStr)-p]
	if string(mode) == "1" {
		return paramValue
	} else {
		return c.memory[paramValue]
	}
}

func (c *computer) doAdd() {
	first := c.getParameterValue(1)
	second := c.getParameterValue(2)
	storeLocation := c.getMemory(c.currentInstructionPointer + 3)
	c.memory[storeLocation] = first + second
	c.currentInstructionPointer += 4
}

func (c *computer) doMultiply() {
	first := c.getParameterValue(1)
	second := c.getParameterValue(2)
	storeLocation := c.getMemory(c.currentInstructionPointer + 3)
	c.memory[storeLocation] = first * second
	c.currentInstructionPointer += 4
}

func (c *computer) doSaveInput() {
	storeLocation := c.getMemory(c.currentInstructionPointer + 1)
	c.memory[storeLocation] = c.input
	c.currentInstructionPointer += 2
}

func (c *computer) doSendOutput() {
	retrievedValue := c.getParameterValue(1)
	c.outputs = append(c.outputs, retrievedValue)
	fmt.Println("Output:", retrievedValue)
	c.currentInstructionPointer += 2
}

func (c *computer) doJumpIfTrue() {
	conditional := c.getParameterValue(1)
	address := c.getParameterValue(2)
	if conditional != 0 {
		c.currentInstructionPointer = address
	} else {
		c.currentInstructionPointer += 3
	}
}

func (c *computer) doJumpIfFalse() {
	conditional := c.getParameterValue(1)
	address := c.getParameterValue(2)
	if conditional == 0 {
		c.currentInstructionPointer = address
	} else {
		c.currentInstructionPointer += 3
	}
}

func (c *computer) doLessThan() {
	first := c.getParameterValue(1)
	second := c.getParameterValue(2)
	address := c.getMemory(c.currentInstructionPointer + 3)
	if first < second {
		c.memory[address] = 1
	} else {
		c.memory[address] = 0
	}
	c.currentInstructionPointer += 4
}

func (c *computer) doEquals() {
	first := c.getParameterValue(1)
	second := c.getParameterValue(2)
	address := c.getMemory(c.currentInstructionPointer + 3)
	if first == second {
		c.memory[address] = 1
	} else {
		c.memory[address] = 0
	}
	c.currentInstructionPointer += 4
}
