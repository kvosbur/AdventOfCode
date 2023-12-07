package adventUtil

import (
	"strconv"
	"strings"
)

func ConvertStringsToInts(input []string) []int {
	dest := []int{}
	for _, str := range input {
		if str == "" {
			continue
		}
		num, _ := strconv.Atoi(strings.Trim(str, " "))
		dest = append(dest, num)
	}
	return dest
}
