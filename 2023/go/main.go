package main

import (
	"day1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day1/input.txt")

	if err != nil {
		fmt.Println("Issue opening file:", err)
		os.Exit(1)
	}

	text, _ := io.ReadAll(f)

	split := strings.Split(string(text), "\n")
	sol := day1.Part2Solution(split)
	fmt.Println("Solution:", sol)
}
