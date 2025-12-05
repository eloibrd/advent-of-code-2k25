package day2

import (
	_ "embed"
	"fmt"
	"log/slog"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Range struct {
	min int
	max int
}

type isIDInvalidFn func(int) bool

func SolveGiftShop(part int) (int, error) {
	if !slices.Contains([]int{1, 2}, part) {
		return 0, fmt.Errorf("called with invalid part")
	}

	slog.Info("Computing invalid IDs...")

	invalidIDs := []int{}

	formattedInput, err := formatInput(input)
	if err != nil {
		return 0, err
	}
	ranges, err := inputsToRanges(formattedInput)
	if err != nil {
		return 0, err
	}
	for _, r := range ranges {
		var validationFn isIDInvalidFn
		if part == 1 {
			validationFn = isIDInvalidPart1
		} else {
			validationFn = isIDInvalidPart2
		}
		invalidIDsInRange := searchInvalidIDsInRange(r, validationFn)
		invalidIDs = append(invalidIDs, invalidIDsInRange...)
	}
	sum := computeSum(invalidIDs)
	return sum, nil
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

func searchInvalidIDsInRange(r Range, isInvalid isIDInvalidFn) []int {
	invalidIDs := []int{}
	for id := r.min; id <= r.max; id++ {
		if isInvalid(id) {
			invalidIDs = append(invalidIDs, id)
		}
	}
	return invalidIDs
}

func isIDInvalidPart1(id int) bool {
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

func isIDInvalidPart2(id int) bool {
	s := strconv.Itoa(id)
	length := len(s)

	// Try every possible pattern length
	// Pattern size must divide total length
	for size := 1; size <= length/2; size++ {
		if length%size != 0 {
			continue
		}

		block := s[:size]
		k := length / size
		ok := true

		// Check all blocks
		for i := 1; i < k; i++ {
			start := i * size
			end := start + size
			if s[start:end] != block {
				ok = false
				break
			}
		}

		if ok {
			return true
		}
	}

	return false
}

func computeSum(ids []int) int {
	sum := 0
	for _, id := range ids {
		sum += id
	}
	return sum
}
