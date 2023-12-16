package day15

import (
	"slices"
	"strconv"
	"strings"
)

type sequence struct {
	label string
	focul int
}

func removeLabel(box_map map[int][]sequence, seq sequence) {
	h := hash(seq.label)
	entries, _ := box_map[h]
	for i, entry := range entries {
		if entry.label == seq.label {
			box_map[h] = slices.Delete(entries, i, i+1)
			return
		}
	}
}

func addLabel(box_map map[int][]sequence, seq sequence) {
	h := hash(seq.label)
	entries, _ := box_map[h]
	for i, entry := range entries {
		if entry.label == seq.label {
			entries[i] = seq
			box_map[h] = entries
			return
		}
	}
	box_map[h] = append(entries, seq)
}

func scoreBoxes(box_map map[int][]sequence) int {
	score := 0
	for key, val := range box_map {
		multiplier := key + 1
		for i, seq := range val {
			score += (multiplier * (i + 1) * seq.focul)
		}
	}
	return score
}

func Part2Solution(input []string) string {
	split_line := strings.Split(input[0], ",")
	box_map := map[int][]sequence{}
	for i := 0; i < 256; i++ {
		box_map[i] = []sequence{}
	}
	for _, seq := range split_line {
		if strings.Contains(seq, "-") {
			split_seq := strings.Split(seq, "-")
			removeLabel(box_map, sequence{label: split_seq[0]})
		} else {
			split_seq := strings.Split(seq, "=")
			num, _ := strconv.Atoi(split_seq[1])
			addLabel(box_map, sequence{split_seq[0], num})
		}
	}

	score := scoreBoxes(box_map)

	return strconv.Itoa(score)
}
