package main

import (
	"advent2023/utils"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	// mtx := utils.ParseFileIntoMatrix("inputs/sample_part1.txt")
	mtx := utils.ParseFileIntoMatrix("inputs/part1.txt")
	fmt.Println(CalculateSchematicSum(mtx))
}

type Coordinate struct {
	row, column int
}

type Symbol struct {
	row int
	column int
	char rune
}

func CalculateSchematicSum(mtx [][] rune) (int, int) {

	// Pre-process a symbols map
	symbols := make(map[Coordinate]rune)
	gears := make(map[Coordinate][]int)
	for i, line := range mtx {
		for j, letter := range line {
			if !unicode.IsDigit(letter) && letter != '.' {
				symbols[Coordinate{i, j}] = letter
			}
		}
	}

	schematic_sum := 0
	gear_ratios := 0
	for i, line := range mtx {
		// Reset per line since we don't have wrapping part numbers
		current_num := ""

		// When we finish a digit string, check if it's adjacent to symbols/gears
		for j, letter := range line {
			if unicode.IsDigit(letter) {
				current_num = current_num + string(letter)
			} else {
				if current_num != "" {
					symbol := getAdjacentSymbol(i, j - 1, len(current_num), symbols)
					if symbol != nil {
						val, _ := strconv.Atoi(current_num)
						schematic_sum += val
						updateGearAdjacency(*symbol, gears, val)
					}
				} 
				current_num = ""
			}
		}

		// End of line check
		if current_num != "" {
			symbol := getAdjacentSymbol(i, len(line) - 1, len(current_num), symbols)
			if symbol != nil {
				val, _ := strconv.Atoi(current_num)
				schematic_sum += val
				updateGearAdjacency(*symbol, gears, val)
			}
		}
	}

	for _, gear := range gears {
		if len(gear) == 2 {
			gear_ratios += gear[0] * gear[1]
		}
	}

	return schematic_sum, gear_ratios
}

func updateGearAdjacency(symbol Symbol, gears map[Coordinate][]int, value int) {
	if symbol.char == '*' {
		coordinate := Coordinate{symbol.row, symbol.column}
		if _, exists := gears[coordinate]; !exists {
			gears[coordinate] = []int{}
		}
		gears[coordinate] = append(gears[coordinate], value)
	}
}

func getAdjacentSymbol(row_end int, col_end int, size int, symbols map[Coordinate]rune ) *Symbol {

	for row := row_end - 1; row <= row_end + 1; row++ {
		for col := col_end - size; col <= col_end + 1; col++ {
			value, exists := symbols[Coordinate{row, col}]
			if exists {
				return &Symbol{row: row, column: col, char: value}
			}
		}
	}
	return nil
}