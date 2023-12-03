package engine_parts

import (
	"fmt"
	"strings"
	"unicode"
)

func getSumOfEngineParts(input string) int {
	parts := extractParts(input)
	sum := 0
	for _, v := range parts {
		sum += v
	}
	return sum
}

func extractParts(input string) []int {
	parts := make([]int, 0)
	matrix := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]int, 0)
		for _, c := range line {
			row = append(row, int(c))
		}
		matrix = append(matrix, row)
	}
	lineLength := len(matrix[0])
	current := 0
	startColumn := 0
	endColumn := 0
	for row, line := range matrix {
		for col, c := range line {
			if unicode.IsDigit(rune(c)) { // is digits
				current = current*10 + (c - 48)
				if startColumn == 0 {
					startColumn = col
				}
			}
			if !unicode.IsDigit(rune(c)) || col == lineLength-1 {
				endColumn = col
				if current > 0 {
					if isAdjacent(matrix, row, startColumn, endColumn, lineLength) {
						parts = append(parts, current)
					}
				}
				startColumn = 0
				current = 0
			}
		}
	}
	return parts
}

func isAdjacent(matrix [][]int, row, startColumn, endColumn, lineLength int) bool {
	for i := -1; i <= 1; i++ {
		for j := startColumn - 1; j <= endColumn; j++ {
			c := getCharacterAt(matrix, row+i, j)
			if !unicode.IsDigit(c) && c != '.' {
				return true
			}
		}
	}
	return false
}

func getCharacterAt(matrix [][]int, i int, j int) rune {
	if i < 0 {
		return '.'
	}
	if i > len(matrix)-1 {
		return '.'
	}
	if j < 0 {
		return '.'
	}
	return rune(matrix[i][j])
}

type Gear struct {
	id    string
	left  int
	right int
}

func (g Gear) getRatio() int {
	return g.left * g.right
}

func getSumOfGears(input string) int {
	gears := extractGears(input)
	sum := 0
	for _, g := range gears {
		sum += g.getRatio()
	}
	return sum
}

func extractGears(input string) map[string]Gear {
	gears := make(map[string]Gear, 0)
	matrix := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]int, 0)
		for _, c := range line {
			row = append(row, int(c))
		}
		matrix = append(matrix, row)
	}
	lineLength := len(matrix[0])
	current := 0
	startColumn := 0
	endColumn := 0
	for row, line := range matrix {
		for col, c := range line {
			if unicode.IsDigit(rune(c)) { // is digits
				current = current*10 + (c - 48)
				if startColumn == 0 {
					startColumn = col
				}
			}
			if !unicode.IsDigit(rune(c)) || col == lineLength-1 {
				endColumn = col
				if current > 0 {
					ok, g := isAdjacentToGear(matrix, row, startColumn, endColumn, lineLength)
					if ok {
						existing, ok := gears[g.id]
						if ok {
							existing.right = current
							gears[g.id] = existing
						} else {
							g.left = current
							gears[g.id] = g
						}
					}
				}
				startColumn = 0
				current = 0
			}
		}
	}
	return gears
}

func isAdjacentToGear(matrix [][]int, row, startColumn, endColumn, lineLength int) (bool, Gear) {
	for i := -1; i <= 1; i++ {
		for j := startColumn - 1; j <= endColumn; j++ {
			c := getCharacterAt(matrix, row+i, j)
			if c == '*' {
				return true, Gear{id: fmt.Sprintf("%d-%d", row+i, j)}
			}
		}
	}
	return false, Gear{}
}
