package day3

import (
	_ "embed"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type PowerBank struct {
	batteries []int
}

func SolveJoltage() {
	powerBanks, err := readInput(input)
	if err != nil {
		panic(err)
	}
	totalJoltageOutput := 0
	for _, powerBank := range powerBanks {
		totalJoltageOutput += powerBank.maxJoltage()
	}
	slog.Info("Max Joltage output for input found : ", slog.Int("value", totalJoltageOutput))
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

func (p PowerBank) maxJoltage() int {
	max, maxIndex := findMax(p.batteries, 0, len(p.batteries)-1)
	max2, _ := findMax(p.batteries, maxIndex+1, len(p.batteries))
	return max*10 + max2
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
