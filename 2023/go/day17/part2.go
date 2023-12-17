package day17

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

var count = 0

func getEndDirectionsPart2(input_direction Direction, straight_done int) []Direction {
	directions := []Direction{}
	if straight_done < 10 {
		directions = append(directions, input_direction)
	}
	if straight_done > 3 {
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
	}
	return directions
}

func getLowestHeatLossPart2(graph [][]*node, x int, y int, direction Direction, heat_loss int, straight_done int) int {
	// fmt.Printf("x:%d y:%d direction:%d heat_loss:%d straight_left: %d\n", x, y, direction, heat_loss, straight_done)
	node_at_position := graph[x][y]

	if x == len(graph)-1 && y == len(graph[0])-1 {
		if straight_done < 4 {
			return math.MaxInt
		}
		count++
		if heat_loss+node_at_position.heat_loss < bestVal {
			bestVal = heat_loss + node_at_position.heat_loss
		}
		if count%100 == 0 {
			fmt.Println(count, bestVal, time.Now())
		}
		// fmt.Println("debug end", heat_loss+node_at_position.heat_loss)

		return heat_loss + node_at_position.heat_loss
	}
	output_directions := getEndDirectionsPart2(direction, straight_done)

	best_heat_loss := math.MaxInt
	var next_heat_loss int
	for _, dir := range output_directions {
		if heat_loss > bestVal {
			return heat_loss
		}
		new_straight_done := straight_done
		if dir == direction {
			new_straight_done += 1
		} else {
			new_straight_done = 1
		}

		should_travel := checkIfShouldTravelInDirection(node_at_position, dir, heat_loss, new_straight_done)
		if should_travel {
			if dir == North {
				if x > 0 {
					next_heat_loss = getLowestHeatLossPart2(graph, x-1, y, dir, heat_loss+node_at_position.heat_loss, new_straight_done)
				} else {
					continue
				}
			} else if dir == East {
				if y < len(graph[x])-1 {
					next_heat_loss = getLowestHeatLossPart2(graph, x, y+1, dir, heat_loss+node_at_position.heat_loss, new_straight_done)
				} else {
					continue
				}
			} else if dir == South {
				if x < len(graph)-1 {
					next_heat_loss = getLowestHeatLossPart2(graph, x+1, y, dir, heat_loss+node_at_position.heat_loss, new_straight_done)
				} else {
					continue
				}
			} else if dir == West {
				if y > 0 {
					next_heat_loss = getLowestHeatLossPart2(graph, x, y-1, dir, heat_loss+node_at_position.heat_loss, new_straight_done)
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

func Part2Solution(input []string) string {
	graph := makeGraphCopy(input)
	bestVal = math.MaxInt

	lowest_heat_loss := getLowestHeatLossPart2(graph, 0, 0, East, -graph[0][0].heat_loss, 0)

	return strconv.Itoa(lowest_heat_loss)
}
