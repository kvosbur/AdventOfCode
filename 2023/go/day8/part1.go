package day8

import (
	"strconv"
	"strings"
)

type node struct {
	left  string
	right string
}

func Part1Solution(input []string) string {
	moves := input[0]
	nodes := map[string]node{}

	for index := 2; index < len(input); index++ {
		line := input[index]
		split_line := strings.Split(line, " = ")
		split_val := strings.Split(split_line[1], ", ")
		nodes[split_line[0]] = node{split_val[0][1:], split_val[1][:len(split_val[1])-1]}
	}

	current_node := "AAA"
	moves_count := 0
	for current_node != "ZZZ" {
		current_move := moves[moves_count%len(moves)]
		n := nodes[current_node]

		if current_move == 'L' {
			current_node = n.left
		} else {
			current_node = n.right
		}

		moves_count = moves_count + 1
	}

	return strconv.Itoa(moves_count)
}
