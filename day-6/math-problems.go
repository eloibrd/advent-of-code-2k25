package day6

import (
	_ "embed"
	"fmt"
	"log/slog"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type problem struct {
	operation string
	operands  []int
}

func (p *problem) solve() (int, error) {
	if len(p.operands) <= 1 {
		return 0, fmt.Errorf("not enough operands in problem")
	}
	result := p.operands[0]
	for i := 1; i < len(p.operands); i++ {
		switch p.operation {
		case "+":
			result += p.operands[i]
		case "*":
			result *= p.operands[i]
		default:
			return 0, fmt.Errorf("operation %s unsupported", p.operation)
		}
	}
	return result, nil
}

func SolveMathProblems(part int) (int, error) {
	if !slices.Contains([]int{1}, part) {
		return 0, fmt.Errorf("called with invalid part")
	}

	slog.Info("Computing math problems total...")

	grid, err := readInput()
	if err != nil {
		return 0, err
	}

	var problems []problem
	if part == 1 {
		problems, err = transformToPart1Problems(grid)
	}
	if err != nil {
		return 0, err
	}

	result, err := computeProblemsSum(problems)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func readInput() ([][]string, error) {
	input = strings.TrimRight(input, "\n")
	// handle CRLF
	input = strings.Replace(input, "\r", "", -1)

	multipleSpacesRegex, err := regexp.Compile(`\s{2,}`)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(input, "\n")
	for i := range lines {
		lines[i] = multipleSpacesRegex.ReplaceAllString(lines[i], " ")
	}

	grid := [][]string{}
	for _, line := range lines {
		grid = append(grid, strings.Split(line, " "))
	}

	return grid, nil
}

func transformToPart1Problems(grid [][]string) ([]problem, error) {
	problems := []problem{}
	for i := range len(grid[0]) {
		prob := problem{}
		for j := range len(grid) {
			if j == len(grid)-1 {
				prob.operation = grid[j][i]
				continue
			}
			num, err := strconv.Atoi(grid[j][i])
			if err != nil {
				return nil, err
			}
			prob.operands = append(prob.operands, num)
		}
		problems = append(problems, prob)
	}
	return problems, nil
}

func computeProblemsSum(problems []problem) (int, error) {
	sum := 0
	for _, prob := range problems {
		result, err := prob.solve()
		if err != nil {
			return 0, err
		}
		sum += result
	}
	return sum, nil
}
