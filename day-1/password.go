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

func SolvePassword() {
	input := "day-1/input.txt"
	position := 50
	password := 0

	steps, err := readInput(input)
	if err != nil {
		panic(err)
	}

	for _, step := range steps {
		position = handleOneInput(step, position)
		if position == 0 {
			password++
		}
	}

	slog.Info("Found the password !", "value", strconv.Itoa(password))
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

func handleOneInput(step string, currentPosition int) int {
	if len(step) == 0 {
		return currentPosition
	}
	direction, shift := parseStep(step)
	switch direction {
	case Left:
		return ((currentPosition-shift)%100 + 100) % 100
	case Right:
		return ((currentPosition+shift)%100 + 100) % 100
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
