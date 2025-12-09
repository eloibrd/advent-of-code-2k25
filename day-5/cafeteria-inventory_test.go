package day5

import (
	"reflect"
	"testing"
)

func TestParseIDRange(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  IDRange
		shouldErr bool
	}{
		{
			name:      "should compute range",
			input:     "14-37",
			expected:  IDRange{14, 37},
			shouldErr: false,
		},
		{
			name:      "should return an error : no delimiter in range input",
			input:     "1345",
			expected:  IDRange{},
			shouldErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parseIDRange(tc.input)
			if err != nil && !tc.shouldErr {
				t.Errorf("parseIDRange(%s) unexpected error: %v", tc.input, err)
				return
			}
			if err == nil && tc.shouldErr {
				t.Errorf("parseIDRange(%s) returns a result; expected an error", tc.input)
				return
			}
			if err != nil && tc.shouldErr {
				return
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("parseIDRange(%s) = %s; want %s", tc.input, result.String(), tc.expected.String())
			}
		})
	}
}

func TestParseInventory(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  []int
		shouldErr bool
	}{
		{
			name:      "should parse single line inventory",
			input:     "14",
			expected:  []int{14},
			shouldErr: false,
		},
		{
			name:      "should parse multiple lines inventory",
			input:     "14\n37",
			expected:  []int{14, 37},
			shouldErr: false,
		},
		{
			name:      "should return an error : empty inventory",
			input:     "",
			expected:  []int{},
			shouldErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parseInventory(tc.input)
			if err != nil && !tc.shouldErr {
				t.Errorf("parseInventory(%s) unexpected error: %v", tc.input, err)
				return
			}
			if err == nil && tc.shouldErr {
				t.Errorf("parseInventory(%s) returns %v; expected an error", tc.input, result)
				return
			}
			if err != nil && tc.shouldErr {
				return
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("parseInventory(%s) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestComputeFreshItemsCount(t *testing.T) {
	tests := []struct {
		name     string
		idRanges []IDRange
		itemList []int
		expected int
	}{
		{
			name:     "should find in one range",
			idRanges: []IDRange{{1, 5}},
			itemList: []int{2},
			expected: 1,
		},
		{
			name:     "should find in multiple ranges",
			idRanges: []IDRange{{1, 5}, {50, 90}},
			itemList: []int{2, 3, 13, 17, 34, 59, 90},
			expected: 4,
		},
		{
			name:     "should not find in multiple ranges",
			idRanges: []IDRange{{1, 5}, {50, 90}},
			itemList: []int{13, 17, 34},
			expected: 0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := computeFreshItemsCount(tc.idRanges, tc.itemList)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("computeFreshItemsCount(%v, %v) = %v; want %v", tc.idRanges, tc.itemList, result, tc.expected)
			}
		})
	}
}

func TestComputeFreshIDsCount(t *testing.T) {
	tests := []struct {
		name     string
		idRanges []IDRange
		expected int
	}{
		{
			name:     "should compute fresh IDs count - one range",
			idRanges: []IDRange{{1, 5}},
			expected: 5,
		},
		{
			name:     "should compute fresh IDs count - multiple ranges",
			idRanges: []IDRange{{1, 5}, {8, 12}},
			expected: 10,
		},
		{
			name:     "should compute fresh IDs count - multiple ranges overlapping",
			idRanges: []IDRange{{1, 5}, {8, 12}, {9, 13}},
			expected: 11,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := computeFreshIDsCount(tc.idRanges)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("computeFreshIDsCount(%v) = %v; want %v", tc.idRanges, result, tc.expected)
			}
		})
	}
}

func TestMergeOverlappingRanges(t *testing.T) {
	tests := []struct {
		name     string
		idRanges []IDRange
		expected []IDRange
	}{
		{
			name:     "should merge",
			idRanges: []IDRange{{1, 5}, {4, 12}},
			expected: []IDRange{{1, 12}},
		},
		{
			name:     "should not merge - single range",
			idRanges: []IDRange{{1, 5}},
			expected: []IDRange{{1, 5}},
		},
		{
			name:     "should not merge - multiple ranges not overlapping",
			idRanges: []IDRange{{1, 5}, {8, 12}},
			expected: []IDRange{{1, 5}, {8, 12}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := mergeOverlappingRanges(tc.idRanges)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("mergeOverlappingRanges(%v) = %v; want %v", tc.idRanges, result, tc.expected)
			}
		})
	}
}

func TestSortRangesAsc(t *testing.T) {
	tests := []struct {
		name     string
		idRanges []IDRange
		expected []IDRange
	}{
		{
			name:     "should sort - by min",
			idRanges: []IDRange{{4, 12}, {1, 5}},
			expected: []IDRange{{1, 5}, {4, 12}},
		},
		{
			name:     "should sort - by max",
			idRanges: []IDRange{{1, 12}, {1, 5}},
			expected: []IDRange{{1, 5}, {1, 12}},
		},
		{
			name:     "should not sort - single range",
			idRanges: []IDRange{{1, 5}},
			expected: []IDRange{{1, 5}},
		},
		{
			name:     "should not sort - multiple ranges already sorted",
			idRanges: []IDRange{{1, 5}, {8, 12}},
			expected: []IDRange{{1, 5}, {8, 12}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := sortRangesAsc(tc.idRanges)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("sortRangesAsc(%v) = %v; want %v", tc.idRanges, result, tc.expected)
			}
		})
	}
}

func TestSumRangesLengths(t *testing.T) {
	tests := []struct {
		name     string
		idRanges []IDRange
		expected int
	}{
		{
			name:     "should sum",
			idRanges: []IDRange{{1, 12}, {50, 59}, {111, 120}},
			expected: 32,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := sumRangesLengths(tc.idRanges)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("sortRangesAsc(%v) = %v; want %v", tc.idRanges, result, tc.expected)
			}
		})
	}
}
