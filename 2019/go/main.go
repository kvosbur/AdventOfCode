package main

import (
	"day8"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day8/input.txt")

	if err != nil {
		fmt.Println("Issue opening file:", err)
		os.Exit(1)
	}

	text, _ := io.ReadAll(f)

	split := strings.Split(string(text), "\r\n")
	sol := day8.Part2Solution(split)
	fmt.Println("Solution:", sol)
}
