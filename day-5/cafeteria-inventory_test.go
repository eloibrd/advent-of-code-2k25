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
