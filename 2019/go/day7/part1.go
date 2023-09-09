package day7

import (
	"strconv"
	"strings"
)

// Perm calls f with each permutation of a.
func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func Part1Solution(input []string) string {
	splitInput := strings.Split(input[0], ",")
	ints := make([]int, len(splitInput))

	for i, s := range splitInput {
		ints[i], _ = strconv.Atoi(s)
	}

	c := makeComputer(ints)

	best := 0
	Perm([]int{0, 1, 2, 3, 4}, func(a []int) {
		input := 0
		for _, val := range a {
			ac := c.copy()
			ac.input = []int{val, input}
			ac.runComputer()
			output := ac.outputs[0]
			input = output
		}
		if input > best {
			best = input
		}
		// fmt.Println("output", a, input, best)
	})

	return strconv.Itoa(best)
}

// answer is 118936
