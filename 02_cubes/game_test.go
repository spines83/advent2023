package main

import "testing"

func setup() Game {
	return CreateGameFromString("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
}

func TestSerializeId(t *testing.T) {

	game := setup()
	
	got := game.id
	want := 1

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSerializeGameLength(t *testing.T) {

	game := setup()
	
	got := len(game.results)
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSerializeRGB(t *testing.T) {

	game := setup()
	
	tests := []struct {
		value int
		expected int
	}{
		{game.results[0].red, 4},
		{game.results[0].green, 0},
		{game.results[0].blue, 3},
		{game.results[1].red, 1},
		{game.results[1].green, 2},
		{game.results[1].blue, 6},
		{game.results[2].red, 0},
		{game.results[2].green, 2},
		{game.results[2].blue, 0},
	}

    for _, tc := range tests {
        if tc.value != tc.expected {
            t.Errorf("Expected %d, got %d", tc.expected, tc.value)
        }
    }
}