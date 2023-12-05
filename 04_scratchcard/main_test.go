package main

import "testing"

func TestLottery01(t *testing.T) {

	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"

	_, got := CalculateTicketPoints(input)
	want := 8

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}


func TestLotteryDupes(t *testing.T) {

	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53 83 83 83 83"

	_, got := CalculateTicketPoints(input)
	want := 8

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestFullLottery01(t *testing.T) {

	input := "Card   1:  5 27 94 20 50  7 98 41 67 34 | 34  9 20 90  7 77 44 71 27 12 98  1 79 96 24 51 25 84 67 41  5 13 78 31 26"

	_, got := CalculateTicketPoints(input)
	want := 128

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}


