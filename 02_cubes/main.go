package main

import (
	"advent2023/utils"
	"fmt"
)

func main() {
	lines, _ := utils.IterLines("inputs/part1.txt")
	
	sum_valid := 0
	sum_power := 0

	for currentLine := range lines {
		game := CreateGameFromString(currentLine)

		valid := true
		min_red := 0
		min_green := 0
		min_blue := 0

		for _, result := range game.results {

			// Part 1
			if result.red > 12 || result.green > 13 || result.blue > 14 {
				valid = false
			}

			// Part 2
			if result.red > min_red {
				min_red = result.red
			}
			if result.green > min_green {
				min_green = result.green
			}
			if result.blue > min_blue {
				min_blue = result.blue
			}
		}

		if valid {
			sum_valid += game.id
		}
		sum_power += min_red * min_green * min_blue
	}
	fmt.Println("ValidIdSum: ", sum_valid)
	fmt.Println("Power: ", sum_power)

}

