package day2

import (
	_ "embed"
	"log/slog"
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
	return []string{}, nil
}

func inputsToRanges(input []string) ([]Range, error) {
	result := []Range{}
	return result, nil
}

func parseRange(entry string) (Range, error) {
	return Range{}, nil
}
