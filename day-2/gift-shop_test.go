package day2

import (
	"fmt"
	"testing"
)

func TestFormatInput(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expectedStrings []string
		expectedError   error
	}{
		{"empty input", "", nil, fmt.Errorf("empty input.txt file")},
		{"single range", "10-20", []string{"10-20"}, nil},
		{"multiple ranges", "5-15,25-35", []string{"5-15", "25-35"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formattedInput, err := formatInput(tt.input)
			if tt.expectedError != nil && err == nil {
				t.Errorf("expected error %v, got nil", tt.expectedError)
				return
			}
			if len(formattedInput) != len(tt.expectedStrings) {
				t.Errorf("expected %d input, got %d", len(tt.expectedStrings), len(formattedInput))
				return
			}
			for i, r := range formattedInput {
				if r != tt.expectedStrings[i] {
					t.Errorf("expected %v, got %v", tt.expectedStrings[i], r)
				}
			}
		})
	}
}

func TestInputsToRanges(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		expectedRanges []Range
		expectedError  error
	}{
		{"valid ranges", []string{"10-20", "30-40"}, []Range{{10, 20}, {30, 40}}, nil},
		{"invalid range format", []string{"10-20", "invalid"}, nil, fmt.Errorf("invalid range invalid")},
		{"non-integer values", []string{"10-20", "30-abc"}, nil, fmt.Errorf("strconv.Atoi: parsing \"abc\": invalid syntax")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ranges, err := inputsToRanges(tt.input)
			if tt.expectedError != nil && err == nil {
				t.Errorf("expected error %v, got nil", tt.expectedError)
				return
			}
			if len(ranges) != len(tt.expectedRanges) {
				t.Errorf("expected %d ranges, got %d", len(tt.expectedRanges), len(ranges))
				return
			}
			for i, r := range ranges {
				if r != tt.expectedRanges[i] {
					t.Errorf("expected %v, got %v", tt.expectedRanges[i], r)
				}
			}
		})
	}
}

func TestParseRange(t *testing.T) {
	tests := []struct {
		name          string
		entry         string
		expectedRange Range
		expectedError error
	}{
		{"valid range", "10-20", Range{10, 20}, nil},
		{"invalid format", "10to20", Range{}, fmt.Errorf("invalid range 10to20")},
		{"non-integer min", "abc-20", Range{}, fmt.Errorf("strconv.Atoi: parsing \"abc\": invalid syntax")},
		{"non-integer max", "10-xyz", Range{}, fmt.Errorf("strconv.Atoi: parsing \"xyz\": invalid syntax")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := parseRange(tt.entry)
			if tt.expectedError != nil && err == nil {
				t.Errorf("expected error %v, got nil", tt.expectedError)
				return
			}
			if r != tt.expectedRange {
				t.Errorf("expected %v, got %v", tt.expectedRange, r)
			}
		})
	}
}

func TestSearchInvalidIDsInRangePart1(t *testing.T) {
	tests := []struct {
		name        string
		r           Range
		expectedIDs []int
	}{
		{"no invalid IDs", Range{10, 10}, []int{}},
		{"one invalid ID", Range{95, 115}, []int{99}},
		{"multiple invalid IDs", Range{11, 22}, []int{11, 22}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ids := searchInvalidIDsInRange(tt.r, isIDInvalidPart1)
			if len(ids) != len(tt.expectedIDs) {
				t.Errorf("expected %d IDs, got %d", len(tt.expectedIDs), len(ids))
				return
			}
			for i, id := range ids {
				if id != tt.expectedIDs[i] {
					t.Errorf("expected ID %d, got %d", tt.expectedIDs[i], id)
				}
			}
		})
	}
}

func TestIsIDInvalidPart1(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		expectedResult bool
	}{
		{"valid ID even digits", 50, false},
		{"valid ID odd digits", 54546, false},
		{"invalid ID", 223223, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isIDInvalidPart1(tt.id)
			if result != tt.expectedResult {
				t.Errorf("expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}

func TestComputeSum(t *testing.T) {
	tests := []struct {
		name        string
		ids         []int
		expectedSum int
	}{
		{"empty list", []int{}, 0},
		{"single ID", []int{5}, 5},
		{"multiple IDs", []int{1, 2, 3, 4, 5}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := computeSum(tt.ids)
			if sum != tt.expectedSum {
				t.Errorf("expected sum %d, got %d", tt.expectedSum, sum)
			}
		})
	}
}
