package day7

import (
	"reflect"
	"testing"
)

var testGrid [][]string = [][]string{
	{".", ".", "S", ".", "."},
	{".", ".", ".", ".", "."},
	{".", ".", "^", ".", "."},
	{".", ".", ".", "^", "."},
	{".", ".", ".", ".", "."},
	{".", "^", ".", "^", "."},
}

func TestComputeBeamSplitCount(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]string
		expected int
	}{
		{
			name:     "computeBeamSplitCount should compute 3",
			grid:     testGrid,
			expected: 3,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := computeBeamSplitCount(tc.grid)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("computeBeamSplitCount(grid...) = %d; expected %d", result, tc.expected)
			}
		})
	}
}
