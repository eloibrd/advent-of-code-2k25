package day1

import (
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Left Direction = iota
	Right
)

func SolvePasswordPart1() {
	input := "day-1/input.txt"
	position := 50
	password := 0

	steps, err := readInput(input)
	if err != nil {
		panic(err)
	}

	for _, step := range steps {
		position, _ = handleOneInput(step, position)
		if position == 0 {
			password += 1
		}
	}

	slog.Info("DAY 1 | Found the password for part 1 : ", "value", strconv.Itoa(password))
}

func SolvePasswordPart2() {
	input := "day-1/input.txt"
	position := 50
	password := 0

	steps, err := readInput(input)
	if err != nil {
		panic(err)
	}

	for _, step := range steps {
		var increment int
		position, increment = handleOneInput(step, position)
		password += increment
	}

	slog.Info("DAY 1 | Found the password for part 2 : ", "value", strconv.Itoa(password))
}

func readInput(pathToFile string) ([]string, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	inputList := strings.Split(string(content), "\n")
	return inputList, nil
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
