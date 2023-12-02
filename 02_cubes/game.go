package main

import (
	"strconv"
	"strings"
)

type Game struct {
	id int
	results []GameResult
}

type GameResult struct {
	red int
	green int
	blue int
}

func CreateGameFromString(input string) Game {
	game := Game{id: 0, results: []GameResult{}}

	// First separate the id from the set of reveals
	parts := strings.Split(input, ":")
	game.id, _ = strconv.Atoi(strings.Split(parts[0], " ")[1])

	// Then we need to iterate over each reveal
	for _, reveal := range strings.Split(parts[1], ";") {
		gameResult := GameResult{red: 0, green: 0, blue: 0}

		// And parse out the number of each color shown
		for _, colors := range strings.Split(reveal, ",") {

			number_color := strings.Split(strings.TrimSpace(colors), " ")
			number, _ := strconv.Atoi(number_color[0])
			if number_color[1] == "red" {
				gameResult.red = number
			} else if number_color[1] == "green" {
				gameResult.green = number
			} else if number_color[1] == "blue" {
				gameResult.blue = number
			}
		}
		game.results = append(game.results, gameResult)
	}

	return game
}