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
