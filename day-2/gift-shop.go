package day2

import (
	_ "embed"
	"fmt"
	"log/slog"
	"math"
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
	invalidIDs := []int{}

	formattedInput, err := formatInput(input)
	if err != nil {
		panic(err)
	}
	ranges, err := inputsToRanges(formattedInput)
	if err != nil {
		panic(err)
	}
	for _, r := range ranges {
		invalidIDsInRange := searchInvalidIDsInRange(r)
		invalidIDs = append(invalidIDs, invalidIDsInRange...)
	}
	sum := computeSum(invalidIDs)
	slog.Info("Sum of invalid IDs", "sum", sum)
}

func formatInput(rawInput string) ([]string, error) {
	trimedRawInput := strings.TrimRight(rawInput, "\n")
	trimedRawInput = strings.TrimRight(trimedRawInput, "\r")
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

func searchInvalidIDsInRange(r Range) []int {
	invalidIDs := []int{}
	for id := r.min; id <= r.max; id++ {
		if isIDInvalid(id) {
			invalidIDs = append(invalidIDs, id)
		}
	}
	return invalidIDs
}

func isIDInvalid(id int) bool {
	// Count digits
	digits := int(math.Log10(float64(id))) + 1

	// Must have an even number of digits
	if digits%2 != 0 {
		return false
	}

	base := int(math.Pow10(digits / 2))

	left := id / base
	right := id % base

	return left == right
}

func computeSum(ids []int) int {
	sum := 0
	for _, id := range ids {
		sum += id
	}
	return sum
}
