package day5

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1Solution(input []string) string {
	splitInput := strings.Split(input[0], ",")
	ints := make([]int, len(splitInput))

	for i, s := range splitInput {
		ints[i], _ = strconv.Atoi(s)
	}

	c := makeComputer(ints, 1)
	c.runComputer()

	fmt.Println(c.outputs)

	return strconv.Itoa(c.outputs[0])
}

// correct answer is 15508323
