package day2

import (
	_ "embed"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Range struct {
	min int
	max int
}

func SolveGiftShop() {
	formattedInput, err := formatInput(input)
	if err != nil {
		panic(err)
	}
	ranges, err := inputsToRanges(formattedInput)
	if err != nil {
		panic(err)
	}
	slog.Info("Number of ranges", "count", len(ranges))
}

func formatInput(rawInput string) ([]string, error) {
	trimedRawInput := strings.TrimRight(rawInput, "\n")
	if len(trimedRawInput) == 0 {
		return nil, fmt.Errorf("empty input.txt file")
	}
	inputList := strings.Split(trimedRawInput, ",")
	return inputList, nil
}

func inputsToRanges(input []string) ([]Range, error) {
	result := []Range{}
	for _, entry := range input {
		rangeEntry, err := parseRange(entry)
		if err != nil {
			return nil, err
		}
		result = append(result, rangeEntry)
	}
	return result, nil
}

func parseRange(entry string) (Range, error) {
	values := strings.Split(entry, "-")
	if len(values) != 2 {
		return Range{}, fmt.Errorf("invalid range %s", entry)
	}
	min, err := strconv.Atoi(values[0])
	if err != nil {
		return Range{}, err
	}
	max, err := strconv.Atoi(values[1])
	if err != nil {
		return Range{}, err
	}
	return Range{min: min, max: max}, nil
}
