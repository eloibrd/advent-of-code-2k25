package day3

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertToPowerBank(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      PowerBank
		expectedError error
	}{
		{
			name:          "should parse multiple batteries",
			input:         "369",
			expected:      PowerBank{batteries: []int{3, 6, 9}},
			expectedError: nil,
		},
		{
			name:          "should not parse empty batteries",
			input:         "",
			expected:      PowerBank{},
			expectedError: fmt.Errorf("Not a power bank: "),
		},
		{
			name:          "should not parse single battery",
			input:         "6",
			expected:      PowerBank{},
			expectedError: fmt.Errorf("Not a power bank: 6"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := convertToPowerBank(tt.input)
			if (err != nil && tt.expectedError == nil) || (err == nil && tt.expectedError != nil) {
				t.Errorf("convertToPowerBank(%q) unexpected error: %v; want %v", tt.input, err, tt.expectedError)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("convertToPowerBank(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMaxJoltagePart1(t *testing.T) {
	tests := []struct {
		name      string
		input     PowerBank
		expected  int
		shouldErr bool
	}{
		{
			name:      "should find max joltage",
			input:     PowerBank{batteries: []int{1, 3, 2, 5, 4}},
			expected:  54,
			shouldErr: false,
		},
		{
			name:      "should find max joltage with duplicates",
			input:     PowerBank{batteries: []int{6, 7, 5, 3, 7, 4, 2, 5}},
			expected:  77,
			shouldErr: false,
		},
		{
			name:      "should find max joltage with highest at the end",
			input:     PowerBank{batteries: []int{6, 7, 5, 3, 7, 4, 2, 9}},
			expected:  79,
			shouldErr: false,
		},
		{
			name:      "should not find max joltage and return error",
			input:     PowerBank{batteries: []int{6}},
			expected:  0,
			shouldErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.maxJoltage(2)
			if err != nil && !tt.shouldErr {
				t.Errorf("%v.MaxJoltage(2) unexpected error: %v", tt.input, err)
				return
			}
			if err != nil && tt.shouldErr {
				return
			}
			if result != tt.expected {
				t.Errorf("%v.MaxJoltage(2) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMaxJoltagePart2(t *testing.T) {
	tests := []struct {
		name      string
		input     PowerBank
		expected  int
		shouldErr bool
	}{
		{
			name:      "should find max joltage",
			input:     PowerBank{batteries: []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}},
			expected:  811111111119,
			shouldErr: false,
		},
		{
			name:      "should find max joltage 2",
			input:     PowerBank{batteries: []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}},
			expected:  434234234278,
			shouldErr: false,
		},
		{
			name:      "should not find max joltage and return error",
			input:     PowerBank{batteries: []int{1, 3, 2, 5, 4}},
			expected:  0,
			shouldErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.maxJoltage(12)
			if err != nil && !tt.shouldErr {
				t.Errorf("%v.MaxJoltage(12) unexpected error: %v", tt.input, err)
				return
			}
			if err != nil && tt.shouldErr {
				return
			}
			if result != tt.expected {
				t.Errorf("%v.MaxJoltage(12) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFindMax(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		start         int
		end           int
		expected      int
		expectedIndex int
	}{
		{
			name:          "should find max",
			input:         []int{1, 3, 2, 5, 4},
			start:         0,
			end:           5,
			expected:      5,
			expectedIndex: 3,
		},
		{
			name:          "should find first max when duplicatas",
			input:         []int{6, 7, 5, 3, 7, 4, 2, 5},
			start:         0,
			end:           8,
			expected:      7,
			expectedIndex: 1,
		},
		{
			name:          "should find max in a subarray",
			input:         []int{9, 7, 5, 3, 7, 4, 8, 5},
			start:         2,
			end:           5,
			expected:      7,
			expectedIndex: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, index := findMax(tt.input, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("findMax(%v) = %d; want %d", tt.input, result, tt.expected)
			}
			if index != tt.expectedIndex {
				t.Errorf("findMax(%v) index = %d; want %d", tt.input, index, tt.expectedIndex)
			}
		})
	}
}
