package main

import (
	"advent2023/utils"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time int
	distance int
}


func GetRaceListFromFile(path string) []Race {

	resp := []Race{}
	re := regexp.MustCompile(`[0-9]+`)

	lines, _ := utils.IterLines(path)
	for currentLine := range lines {
		if strings.Contains(currentLine, "Time") {
			matches := re.FindAllString(currentLine, -1)
			for _, match := range matches {
				match_int, _ := strconv.Atoi(match)
				resp = append(resp, Race{time: match_int, distance: 0})
			}
		} else if strings.Contains(currentLine, "Distance") {
			matches := re.FindAllString(currentLine, -1)
			for i, match := range matches {
				match_int, _ := strconv.Atoi(match)
				resp[i].distance = match_int
			}
		}
	}

	return resp
}