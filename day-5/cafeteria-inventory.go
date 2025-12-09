package day5

import (
	_ "embed"
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type IDRange struct {
	min int
	max int
}

func (r IDRange) String() string {
	return fmt.Sprintf("[%d, %d]", r.min, r.max)
}

func SolveCafeteriaInventory(part int) (int, error) {
	if !slices.Contains([]int{1}, part) {
		return 0, fmt.Errorf("called with invalid part")
	}

	slog.Info("Computing number of fresh ingredients...")

	idRanges, inventory, err := readInput()
	if err != nil {
		return 0, err
	}

	freshItemsCount := computeFreshItemsCount(idRanges, inventory)

	return freshItemsCount, nil
}

func readInput() ([]IDRange, []int, error) {
	input = strings.TrimRight(input, "\n")
	// handle CRLF
	input = strings.Replace(input, "\r", "", -1)
	if len(input) == 0 {
		return nil, nil, fmt.Errorf("empty input.txt file")
	}

	splittedInput := strings.Split(input, "\n\n")
	if len(splittedInput) != 2 {
		return nil, nil, fmt.Errorf("invalid input.txt file, splitted in %d insteand of 2", len(splittedInput))
	}
	freshIDRanges, err := parseIDRanges(splittedInput[0])
	if err != nil {
		return nil, nil, err
	}

	inventory, err := parseInventory(splittedInput[1])
	if err != nil {
		return nil, nil, err
	}

	return freshIDRanges, inventory, nil
}

func parseIDRanges(rangesAsStr string) ([]IDRange, error) {
	idRanges := []IDRange{}
	splittedInput := strings.Split(rangesAsStr, "\n")
	for _, rangeAsStr := range splittedInput {
		idRange, err := parseIDRange(rangeAsStr)
		if err != nil {
			return nil, err
		}
		idRanges = append(idRanges, idRange)
	}
	return idRanges, nil
}

func parseIDRange(rangeAsStr string) (IDRange, error) {
	thresholds := strings.Split(rangeAsStr, "-")
	if len(thresholds) != 2 {
		return IDRange{}, fmt.Errorf("invalid fresh ID range")
	}
	min, err := strconv.Atoi(thresholds[0])
	if err != nil {
		return IDRange{}, err
	}
	max, err := strconv.Atoi(thresholds[1])
	if err != nil {
		return IDRange{}, err
	}
	return IDRange{min, max}, nil
}

func parseInventory(inventoryAsStr string) ([]int, error) {
	inventory := []int{}
	for _, idAsStr := range strings.Split(inventoryAsStr, "\n") {
		id, err := strconv.Atoi(idAsStr)
		if err != nil {
			return nil, err
		}
		inventory = append(inventory, id)
	}
	return inventory, nil
}

func computeFreshItemsCount(idRanges []IDRange, itemList []int) int {
	freshItemsCount := 0
	for _, itemId := range itemList {
		isFresh := false
		for _, idRange := range idRanges {
			if isItemFresh(idRange, itemId) {
				isFresh = true
				break
			}
		}
		if isFresh {
			freshItemsCount++
		}
	}
	return freshItemsCount
}

func isItemFresh(idRange IDRange, itemId int) bool {
	return itemId >= idRange.min && itemId <= idRange.max
}
