package day7

import (
	"strconv"
	"strings"
)

func Part2Solution(input []string) string {
	splitInput := strings.Split(input[0], ",")
	ints := make([]int, len(splitInput))

	for i, s := range splitInput {
		ints[i], _ = strconv.Atoi(s)
	}

	c := makeComputer(ints)

	best := 0
	Perm([]int{9, 8, 7, 6, 5}, func(a []int) {
		input := 0
		acs := []*computer{c.copy(), c.copy(), c.copy(), c.copy(), c.copy()}

		for i, val := range a {
			ac := acs[i]
			ac.input = []int{val}
		}

		for {
			halted := false

			for i := range a {
				ac := acs[i]
				ac.input = append(ac.input, input)
				ac.runComputerTillEndOrOutput()
				if !ac.executing {
					halted = true
					break
				}
				output := ac.outputs[0]
				ac.outputs = []int{}
				input = output
			}
			if halted {
				break
			}
		}

		if input > best {
			best = input
		}
	})

	return strconv.Itoa(best)
}

// answer is 57660948
