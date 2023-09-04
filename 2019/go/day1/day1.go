package day1

import (
	"fmt"
	"strconv"
)

func Day1Part1Solution(input []string) string {
	sum := 0
	for _, entry := range input {
		numeric, err := strconv.Atoi(entry)
		if err != nil {
			fmt.Println("Issue converting to int", entry, err)
		}
		sum += numeric/3 - 2
	}

	return strconv.Itoa(sum)
}

// 3368364 is correct

func Day1Part2Solution(input []string) string {
	sum := 0
	for _, entry := range input {
		numeric, err := strconv.Atoi(entry)
		if err != nil {
			fmt.Println("Issue converting to int", entry, err)
		}
		for {
			temp := numeric/3 - 2
			if temp <= 0 {
				break
			}
			sum += temp
			numeric = temp
		}
	}

	return strconv.Itoa(sum)
}

// 5049684 is correct
