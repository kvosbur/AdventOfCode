package day14

import (
	"fmt"
	"strconv"
	"time"
)

type Node struct {
	char  rune
	up    *Node
	right *Node
	down  *Node
	left  *Node
}

func makeGraph(input [][]rune) [][]*Node {
	nodes := [][]*Node{}
	// make all nodes
	for x := range input {
		line := []*Node{}
		for y := range input[x] {
			line = append(line, &Node{char: input[x][y], up: nil, right: nil, down: nil, left: nil})
		}
		nodes = append(nodes, line)
	}

	// setup all edges
	for x := range input {
		for y := range input[x] {
			current_node := nodes[x][y]
			// set up
			if x > 0 {
				current_node.up = nodes[x-1][y]
			}
			// set right
			if y < len(input[0])-1 {
				current_node.right = nodes[x][y+1]
			}
			// set down
			if x < len(input)-1 {
				current_node.down = nodes[x+1][y]
			}
			// set left
			if y > 0 {
				current_node.left = nodes[x][y-1]
			}
		}
	}
	return nodes
}

func scoreGraphNorth(nodes [][]*Node) int {
	start := nodes[len(nodes)-1][0]
	row_value := 1
	sum := 0
	for row := start; row != nil; row = row.up {
		for col := row; col != nil; col = col.right {
			if col.char == 'O' {
				sum += row_value
			}
		}
		row_value++
	}
	return sum
}

func shiftGraphRockNorth(node *Node) {
	for n := node; n.up != nil; n = n.up {
		if n.up.char == '.' {
			n.char = '.'
			n.up.char = 'O'
		} else {
			return
		}
	}
}

func shiftGraphAllNorth(start *Node) {
	for next_right := start; next_right != nil; next_right = next_right.right {
		for next_down := next_right.down; next_down != nil; next_down = next_down.down {
			if next_down.char == 'O' {
				shiftGraphRockNorth(next_down)
			}
		}
	}
}

func shiftGraphRockSouth(node *Node) {
	for n := node; n.down != nil; n = n.down {
		if n.down.char == '.' {
			n.char = '.'
			n.down.char = 'O'
		} else {
			return
		}
	}
}

func shiftGraphAllSouth(start *Node) {
	for next_right := start; next_right != nil; next_right = next_right.right {
		for next_up := next_right.up; next_up != nil; next_up = next_up.up {
			if next_up.char == 'O' {
				shiftGraphRockSouth(next_up)
			}
		}
	}
}

func shiftGraphRockWest(node *Node) {
	for n := node; n.left != nil; n = n.left {
		if n.left.char == '.' {
			n.char = '.'
			n.left.char = 'O'
		} else {
			return
		}
	}
}

func shiftGraphAllWest(start *Node) {
	for next_down := start; next_down != nil; next_down = next_down.down {
		for next_right := next_down; next_right != nil; next_right = next_right.right {
			if next_right.char == 'O' {
				shiftGraphRockWest(next_right)
			}
		}
	}
}

func shiftGraphRockEast(node *Node) {
	for n := node; n.right != nil; n = n.right {
		if n.right.char == '.' {
			n.char = '.'
			n.right.char = 'O'
		} else {
			return
		}
	}
}

func shiftGraphAllEast(start *Node) {
	for next_down := start; next_down != nil; next_down = next_down.down {
		for next_left := next_down; next_left != nil; next_left = next_left.left {
			if next_left.char == 'O' {
				shiftGraphRockEast(next_left)
			}
		}
	}
}

func printGraph(nodes [][]*Node) {
	for x := range nodes {
		for _, y := range nodes[x] {
			fmt.Print(string(y.char))
		}
		fmt.Println("")
	}
}

func Part2BruteGraphSolution(input []string) string {
	input_split := [][]rune{}
	cycle_count := 1000000000
	for _, val := range input {
		temp := []rune{}
		for _, char := range val {
			temp = append(temp, char)
		}
		input_split = append(input_split, temp)
	}
	fmt.Println("++++++++++++++")

	nodes := makeGraph(input_split)
	north_start := nodes[0][0]
	west_start := nodes[0][0]
	south_start := nodes[len(nodes)-1][0]
	east_start := nodes[0][len(nodes[0])-1]

	for i := 0; i < cycle_count; i++ {
		shiftGraphAllNorth(north_start)
		shiftGraphAllWest(west_start)
		shiftGraphAllSouth(south_start)
		shiftGraphAllEast(east_start)
		if i%100000 == 0 {
			fmt.Println(float64(i)/float64(cycle_count), time.Now())
		}
	}
	printGraph(nodes)

	value := scoreGraphNorth(nodes)
	return strconv.Itoa(value)
}

// existing brute force
// 0.01 2023-12-14 22:52:16.891055 -0500 EST m=+4.631061486
// 0.1 2023-12-14 22:52:59.673312 -0500 EST m=+47.413830315
//

// first graph shot
// 0.01 2023-12-14 22:50:04.24575 -0500 EST m=+4.548163865
// 0.1 2023-12-14 22:49:04.26939 -0500 EST m=+46.850838779
// 0.1 2023-12-14 22:51:15.202099 -0500 EST m=+45.261446818

// WITH FULL INPUT
// exisitng
// 0.0001 2023-12-14 22:59:19.928284 -0500 EST m=+17.033106495
// 0.0002 2023-12-14 22:59:38.069107 -0500 EST m=+35.174147103

// graphing
//
//
