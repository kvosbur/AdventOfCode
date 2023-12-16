package day16

import "strconv"

func makeGraphCopy(input []string) [][]*node {
	graph := [][]*node{}

	for _, line := range input {
		graph_line := []*node{}
		for _, r := range line {
			graph_line = append(graph_line, &node{char: r})
		}
		graph = append(graph, graph_line)
	}
	return graph
}

func findBestGraphIteration(input []string) int {
	best_so_far := 0
	var graph_copy [][]*node
	var val int
	// iterate rows
	for y := range input[0] {
		// top going down
		graph_copy = makeGraphCopy(input)
		iterateGraph(graph_copy, 0, y, South)
		val = countEnergizedGraph(graph_copy)
		if val > best_so_far {
			best_so_far = val
		}

		// bottom going up
		graph_copy = makeGraphCopy(input)
		iterateGraph(graph_copy, len(input)-1, y, North)
		val = countEnergizedGraph(graph_copy)
		if val > best_so_far {
			best_so_far = val
		}
	}

	for x := range input {
		// left going right
		graph_copy = makeGraphCopy(input)
		iterateGraph(graph_copy, x, 0, East)
		val = countEnergizedGraph(graph_copy)
		if val > best_so_far {
			best_so_far = val
		}

		// right going left
		graph_copy = makeGraphCopy(input)
		iterateGraph(graph_copy, x, len(input[0])-1, West)
		val = countEnergizedGraph(graph_copy)
		if val > best_so_far {
			best_so_far = val
		}
	}
	return best_so_far
}

func Part2Solution(input []string) string {

	return strconv.Itoa(findBestGraphIteration(input))
}
