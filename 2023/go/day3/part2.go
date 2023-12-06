package day3

import (
	"slices"
	"strconv"
)

type gear struct {
	x       int
	y       int
	numbers []int
}

type loc struct {
	x int
	y int
}

func isGearSymbol(character rune) bool {
	return character == '*'
}

func hasAdjacentGearSymbol(input []string, x int, y int) (bool, loc) {
	// top left
	if x > 0 && y > 0 && isGearSymbol(rune(input[x-1][y-1])) {
		return true, loc{x - 1, y - 1}
	}
	// top
	if x > 0 && isGearSymbol(rune(input[x-1][y])) {
		return true, loc{x - 1, y}
	}
	// top right
	if x > 0 && y < len(input[0])-1 && isGearSymbol(rune(input[x-1][y+1])) {
		return true, loc{x - 1, y + 1}
	}
	// right
	if y < len(input[0])-1 && isGearSymbol(rune(input[x][y+1])) {
		return true, loc{x, y + 1}
	}
	// bottom right
	if x < len(input)-1 && y < len(input[0])-1 && isGearSymbol(rune(input[x+1][y+1])) {
		return true, loc{x + 1, y + 1}
	}
	// bottom
	if x < len(input)-1 && isGearSymbol(rune(input[x+1][y])) {
		return true, loc{x + 1, y}
	}
	// bottom left
	if x < len(input)-1 && y > 0 && isGearSymbol(rune(input[x+1][y-1])) {
		return true, loc{x + 1, y - 1}
	}
	// left
	if y > 0 && isGearSymbol(rune(input[x][y-1])) {
		return true, loc{x, y - 1}
	}

	return false, loc{}
}

func updateGears(x int, y int, number int, gears []*gear) []*gear {
	// fmt.Println(x, y, number, gears)
	for _, g := range gears {
		if g.x == x && g.y == y {
			g.numbers = append(g.numbers, number)
			return gears
		}
	}
	new_gears := append(gears, &gear{x, y, []int{number}})
	return new_gears
}

func filterDuplicateGearLocations(gears []gear) []gear {
	new_gears := []gear{}
	for _, g := range gears {
		already_exists := false
		for _, new_g := range new_gears {
			if g.x == new_g.x && g.y == new_g.y {
				already_exists = true
			}
		}
		if !already_exists {
			new_gears = append(new_gears, g)
		}
	}
	return new_gears
}

func Part2Solution(input []string) string {
	gears := []*gear{}
	for x, line := range input {
		current_num := ""
		num_should_be_counted := false
		gear_locations := []gear{}
		for y, character := range line {
			if slices.Contains(digits, character) {
				current_num += string(character)
				b, location := hasAdjacentGearSymbol(input, x, y)
				if b {
					num_should_be_counted = true
					gear_locations = append(gear_locations, gear{x: location.x, y: location.y})
				}
			} else if len(current_num) > 0 {
				if num_should_be_counted {
					num, _ := strconv.Atoi(current_num)
					filtered := filterDuplicateGearLocations(gear_locations)
					for _, g := range filtered {
						gears = updateGears(g.x, g.y, num, gears)
					}
				}
				current_num = ""
				num_should_be_counted = false
				gear_locations = []gear{}
			}
		}
		if len(current_num) > 0 {
			if num_should_be_counted {
				num, _ := strconv.Atoi(current_num)
				for _, g := range gear_locations {
					gears = updateGears(g.x, g.y, num, gears)
				}
			}
		}
	}

	sum := 0
	for _, g := range gears {
		if len(g.numbers) == 2 {
			sum += (g.numbers[0] * g.numbers[1])
		}
	}

	return strconv.Itoa(sum)
}
