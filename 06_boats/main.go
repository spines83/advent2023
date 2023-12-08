package main

import (
	"fmt"
	"strconv"
)

func main() {
	race_list := GetRaceListFromFile("inputs/part1.txt")
	
	// Part 1
	counter_mul := 1
	for _, race := range race_list {
		win_opts := countWinOptions(race) 
		counter_mul *= win_opts
	}
	fmt.Println("p1: ", counter_mul)

	// Part 2
	big_time := ""
	big_dist := ""
	for _, race := range race_list {
		big_time = big_time + strconv.Itoa(race.time)
		big_dist = big_dist + strconv.Itoa(race.distance)
	}
	big_time_int, _ := strconv.Atoi(big_time)
	big_dist_int, _ := strconv.Atoi(big_dist)
	fmt.Println("p2: ", countWinOptions(Race{time: big_time_int, distance: big_dist_int}))
}

func countWinOptions(race Race) int {
	counter := 0
	for button := 0; button <= race.time; button++ {
		if button * (race.time - button) > race.distance {
			counter += 1
		}
	}
	return counter
}