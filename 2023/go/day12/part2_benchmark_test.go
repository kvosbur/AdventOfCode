package day12_test

import (
	"day12"
	"strings"
	"testing"
)

func BenchmarkLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day12.Part2Solution(strings.Split(`?????.??##?????????. 2,6,2`, "\n"))
	}
}
