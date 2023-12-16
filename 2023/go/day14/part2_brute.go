package day14

import (
	"fmt"
	"strconv"
	"time"
)

func Part2BruteSolution(input []string) string {
	input_split := [][]rune{}
	cycle_count := 1000000000
	for _, val := range input {
		temp := []rune{}
		for _, char := range val {
			temp = append(temp, char)
		}
		input_split = append(input_split, temp)
	}
	fmt.Println("++++++++++++++")

	// day11.PrintInput(input_split)
	for i := 0; i < cycle_count; i++ {
		shiftAllNorth(input_split)
		shiftAllWest(input_split)
		shiftAllSouth(input_split)
		shiftAllEast(input_split)
		if i%100000 == 0 {
			fmt.Println(float64(i)/float64(cycle_count), time.Now())
		}
	}

	value := scoreNorth(input_split)
	return strconv.Itoa(value)
}
