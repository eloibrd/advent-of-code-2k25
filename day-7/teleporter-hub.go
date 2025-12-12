package day7

import (
	_ "embed"
	"fmt"
	"log/slog"
	"slices"
	"strings"
)

const START_CHAR string = "S"
const BEAM_CHAR string = "|"
const SPLITTER_CHAR string = "^"

//go:embed input.txt
var input string

func SolveTeleporterHub(part int) (int, error) {
	if !slices.Contains([]int{1}, part) {
		return 0, fmt.Errorf("called with invalid part")
	}

	slog.Info("Computing tachyon manifold solution...")

	grid := readInput()

	result := computeBeamSplitCount(grid)

	return result, nil
}

func readInput() [][]string {
	input = strings.TrimRight(input, "\n")
	// handle CRLF
	input = strings.Replace(input, "\r", "", -1)

	grid := [][]string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		grid = append(grid, strings.Split(strings.TrimSpace(line), ""))
	}
	return grid
}

func computeBeamSplitCount(grid [][]string) int {
	count := 0
	for row := 1; row < len(grid); row++ {
		for col := range len(grid[0]) {
			if grid[row-1][col] == START_CHAR {
				grid[row][col] = BEAM_CHAR
				continue
			}
			if grid[row-1][col] == BEAM_CHAR {
				if grid[row][col] == SPLITTER_CHAR {
					grid[row][col+1] = BEAM_CHAR
					grid[row][col-1] = BEAM_CHAR
					count++
					continue
				}
				grid[row][col] = BEAM_CHAR
			}
		}
	}
	return count
}
