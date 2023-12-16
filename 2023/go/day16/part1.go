package day16

import (
	"fmt"
	"strconv"
)

type node struct {
	char         rune
	energized    bool
	hasGoneNorth bool
	hasGoneEast  bool
	hasGoneSouth bool
	hasGoneWest  bool
}

type Direction int

const (
	North Direction = 0
	East  Direction = 1
	South Direction = 2
	West  Direction = 3
)

func printGraph(graph [][]*node) {
	for x := range graph {
		for y := range graph[x] {
			char := graph[x][y].char
			if graph[x][y].energized {
				char = 'x'
			}
			fmt.Print(string(char))
		}
		fmt.Println("")
	}
}

func getEndDirection(input_direction Direction, char rune) []Direction {
	if char == '/' {
		if input_direction == North {
			return []Direction{East}
		} else if input_direction == East {
			return []Direction{North}
		} else if input_direction == South {
			return []Direction{West}
		} else {
			// input_direction == West
			return []Direction{South}
		}
	} else if char == '\\' {
		if input_direction == North {
			return []Direction{West}
		} else if input_direction == East {
			return []Direction{South}
		} else if input_direction == South {
			return []Direction{East}
		} else {
			// input_direction == West
			return []Direction{North}
		}
	} else if char == '-' {
		if input_direction == West || input_direction == East {
			return []Direction{input_direction}
		} else {
			return []Direction{West, East}
		}

	} else if char == '|' {
		if input_direction == North || input_direction == South {
			return []Direction{input_direction}
		} else {
			return []Direction{North, South}
		}
	} else {
		return []Direction{input_direction}
	}
}

func iterateGraph(graph [][]*node, x int, y int, direction Direction) {
	node_at_position := graph[x][y]
	node_at_position.energized = true
	output_directions := getEndDirection(direction, node_at_position.char)
	for _, dir := range output_directions {
		if dir == North {
			if node_at_position.hasGoneNorth {
				continue
			}
			if x > 0 {
				node_at_position.hasGoneNorth = true
				iterateGraph(graph, x-1, y, dir)
			}
		} else if dir == East {
			if node_at_position.hasGoneEast {
				continue
			}
			if y < len(graph[x])-1 {
				node_at_position.hasGoneEast = true
				iterateGraph(graph, x, y+1, dir)
			}
		} else if dir == South {
			if node_at_position.hasGoneSouth {
				continue
			}
			if x < len(graph)-1 {
				node_at_position.hasGoneSouth = true
				iterateGraph(graph, x+1, y, dir)
			}
		} else if dir == West {
			if node_at_position.hasGoneWest {
				continue
			}
			if y > 0 {
				node_at_position.hasGoneWest = true
				iterateGraph(graph, x, y-1, dir)
			}
		}
	}
}

func countEnergizedGraph(graph [][]*node) int {
	sum := 0
	for x := range graph {
		for y := range graph[x] {
			if graph[x][y].energized {
				sum += 1
			}
		}
	}
	return sum
}

func Part1Solution(input []string) string {
	graph := [][]*node{}

	for _, line := range input {
		graph_line := []*node{}
		for _, r := range line {
			graph_line = append(graph_line, &node{char: r})
		}
		graph = append(graph, graph_line)
	}

	iterateGraph(graph, 0, 0, East)

	printGraph(graph)

	return strconv.Itoa(countEnergizedGraph(graph))
}
