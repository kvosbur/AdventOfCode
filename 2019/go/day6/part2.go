package day6

import (
	"fmt"
	"strconv"
	"strings"
)

var santa string = "SAN"
var me string = "YOU"

func findConnector(so *spaceObject) (int, int, bool) {
	if so.identifier == santa {
		fmt.Println("found santa", so)
		return 1, 0, false
	} else if so.identifier == me {
		fmt.Println("found me", so)
		return 0, 1, false
	}

	meDistance := 0
	santaDistance := 0
	// fmt.Println(so.identifier)
	for _, child := range so.orbits {
		// fmt.Println(so.identifier, "->", child.identifier)
		a, b, found := findConnector(child)
		meDistance += a
		santaDistance += b
		if found {
			// fmt.Println("found", a, b)
			return a, b, found
		} else if meDistance > 0 && santaDistance > 0 {
			return meDistance, santaDistance, true
		}
	}

	if meDistance > 0 {
		meDistance += 1
	}
	if santaDistance > 0 {
		santaDistance += 1
	}

	return meDistance, santaDistance, false
}

func Part2Solution(input []string) string {
	spaceObjects := make(map[string]*spaceObject)

	for _, item := range input {
		split := strings.Split(item, ")")

		first, found := spaceObjects[split[0]]
		if !found {
			first = &spaceObject{
				orbits:     nil,
				identifier: split[0],
			}
			spaceObjects[split[0]] = first
		}

		second, found := spaceObjects[split[1]]
		if !found {
			second = &spaceObject{
				orbits:     nil,
				identifier: split[1],
			}
			spaceObjects[split[1]] = second
		}

		first.orbits = append(first.orbits, second)
	}

	start := spaceObjects["COM"]
	a, b, _ := findConnector(start)

	return strconv.Itoa(a + b - 2)
}

// answer 547
