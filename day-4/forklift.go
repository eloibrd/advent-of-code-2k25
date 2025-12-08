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

type Position struct {
	i int
	j int
}

const ROLL_CHAR string = "@"
const MAX_ADJACENT_ROLLS int = 3

func SolveForklift(part int) (int, error) {
	if !slices.Contains([]int{1, 2}, part) {
		return 0, fmt.Errorf("called with invalid part")
	}

	slog.Info("Computing number of accessible rolls of paper...")

	table, err := readInput()

	if err != nil {
		return 0, err
	}

	if part == 1 {
		accessibleRolls := computeAccessibleRolls(table)
		return len(accessibleRolls), nil
	}

	totalCount := 0

	return totalCount, nil
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

func computeAccessibleRolls(table [][]string) []Position {
	accessiblePositions := []Position{}
	for i, row := range table {
		for j := range row {
			if table[i][j] == ROLL_CHAR && isRollAccessible(Position{i, j}, table) {
				accessiblePositions = append(accessiblePositions, Position{i, j})
			}
		}
	}
	return accessiblePositions
}

func isRollAccessible(pos Position, table [][]string) bool {
	minY := max(pos.i-1, 0)
	maxY := min(pos.i+1, len(table)-1)
	minX := max(pos.j-1, 0)
	maxX := min(pos.j+1, len(table[0])-1)

	adjacentRollsCount := 0
	for y := minY; y < maxY+1; y++ {
		for x := minX; x < maxX+1; x++ {
			if x == pos.j && y == pos.i {
				continue
			}
			if table[y][x] == ROLL_CHAR {
				adjacentRollsCount++
			}
		}
	}
	return adjacentRollsCount <= MAX_ADJACENT_ROLLS
}
