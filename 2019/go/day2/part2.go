package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makeCopy(ints []int) []int {
	newInts := make([]int, len(ints))
	copy(newInts, ints)
	return newInts
}

func Part2Solution(input []string) string {
	splitInput := strings.Split(input[0], ",")
	ints := make([]int, len(splitInput))

	for i, s := range splitInput {
		ints[i], _ = strconv.Atoi(s)
	}

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			intsCopy := makeCopy(ints)
			intsCopy[1] = noun
			intsCopy[2] = verb

			index := 0
			finish := false
			for {
				opcode := intsCopy[index]
				switch opcode {
				case 99:
					finish = true
				case 1:
					doAdd(intsCopy, index)
					index += 4
				case 2:
					doMultiply(intsCopy, index)
					index += 4
				default:
					fmt.Println("unknown opcode", opcode)
					os.Exit(1)
				}

				if finish {
					break
				}
			}

			if intsCopy[0] == 19690720 {
				fmt.Println("noun:", noun, "verb:", verb)
				return strconv.Itoa(100*noun + verb)
			}
		}
	}

	return "NA"
}

// noun: 64 verb: 29
// solution is 6429
