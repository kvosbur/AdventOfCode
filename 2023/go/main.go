package main

import (
	"day17"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day17/input.txt")

	if err != nil {
		fmt.Println("Issue opening file:", err)
		os.Exit(1)
	}

	text, _ := io.ReadAll(f)

	split := strings.Split(string(text), "\n")
	sol := day17.Part1Solution(split)
	fmt.Println("Solution:", sol)
}
