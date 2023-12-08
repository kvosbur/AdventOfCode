package day7

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var card_strength_2 = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

type rankedHandList2 []rankedHand

func (rhl rankedHandList2) Len() int {
	return len(rhl)
}

func (rhl rankedHandList2) Less(i, j int) bool {
	if rhl[i].rank == rhl[j].rank {
		for index, card_i := range rhl[i].h.cards {
			card_j := rhl[j].h.cards[index]
			if card_i != card_j {
				return slices.Index(card_strength_2, card_i) > slices.Index(card_strength_2, card_j)
			}
		}
	}
	return rhl[i].rank > rhl[j].rank
}

func (rhl rankedHandList2) Swap(i, j int) {
	rhl[i], rhl[j] = rhl[j], rhl[i]
}

func rankHand2(h hand) rankedHand {
	cards_count := map[string]int{}
	for _, card := range h.cards {
		_, found := cards_count[card]
		if !found {
			cards_count[card] = 1
		} else {
			cards_count[card] += 1
		}
	}

	count_5 := 0
	count_4 := 0
	count_3 := 0
	count_2 := 0
	count_J := 0
	for key, val := range cards_count {
		if key == "J" {
			count_J += val
		} else if val == 5 {
			count_5 += 1
		} else if val == 4 {
			count_4 += 1
		} else if val == 3 {
			count_3 += 1
		} else if val == 2 {
			count_2 += 1
		}
	}

	// 5 of a kind
	if count_5 == 1 || (count_4 == 1 && count_J == 1) || (count_3 == 1 && count_J == 2) || (count_2 == 1 && count_J == 3) ||
		count_J == 4 || count_J == 5 {
		return rankedHand{h, 6}
	}
	// 4 of a kind
	if count_4 == 1 || (count_3 == 1 && count_J == 1) || (count_2 == 1 && count_J == 2) || (count_2 == 0 && count_J == 3) {
		return rankedHand{h, 5}
	}
	// full house
	if count_3 == 1 && count_2 == 1 || (count_2 == 2 && count_J == 1) {
		return rankedHand{h, 4}
	}
	// 3 of a kind
	if count_3 == 1 || (count_2 == 1 && count_J == 1) || count_J == 2 {
		return rankedHand{h, 3}
	}
	// 2 pairs
	if count_2 == 2 {
		return rankedHand{h, 2}
	}
	//1 pair
	if count_2 == 1 || count_J == 1 {
		return rankedHand{h, 1}
	}
	// high card
	return rankedHand{h, 0}
}

func Part2Solution(input []string) string {
	hands := []hand{}

	for _, line := range input {
		split_line := strings.Split(line, " ")
		bid, _ := strconv.Atoi(split_line[1])
		cards := strings.Split(split_line[0], "")
		hands = append(hands, hand{cards, bid})
	}

	ranked_hands := rankedHandList2{}
	for _, h := range hands {
		ranked_hands = append(ranked_hands, rankHand2(h))
	}

	sort.Sort(ranked_hands)

	winnings := 0
	rank := len(ranked_hands)
	for _, rh := range ranked_hands {
		fmt.Println(rh)
		winnings += (rh.h.bid * rank)
		rank -= 1
	}
	return strconv.Itoa(winnings)
}
