package day2

import (
	"strconv"
	"strings"
)

type gameCubes struct {
	red   int
	blue  int
	green int
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (gc gameCubes) maxCombine(other gameCubes) gameCubes {
	return gameCubes{
		red:   maxInt(gc.red, other.red),
		green: maxInt(gc.green, other.green),
		blue:  maxInt(gc.blue, other.blue),
	}
}

func (gc gameCubes) getPower() int {
	return gc.red * gc.blue * gc.green
}

func getGameCubes(game string) gameCubes {
	colored_cubes := strings.Split(game, ",")
	gameCube := gameCubes{}
	for _, cube_str := range colored_cubes {
		split_cube := strings.Split(strings.Trim(cube_str, " "), " ")
		num, _ := strconv.Atoi(split_cube[0])
		if strings.Contains(cube_str, "red") {
			gameCube.red = num
		}
		if strings.Contains(cube_str, "green") {
			gameCube.green = num
		}
		if strings.Contains(cube_str, "blue") {
			gameCube.blue = num
		}
	}
	return gameCube
}

func Part2Solution(input []string) string {
	sum := 0
	for _, line := range input {
		items := strings.Split(line, ":")
		games := strings.Split(items[1], ";")

		isValid := true
		current_max := gameCubes{}
		for _, game := range games {
			current_max = current_max.maxCombine(getGameCubes(game))
		}

		if isValid {
			sum += current_max.getPower()
		}
	}
	return strconv.Itoa(sum)
}
