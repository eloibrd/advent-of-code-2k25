package day1

import (
	_ "embed"
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"
)

type Direction int

const (
	Left Direction = iota
	Right
)

//go:embed input.txt
var input string

func SolvePassword(part int) {
	if !slices.Contains([]int{1, 2}, part) {
		panic("Called with invalid part")
	}
	position := 50
	password := 0

	steps := readInput(input)

	for _, step := range steps {
		var increment int
		position, increment = handleOneInput(step, position)
		if part == 1 && position == 0 {
			password += 1
		} else if part == 2 {
			password += increment
		}
	}

	slog.Info(fmt.Sprintf("DAY 1 | Found the password for part %d : ", part), slog.String("value", strconv.Itoa(password)))
}

func readInput(pathToFile string) []string {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
	inputList := strings.Split(input, "\n")
	return inputList
}

func handleOneInput(step string, currentPosition int) (int, int) {
	if len(step) == 0 {
		return currentPosition, 0
	}
	direction, shift := parseStep(step)
	switch direction {
	case Left:
		newPosition := currentPosition - shift
		increment := 0
		startedOnZero := currentPosition == 0
		for newPosition < 0 {
			newPosition += 100
			increment++
		}
		if startedOnZero {
			increment--
		}
		if newPosition == 0 {
			increment++
		}
		return newPosition, increment
	case Right:
		newPosition := currentPosition + shift
		increment := newPosition / 100
		newPosition %= 100
		return newPosition, increment
	default:
		panic("invalid direction")
	}
}

func parseStep(step string) (Direction, int) {
	// handle CRLF, \n is already handled in readInput by Split
	value := strings.TrimRight(step[1:], "\r")
	shift, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	switch string(step[0]) {
	case "L":
		return Left, shift
	case "R":
		return Right, shift
	default:
		panic("invalid step")
	}
}
