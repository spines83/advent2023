package main

import (
	"sort"
	"strconv"
	"strings"
)

type HandStrength int
const (
	HIGH_CARD    HandStrength = iota + 1 // EnumIndex = 1
	ONE_PAIR                             // EnumIndex = 2
	TWO_PAIR                             // EnumIndex = 3
	THREE_KIND                           // EnumIndex = 4
	FULL_HOUSE                           // EnumIndex = 5
	FOUR_KIND                            // EnumIndex = 6
	FIVE_KIND                            // EnumIndex = 7
)

type Hand struct {
	cards []int
	strength HandStrength
	bid int
}

func ParseHand(input string, use_joker_rule bool) Hand {
	parts := strings.Split(input, " ")
	hand := Hand{cards: []int{}, strength: 0}
	for _, card := range parts[0] {
		hand.cards = append(hand.cards, getIntValue(card, use_joker_rule))
	}
	hand.strength = getHandStrength(hand.cards)
	hand.bid, _ = strconv.Atoi(parts[1])
	return hand
}

func SortHands(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].strength > hands[j].strength {
			return false
		} else if hands[i].strength < hands[j].strength {
			return true
		} else {
			for id := 0; id < len(hands[i].cards); id++ {
				if hands[i].cards[id] > hands[j].cards[id] {
					return false
				} else if hands[i].cards[id] < hands[j].cards[id] {
					return true
				}
			}
		}
		return true
	})
}

func getHandStrength(cards[] int) HandStrength {
	card_counts := make(map[int]int)
	joker_count := 0
	for _, card := range cards {
		// Part 2 - Capture Jokers
		if card == 0 {
			joker_count += 1
		} else {
			if _, exists := card_counts[card]; !exists {
				card_counts[card] = 0
			}
			card_counts[card] += 1
		}
	}

	// Extract + sort values from the above map
	cardinality := []int{}
	for _, v := range card_counts {
		cardinality = append(cardinality, v)
	}
	sort.Ints(cardinality)

	// Part 2 - Add Jokers to highest cardinality 
	if len(cardinality) > 0 {
		cardinality[len(cardinality) - 1] += joker_count
	} else if joker_count > 0 {
		cardinality = append(cardinality, joker_count)
	}

	if len(cardinality) == 1 && cardinality[0] == 5 {
		return FIVE_KIND // [5]
	} else if len(cardinality) == 2 && cardinality[1] == 4 {
		return FOUR_KIND // [1, 4]
	} else if len(cardinality) == 2 && cardinality[1] == 3 {
		return FULL_HOUSE // [2, 3]
	} else if len(cardinality) == 3 && cardinality[2] == 3 {
		return THREE_KIND // [1, 1, 3]
	} else if len(cardinality) == 3 && cardinality[2] == 2 {
		return TWO_PAIR // [1, 2, 2]
	} else if len(cardinality) == 4 && cardinality[3] == 2 {
		return ONE_PAIR // [1, 1, 1, 2]
	} else {
		return HIGH_CARD // [1, 1, 1, 1, 1]
	}
}

func getIntValue(card rune, use_joker_rule bool) int {
	if card == 'T' {
		return 10 
	} else if card == 'J' {
		if use_joker_rule {
			return 0
		} else {
			return 11 
		}
	} else if card == 'Q' {
		return 12
	} else if card == 'K' {
		return 13 
	} else if card == 'A' {
		return 14 
	} else {
		return int(card - '0')
	}
}