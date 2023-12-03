package utils

import (
	"bufio"
	"os"
)

func IterLines(path string) (<-chan string, error) {
	readFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	chn := make(chan string)
	go func() {
		for fileScanner.Scan() {
			chn <- fileScanner.Text()
		}
		readFile.Close()
		close(chn)
	}()

	return chn, nil
}

// Utility for problems where we need a 2d matrix
func ParseFileIntoMatrix(path string) [][]rune {

	// Open the file
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Initialize a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Initialize a dynamic 2D slice of characters
	var matrix [][]rune

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return matrix
}

// Quick and dirty utility to count lines in a file
func CountLines(inputFile string) int {
	file, _ := os.Open(inputFile)
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	return lineCount
}