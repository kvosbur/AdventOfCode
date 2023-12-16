package day12_test

import (
	"day12"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sol := day12.Part2BruteSolution(strings.Split(`?????.??##?????????. 2,6,2`, "\n"))
		fmt.Println(sol == "1687626340")
	}
}
