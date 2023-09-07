package day6

import (
	"strconv"
	"strings"
)

type spaceObject struct {
	orbits     []*spaceObject
	identifier string
	checksum   int
}

func Part1Solution(input []string) string {
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

	next := []*spaceObject{}
	index := 0
	for _, val := range spaceObjects {
		next = append(next, val.orbits...)
		for {
			if index == len(next) {
				break
			}
			objPointer := next[index]
			objPointer.checksum += 1
			next = append(next, objPointer.orbits...)

			index++
		}
	}

	sum := 0
	for _, val := range spaceObjects {
		sum += val.checksum
	}
	return strconv.Itoa(sum)
}

// answer 301100
