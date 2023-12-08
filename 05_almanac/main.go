package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"sync"
)

func main() {
	input := "inputs/part1.txt"
	almanac := BuildAlmanac(input)
	re := regexp.MustCompile(`[0-9]+`)

	var wg sync.WaitGroup
    var mu sync.Mutex

	// Read the first line to grab seeds
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)

		smallest_p1 := math.MaxInt
		smallest_p2 := math.MaxInt
		for i, match := range matches {
			match_int, _ := strconv.Atoi(match)
			
			// Part 1
			filtered_num = FilterNumber(match_int, almanac)
			if filtered_num < smallest_p1 {
				smallest_p1 = filtered_num
			}

			// Part 2 - Runs in O(my poor CPU)
			if i % 2 == 1 {
				start, _ := strconv.Atoi(matches[i - 1])
				len , _ := strconv.Atoi(match)
				fmt.Println(start, len)
				for seed := start; seed < start + len; seed++ {
					wg.Add(1)
					go func(seed int) {
						defer wg.Done()
						filtered_num = FilterNumber(seed, almanac)
						mu.Lock()
						if filtered_num < smallest_p2 {
							smallest_p2 = filtered_num
						}
						mu.Unlock()
					}(seed)
				}
			}
		}

		wg.Wait()
		fmt.Println(smallest_p1)
		fmt.Println(smallest_p2)
	}
}