package main

import (
	"advent2023/utils"
	"fmt"
	"strconv"
	"regexp"
	"strings"
)

func main() {
	runCalibration("inputs/sample_part1.txt")
	runCalibration("inputs/part1.txt")
	runCalibrationWithReplace("inputs/sample_part2.txt", true)
	runCalibrationWithReplace("inputs/part1.txt", true)
}

// To maintain backwards compatiability with part 1
func runCalibration(path string) {
	runCalibrationWithReplace(path, false)
}

func runCalibrationWithReplace(path string, replaceText bool) {
	lines, _ := utils.IterLines(path)
	totalCalibration := 0
	for currentLine := range lines {
		totalCalibration += getCalibrationValueWithReplace(currentLine, replaceText)
	}
	fmt.Printf("%d\n", totalCalibration)
}

// To maintain backwards compatiability with part 1
func getCalibrationValue(input string) int {
	return getCalibrationValueWithReplace(input, false)
}

func getCalibrationValueWithReplace(input string, replaceText bool) int {

	var textToCalibrate string
	if replaceText {
		textToCalibrate = replaceStringDigits(input)
	} else {
		textToCalibrate = input
	}
	digits, _ := getDigits(textToCalibrate)
	return 10 * digits[0] + digits[len(digits) - 1]
}

func replaceStringDigits(input string) string {
	output := strings.Replace(input, "one", "on1ne", -1)
	output = strings.Replace(output, "two", "tw2wo", -1)
	output = strings.Replace(output, "three", "thre3hree", -1)
	output = strings.Replace(output, "four", "fou4our", -1)
	output = strings.Replace(output, "five", "fiv5ive", -1)
	output = strings.Replace(output, "six", "si6ix", -1)
	output = strings.Replace(output, "seven", "seve7even", -1)
	output = strings.Replace(output, "eight", "eigh8ight", -1)
	output = strings.Replace(output, "nine", "nin9ine", -1)
	return output
}

func getDigits(s string) ([]int, error) {
    var digits []int
    re := regexp.MustCompile(`\d`)
    matches := re.FindAllString(s, -1)
    for _, match := range matches {
        digit, err := strconv.Atoi(match)
        if err != nil {
            return nil, err
        }
        digits = append(digits, digit)
    }
    return digits, nil
}