package main

import (
	"day3"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day3/input.txt")

	if err != nil {
		fmt.Println("Issue opening file:", err)
		os.Exit(1)
	}

	text, _ := io.ReadAll(f)

	split := strings.Split(string(text), "\r\n")
	sol := day3.Part2Solution(split)
	fmt.Println("Solution:", sol)
}
