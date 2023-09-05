package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func doAdd(ints []int, index int) {
	firstLocation := ints[index+1]
	secondLocation := ints[index+2]
	storeLocation := ints[index+3]
	ints[storeLocation] = ints[firstLocation] + ints[secondLocation]
}

func doMultiply(ints []int, index int) {
	firstLocation := ints[index+1]
	secondLocation := ints[index+2]
	storeLocation := ints[index+3]
	ints[storeLocation] = ints[firstLocation] * ints[secondLocation]
}

func Part1Solution(input []string) string {
	splitInput := strings.Split(input[0], ",")
	ints := make([]int, len(splitInput))

	for i, s := range splitInput {
		ints[i], _ = strconv.Atoi(s)
	}

	index := 0
	finish := false
	for {
		opcode := ints[index]
		switch opcode {
		case 99:
			finish = true
		case 1:
			doAdd(ints, index)
			index += 4
		case 2:
			doMultiply(ints, index)
			index += 4
		default:
			fmt.Println("unknown opcode", opcode)
			os.Exit(1)
		}

		if finish {
			break
		}
	}
	return strconv.Itoa(ints[0])
}

// solution is 3716293
