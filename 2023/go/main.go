package main

import (
	"day12"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day12/input.txt")

	if err != nil {
		fmt.Println("Issue opening file:", err)
		os.Exit(1)
	}

	text, _ := io.ReadAll(f)

	split := strings.Split(string(text), "\n")
	sol := day12.Part2Solution(split)
	fmt.Println("Solution:", sol)
}
