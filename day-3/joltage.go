package day3

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

type PowerBank struct {
	batteries []int
}

func SolveJoltage(part int) (int, error) {
	if !slices.Contains([]int{1, 2}, part) {
		return 0, fmt.Errorf("called with invalid part")
	}

	slog.Info("Computing max Joltage output...")

	var numberOfBatteries int
	if part == 1 {
		numberOfBatteries = 2
	} else {
		numberOfBatteries = 12
	}

	powerBanks, err := readInput(input)
	if err != nil {
		return 0, err
	}

	totalJoltageOutput := 0
	for _, powerBank := range powerBanks {
		joltage, err := powerBank.maxJoltage(numberOfBatteries)
		if err != nil {
			return 0, err
		}
		totalJoltageOutput += joltage
	}

	return totalJoltageOutput, nil
}

func readInput(input string) ([]PowerBank, error) {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		return nil, fmt.Errorf("empty input.txt file")
	}
	inputList := strings.Split(input, "\n")
	powerBanks := []PowerBank{}
	for _, input := range inputList {
		// handle CRLF
		trimedInput := strings.TrimRight(input, "\r")
		powerBank, err := convertToPowerBank(trimedInput)
		if err != nil {
			return nil, err
		}
		powerBanks = append(powerBanks, powerBank)
	}
	return powerBanks, nil
}

func convertToPowerBank(input string) (PowerBank, error) {
	powerBank := PowerBank{batteries: []int{}}
	if len(input) < 2 {
		return PowerBank{}, fmt.Errorf("not a power bank: %q", input)
	}
	for _, char := range input {
		value, err := strconv.Atoi(string(char))
		if err != nil {
			return PowerBank{}, err
		}
		powerBank.batteries = append(powerBank.batteries, value)
	}
	return powerBank, nil
}

func (p PowerBank) maxJoltage(numberOfBatteries int) (int, error) {
	if len(p.batteries) < numberOfBatteries {
		return 0, fmt.Errorf("not enough batteries to compute max joltage")
	}
	batteriesToTurnOn := []int{}
	currentMaxIndex := 0
	for i := numberOfBatteries - 1; i >= 0; i-- {
		max, maxIndex := findMax(p.batteries, currentMaxIndex, len(p.batteries)-i)
		currentMaxIndex = maxIndex + 1
		batteriesToTurnOn = append(batteriesToTurnOn, max)
	}
	if len(batteriesToTurnOn) != numberOfBatteries {
		return 0, fmt.Errorf("an error occured while computing joltage")
	}
	joltage := 0
	pow := numberOfBatteries - 1
	for _, battery := range batteriesToTurnOn {
		joltage += battery * int(math.Pow(float64(10), float64(pow)))
		pow--
	}
	return joltage, nil
}

func findMax(array []int, start int, end int) (int, int) {
	max := 0
	maxIndex := start
	for i := start; i < end; i++ {
		if array[i] > max {
			max = array[i]
			maxIndex = i
		}
	}
	return max, maxIndex
}
