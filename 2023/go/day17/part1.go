package day17

import (
	"fmt"
	"math"
	"strconv"
)

type best struct {
	heat_loss     int
	straight_left int
}

type node struct {
	heat_loss  int
	best_north []*best
	best_east  []*best
	best_south []*best
	best_west  []*best
}

type Direction int

const (
	North Direction = 0
	East  Direction = 1
	South Direction = 2
	West  Direction = 3
)

type bfsItem struct {
	node_x        int
	node_y        int
	heat_loss     int
	straight_left int
	direction     Direction
}

var bestVal = math.MaxInt

func printGraph(graph [][]*node) {
	for x := range graph {
		for y := range graph[x] {
			fmt.Print(string(rune(graph[x][y].heat_loss)))
		}
		fmt.Println("")
	}
}

func getEndDirections(input_direction Direction, straight_left int) []Direction {
	directions := []Direction{}
	if straight_left > 0 {
		directions = append(directions, input_direction)
	}
	if input_direction == North {
		directions = append(directions, East)
		directions = append(directions, West)
	} else if input_direction == East {
		directions = append(directions, South)
		directions = append(directions, North)
	} else if input_direction == South {
		directions = append(directions, East)
		directions = append(directions, West)
	} else {
		// input_direction == West
		directions = append(directions, South)
		directions = append(directions, North)
	}
	return directions
}

func checkIfShouldTravelInDirection(node *node, input_direction Direction, heat_loss int, straight_left int) bool {
	current_best := node.best_west
	if input_direction == North {
		current_best = node.best_north
	} else if input_direction == East {
		current_best = node.best_east
	} else if input_direction == South {
		current_best = node.best_south
	}

	found := false
	for _, b := range current_best {
		if b.straight_left == straight_left {
			found = true
			if heat_loss < b.heat_loss {
				b.heat_loss = heat_loss
				return true
			}
		}
	}
	if !found {
		current_best = append(current_best, &best{heat_loss, straight_left})
	}
	if input_direction == North {
		node.best_north = current_best
	} else if input_direction == East {
		node.best_east = current_best
	} else if input_direction == South {
		node.best_south = current_best
	} else {
		node.best_west = current_best
	}
	return !found
}

func getLowestHeatLoss(graph [][]*node, x int, y int, direction Direction, heat_loss int, straight_left int) int {
	// fmt.Printf("x:%d y:%d direction:%d heat_loss:%d straight_left: %d\n", x, y, direction, heat_loss, straight_left)
	node_at_position := graph[x][y]

	if x == len(graph)-1 && y == len(graph[0])-1 {
		if heat_loss+node_at_position.heat_loss < bestVal {
			bestVal = heat_loss + node_at_position.heat_loss
		}

		return heat_loss + node_at_position.heat_loss
	}
	output_directions := getEndDirections(direction, straight_left)

	best_heat_loss := math.MaxInt
	var next_heat_loss int
	for _, dir := range output_directions {
		if heat_loss > bestVal {
			return heat_loss
		}
		new_straight_left := straight_left
		if dir == direction {
			new_straight_left -= 1
		} else {
			new_straight_left = 2
		}

		should_travel := checkIfShouldTravelInDirection(node_at_position, dir, heat_loss, new_straight_left)
		if should_travel {
			if dir == North {
				if x > 0 {
					next_heat_loss = getLowestHeatLoss(graph, x-1, y, dir, heat_loss+node_at_position.heat_loss, new_straight_left)
				} else {
					continue
				}
			} else if dir == East {
				if y < len(graph[x])-1 {
					next_heat_loss = getLowestHeatLoss(graph, x, y+1, dir, heat_loss+node_at_position.heat_loss, new_straight_left)
				} else {
					continue
				}
			} else if dir == South {
				if x < len(graph)-1 {
					next_heat_loss = getLowestHeatLoss(graph, x+1, y, dir, heat_loss+node_at_position.heat_loss, new_straight_left)
				} else {
					continue
				}
			} else if dir == West {
				if y > 0 {
					next_heat_loss = getLowestHeatLoss(graph, x, y-1, dir, heat_loss+node_at_position.heat_loss, new_straight_left)
				} else {
					continue
				}
			}
			if next_heat_loss < best_heat_loss {
				best_heat_loss = next_heat_loss
			}
		}
	}
	return best_heat_loss
}

func doTravel(graph [][]*node) int {
	nodes_to_look_at := []bfsItem{{0, 0, 0, 2, East}}
	best_heat_loss := math.MaxInt
	var next bfsItem
	for len(nodes_to_look_at) > 0 {
		next, nodes_to_look_at = nodes_to_look_at[0], nodes_to_look_at[1:]
		if next.node_x == len(graph)-1 && next.node_y == len(graph[0])-1 {
			if next.heat_loss < best_heat_loss {
				best_heat_loss = next.heat_loss
				continue
			}
		} else if next.heat_loss > best_heat_loss {
			continue
		}
		output_directions := getEndDirections(next.direction, next.straight_left)
		for _, dir := range output_directions {
			new_straight_left := next.straight_left
			if dir == next.direction {
				new_straight_left -= 1
			} else {
				new_straight_left = 2
			}

			if dir == North {
				if next.node_x > 0 {
					nodes_to_look_at = append(nodes_to_look_at, bfsItem{
						next.node_x - 1, next.node_y,
						next.heat_loss + graph[next.node_x-1][next.node_y].heat_loss,
						new_straight_left, dir})
				}
			} else if dir == East {
				if next.node_y < len(graph[0])-1 {
					nodes_to_look_at = append(nodes_to_look_at, bfsItem{
						next.node_x, next.node_y + 1,
						next.heat_loss + graph[next.node_x][next.node_y+1].heat_loss,
						new_straight_left, dir})
				}
			} else if dir == South {
				if next.node_x < len(graph)-1 {
					nodes_to_look_at = append(nodes_to_look_at, bfsItem{
						next.node_x + 1, next.node_y,
						next.heat_loss + graph[next.node_x+1][next.node_y].heat_loss,
						new_straight_left, dir})
				}
			} else if dir == West {
				if next.node_y > 0 {
					nodes_to_look_at = append(nodes_to_look_at, bfsItem{
						next.node_x, next.node_y - 1,
						next.heat_loss + graph[next.node_x][next.node_y-1].heat_loss,
						new_straight_left, dir})
				}
			}
		}
	}
	return best_heat_loss
}

func makeGraphCopy(input []string) [][]*node {
	graph := [][]*node{}

	for _, line := range input {
		graph_line := []*node{}
		for _, r := range line {
			num, _ := strconv.Atoi(string(r))
			graph_line = append(graph_line, &node{num, []*best{}, []*best{}, []*best{}, []*best{}})
		}
		graph = append(graph, graph_line)
	}
	return graph
}

func Part1Solution(input []string) string {
	bestVal = math.MaxInt
	graph := makeGraphCopy(input)

	lowest_heat_loss := getLowestHeatLoss(graph, 0, 0, East, -graph[0][0].heat_loss, 3)

	// printGraph(graph)

	return strconv.Itoa(lowest_heat_loss)
}

func Part1SolutionBFS(input []string) string {
	graph := makeGraphCopy(input)

	lowest_heat_loss := doTravel(graph)

	// printGraph(graph)

	return strconv.Itoa(lowest_heat_loss)
}
