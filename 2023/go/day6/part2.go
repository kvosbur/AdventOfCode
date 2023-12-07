package day6

import (
	"fmt"
	"strconv"
	"strings"
)

func Part2Solution(input []string) string {
	length_all := strings.Split(input[0], ":")
	length, _ := strconv.Atoi(strings.Replace(length_all[1], " ", "", -1))

	records_all := strings.Split(input[1], ":")
	record, _ := strconv.Atoi(strings.Replace(records_all[1], " ", "", -1))

	fmt.Println(length)
	fmt.Println(record)

	answer := getWinCount(length, record)
	return strconv.Itoa(answer)
}
