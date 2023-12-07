package day6

import (
	"adventUtil"
	"fmt"
	"strconv"
	"strings"
)

func getWinCount(time int, record int) int {
	count := 0
	for wait := 0; wait < time; wait++ {
		time_remaining := time - wait
		distance := time_remaining * wait
		if distance > record {
			count++
		}
	}
	return count
}

func Part1Solution(input []string) string {

	length_all := strings.Split(input[0], ":")
	lengths := adventUtil.ConvertStringsToInts(strings.Split(length_all[1], " "))

	records_all := strings.Split(input[1], ":")
	records := adventUtil.ConvertStringsToInts(strings.Split(records_all[1], " "))

	fmt.Println(lengths)
	fmt.Println(records)

	answer := 1
	for i, length := range lengths {
		record := records[i]
		answer *= getWinCount(length, record)
	}
	return strconv.Itoa(answer)
}
