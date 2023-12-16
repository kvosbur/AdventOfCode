package day14_test

import (
	"day14"
	"strings"
	"testing"
)

func BenchmarkLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day14.Part2BruteSolution(strings.Split(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`, "\n"))
	}
}
