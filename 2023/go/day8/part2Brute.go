package day8

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type travelerBrute struct {
	current_node *nodeRefs
}

type nodeBrute struct {
	left  string
	right string
}

type nodeRefs struct {
	identity string
	is_end   bool
	left     *nodeRefs
	right    *nodeRefs
}

type moveRef struct {
	is_left bool
	next    *moveRef
}

func convertToNodeRefs(node_map map[string]nodeBrute) []*nodeRefs {
	nodes := []*nodeRefs{}
	for name, _ := range node_map {
		nodes = append(nodes, &nodeRefs{identity: name, is_end: name[2] == 'Z'})
	}

	for _, n := range nodes {
		val := node_map[n.identity]
		for _, n2 := range nodes {
			if n2.identity == val.left {
				n.left = n2
			}
			if n2.identity == val.right {
				n.right = n2
			}
		}
	}
	// for _, n := range nodes {
	// 	fmt.Printf("%+v\n", n)
	// }

	return nodes
}

func convertoMoveRefs(moves string) *moveRef {
	moveRefs := []*moveRef{}
	for _, char := range moves {
		moveRefs = append(moveRefs, &moveRef{is_left: char == 'L'})
	}

	for i, ref := range moveRefs {
		ref.next = moveRefs[(i+1)%len(moveRefs)]
	}
	// for _, ref := range moveRefs {
	// 	fmt.Println(ref)
	// }
	return moveRefs[0]
}

func Part2BruteSolution(input []string) string {
	moves := input[0]
	nodes := map[string]nodeBrute{}

	for index := 2; index < len(input); index++ {
		line := input[index]
		split_line := strings.Split(line, " = ")
		split_val := strings.Split(split_line[1], ", ")
		nodes[split_line[0]] = nodeBrute{split_val[0][1:], split_val[1][:len(split_val[1])-1]}
	}

	node_refs := convertToNodeRefs(nodes)

	travelers := []*travelerBrute{}
	for _, n := range node_refs {
		if n.identity[2] == 'A' {
			travelers = append(travelers, &travelerBrute{n})
		}
	}
	fmt.Println(len(travelers))

	current_move := convertoMoveRefs(moves)

	goal := 13289612809129

	moves_count := 0
	all_at_dest := false
	for !all_at_dest {
		all_at_dest = true

		for _, t := range travelers {
			if current_move.is_left {
				t.current_node = t.current_node.left
			} else {
				t.current_node = t.current_node.right
			}

			if !t.current_node.is_end {
				all_at_dest = false
			}
		}

		moves_count = moves_count + 1
		current_move = current_move.next

		if moves_count%100000000 == 0 {
			fmt.Println(float64(moves_count)/float64(goal), time.Now())
			// return ""
		}
	}

	return strconv.Itoa(moves_count)
}

// baseline 1mil moves 100 times   192374234 ns/op

// baseline 100000000   time elapsed 20.363297971

// baseline with refs about a factor of 8 faster than by value
// 0.00010534543181260342 2023-12-09 18:21:12.629314 -0500 EST m=+29.216865531

// improved by pre computing is_end
// 0.00010534543181260342 2023-12-09 18:26:09.14118 -0500 EST m=+26.323272242

// putting moves into a linked list
// 0.00010534543181260342 2023-12-09 18:33:43.771862 -0500 EST m=+17.294891510

// not that big of a difference to precompute left by comparing 'L'
