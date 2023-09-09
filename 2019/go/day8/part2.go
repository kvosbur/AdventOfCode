package day8

import (
	"fmt"
	"strconv"
)

func createTransparentImage(width int, height int) [][]int {
	image := [][]int{}
	for y := 0; y < height; y++ {
		row := []int{}
		for x := 0; x < width; x++ {
			row = append(row, 2)
		}
		image = append(image, row)
	}
	return image
}

func Part2Solution(input []string) string {
	width := 25
	height := 6

	full := input[0]
	index := 0

	resultImage := createTransparentImage(width, height)

	for {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if resultImage[y][x] == 2 {
					val, _ := strconv.Atoi(string(full[index]))
					resultImage[y][x] = val
				}
				index++
			}
		}

		if index >= len(full)-1 {
			break
		}
	}

	for _, row := range resultImage {
		for _, element := range row {
			temp := strconv.Itoa(element)
			if temp == "0" {
				temp = " "
			}
			fmt.Print(temp)
		}
		fmt.Print("\n")
	}
	return ""
}

// answer is KYHFE
