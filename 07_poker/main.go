package main

import (
	"advent2023/utils"
	"fmt"
)

func main() {
	lines, _ := utils.IterLines("inputs/input1.txt")
	hands := []Hand{}
	for currentLine := range lines {
		hand := ParseHand(currentLine, true)
		hands = append(hands, hand)
	}

	SortHands(hands)

	winnings := 0
	for i, hand := range hands {
		for _, card := range hand.cards {
			fmt.Printf("%d,", card)
		}
		fmt.Printf(" -- %d, %d\n", hand.strength, hand.bid)

		winnings += hand.bid * (i + 1)
	}

	fmt.Println(winnings)
}