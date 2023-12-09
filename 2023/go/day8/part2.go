package day8

import (
	"strconv"
	"strings"
)

type traveler struct {
	current_node string
	last_z_move  int
	final_z_move int
}

func euclidean_gcd(a int, b int) int {
	if a == 0 || b == 0 {
		return a + b
	}
	larger := a
	smaller := b
	if a < b {
		larger = b
		smaller = a
	}

	return euclidean_gcd(larger%smaller, smaller)
}

func calculate_lcm(a int, b int) int {
	return (a * b) / euclidean_gcd(a, b)
}

func Part2Solution(input []string) string {
	moves := input[0]
	nodes := map[string]node{}

	for index := 2; index < len(input); index++ {
		line := input[index]
		split_line := strings.Split(line, " = ")
		split_val := strings.Split(split_line[1], ", ")
		nodes[split_line[0]] = node{split_val[0][1:], split_val[1][:len(split_val[1])-1]}
	}

	travelers := []traveler{}
	for node, _ := range nodes {
		if node[2] == 'A' {
			travelers = append(travelers, traveler{node, 0, 0})
		}
	}

	moves_count := 0
	all_at_dest := false
	for !all_at_dest {
		all_at_dest = true
		current_move := moves[moves_count%len(moves)]

		for i, t := range travelers {
			n := nodes[t.current_node]
			if current_move == 'L' {
				travelers[i].current_node = n.left
			} else {
				travelers[i].current_node = n.right
			}

			if travelers[i].current_node[2] == 'Z' {
				if travelers[i].last_z_move == 0 {
					travelers[i].last_z_move = moves_count
				} else if travelers[i].final_z_move == 0 {
					travelers[i].final_z_move = moves_count
				}
			}
		}

		for _, t := range travelers {
			if t.final_z_move == 0 {
				all_at_dest = false
			}
		}

		moves_count = moves_count + 1
	}

	our_lcm := travelers[0].final_z_move - travelers[0].last_z_move
	for i := 1; i < len(travelers); i++ {
		next_step := travelers[i].final_z_move - travelers[i].last_z_move
		our_lcm = calculate_lcm(our_lcm, next_step)
	}

	return strconv.Itoa(our_lcm)
}
