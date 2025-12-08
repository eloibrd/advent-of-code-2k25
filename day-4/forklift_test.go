package day4

import (
	"reflect"
	"testing"
)

var testTable = [][]string{
	[]string{".", ".", "@", "@", ".", "@", "@", "@", "@", "."},
	[]string{"@", "@", "@", ".", "@", ".", "@", ".", "@", "@"},
	[]string{"@", "@", "@", "@", "@", ".", "@", ".", "@", "@"},
	[]string{"@", ".", "@", "@", "@", "@", ".", ".", "@", "."},
	[]string{"@", "@", ".", "@", "@", "@", "@", ".", "@", "@"},
	[]string{".", "@", "@", "@", "@", "@", "@", "@", ".", "@"},
	[]string{".", "@", ".", "@", ".", "@", ".", "@", "@", "@"},
	[]string{"@", ".", "@", "@", "@", ".", "@", "@", "@", "@"},
	[]string{".", "@", "@", "@", "@", "@", "@", "@", "@", "."},
	[]string{"@", ".", "@", ".", "@", "@", "@", ".", "@", "."},
}

func TestIsRollAccessible(t *testing.T) {
	tests := []struct {
		name     string
		table    [][]string
		i        int
		j        int
		expected bool
	}{
		{
			name:     "should be accessible",
			table:    testTable,
			i:        2,
			j:        6,
			expected: true,
		},
		{
			name:     "should be accessible on left border",
			table:    testTable,
			i:        1,
			j:        0,
			expected: true,
		},
		{
			name:     "should be accessible on right border",
			table:    testTable,
			i:        4,
			j:        9,
			expected: true,
		},
		{
			name:     "should be accessible on top border",
			table:    testTable,
			i:        0,
			j:        3,
			expected: true,
		},
		{
			name:     "should be accessible on bottom border",
			table:    testTable,
			i:        9,
			j:        2,
			expected: true,
		},
		{
			name:     "should not be accessible",
			table:    testTable,
			i:        3,
			j:        4,
			expected: false,
		},
		{
			name:     "should not be accessible on top border",
			table:    testTable,
			i:        0,
			j:        7,
			expected: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isRollAccessible(Position{tc.i, tc.j}, tc.table)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("isRollAccessible(%d, %d, table...) = %t; want %t", tc.i, tc.j, result, tc.expected)
			}
		})
	}
}
