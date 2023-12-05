package main

import (
	"advent2023/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main () {

	total_score := 0
	num_tickets := 0

	ticket_hits := make(map[int]int)
	ticket_counts := make(map[int]int)
	
	// Part 1
	lines, _ := utils.IterLines("inputs/part1.txt")
	for currentLine := range lines {
		id, hits := CalculateTicketPoints(currentLine)
		ticket_hits[id] = hits
		ticket_counts[id] = 1
		total_score += getPoints(hits)
		num_tickets += 1
	}

	// Part 2
	// for each ticket 1..n
	for id := 1; id <= num_tickets; id++ {
		
		// how many times are we scoring the ticket? 
		multiplier := ticket_counts[id]

		// how many copies of future tickets do we need to add
		for i := 1; i <= ticket_hits[id]; i++ {
			ticket_counts[id + i] = ticket_counts[id + i] + multiplier
		}
	}

	total_tickets := 0
	for _, val := range ticket_counts {
		total_tickets += val
	}

	fmt.Println(total_score)
	fmt.Println(total_tickets)
}

func CalculateTicketPoints(input string) (int, int) {
	id_card := strings.Split(input, ":")
	winner_player := strings.Split(id_card[1], "|")
	re := regexp.MustCompile(`[0-9]+`)

	// Ticket id
	matches := re.FindAllString(id_card[0], -1)
	ticket_id, _ := strconv.Atoi(matches[0])

	// Winner half of the ticket
	winner_map := make(map[int]bool)
	matches = re.FindAllString(winner_player[0], -1)
	for _, winner := range matches {
		winner_int, _ := strconv.Atoi(winner)
		winner_map[winner_int] = true
	}

	// Player half of the ticket
	hits := 0
	matches = re.FindAllString(winner_player[1], -1)
	for _, player := range matches {
		player_int, _ := strconv.Atoi(player)
		_, exists := winner_map[player_int]
		if exists {
			hits += 1
			delete(winner_map, player_int)
		}
	}

	return ticket_id, hits
}

func getPoints(hits int) int {
	if hits == 0 {
		return 0
	} else {
		return int(math.Pow(2, float64(hits) - 1))
	}
}