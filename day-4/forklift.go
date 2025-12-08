package day4

import (
	_ "embed"
	"fmt"
	"log/slog"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

const ROLL_CHAR string = "@"
const MAX_ADJACENT_ROLLS int = 3

func SolveForklift(part int) (int, error) {
	if !slices.Contains([]int{1}, part) {
		return 0, fmt.Errorf("called with invalid part")
	}

	slog.Info("Computing number of accessible rolls of paper...")

	table, err := readInput()

	if err != nil {
		return 0, err
	}

	result := computeNumberOfAccessibleRolls(table)

	return result, nil
}

func readInput() ([][]string, error) {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		return nil, fmt.Errorf("empty input.txt file")
	}

	table := [][]string{}

	splittedInput := strings.Split(input, "\n")

	for _, line := range splittedInput {
		row := []string{}
		// handle CRLF
		trimedRow := strings.TrimRight(line, "\r")
		for _, c := range trimedRow {
			row = append(row, string(c))
		}
		table = append(table, row)
	}
	return table, nil
}

func computeNumberOfAccessibleRolls(table [][]string) int {
	accessibleRollsCounter := 0
	for i, row := range table {
		for j := range row {
			if table[i][j] == ROLL_CHAR && isRollAccessible(i, j, table) {
				accessibleRollsCounter++
			}
		}
	}
	return accessibleRollsCounter
}

func isRollAccessible(i, j int, table [][]string) bool {
	minY := max(i-1, 0)
	maxY := min(i+1, len(table)-1)
	minX := max(j-1, 0)
	maxX := min(j+1, len(table[0])-1)

	adjacentRollsCount := 0
	for y := minY; y < maxY+1; y++ {
		for x := minX; x < maxX+1; x++ {
			if x == j && y == i {
				continue
			}
			if table[y][x] == ROLL_CHAR {
				adjacentRollsCount++
			}
		}
	}
	return adjacentRollsCount <= MAX_ADJACENT_ROLLS
}
