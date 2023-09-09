package day8

import (
	"fmt"
	"strconv"
)

func getValueOfImage(image [][]byte) int {
	counts := map[byte]int{}
	for _, row := range image {
		for _, column := range row {
			entry, ok := counts[column]
			if !ok {
				entry = 0
			}
			counts[column] = entry + 1
		}
	}
	return counts[byte('1')] * counts[byte('2')]
}

func Part1Solution(input []string) string {
	width := 25
	height := 6

	bestImage := [][]byte{}
	bestCount := 500

	full := input[0]
	index := 0

	for {
		image := [][]byte{}
		count := 0
		for y := 0; y < height; y++ {
			row := []byte{}
			for x := 0; x < width; x++ {
				row = append(row, full[index])
				if full[index] == byte('0') {
					count++
				}
				index++
			}
			image = append(image, row)
		}

		if count < bestCount {
			bestImage = image
			bestCount = count
		}
		if index >= len(full)-1 {
			break
		}
	}
	fmt.Println(bestImage)
	fmt.Println(bestCount)
	return strconv.Itoa(getValueOfImage(bestImage))
}

// answer is 1572
