package day2

import (
	"strconv"
	"strings"
)

func getGameNum(gameIdent string) int {
	split := strings.Split(gameIdent, " ")
	res, _ := strconv.Atoi(split[1])
	return res
}

func isGameValid(game string) bool {
	colored_cubes := strings.Split(game, ",")
	for _, cube_str := range colored_cubes {
		split_cube := strings.Split(strings.Trim(cube_str, " "), " ")
		num, _ := strconv.Atoi(split_cube[0])
		if strings.Contains(cube_str, "red") {
			if num > 12 {
				return false
			}
		}
		if strings.Contains(cube_str, "green") {
			if num > 13 {
				return false
			}
		}
		if strings.Contains(cube_str, "blue") {
			if num > 14 {
				return false
			}
		}
	}
	return true
}

func Part1Solution(input []string) string {

	sum := 0
	for _, line := range input {
		items := strings.Split(line, ":")
		game_num := getGameNum(items[0])
		games := strings.Split(items[1], ";")

		isValid := true
		for _, game := range games {
			isValid = isValid && isGameValid(game)
		}

		if isValid {
			sum += game_num
		}
	}
	return strconv.Itoa(sum)
}
