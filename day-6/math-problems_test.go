package day6

import (
	"reflect"
	"testing"
)

func TestProblemSolve(t *testing.T) {
	tests := []struct {
		name      string
		problem   problem
		expected  int
		shouldErr bool
	}{
		{
			name:      "should solve addition",
			problem:   problem{operation: "+", operands: []int{34, 46, 10}},
			expected:  90,
			shouldErr: false,
		},
		{
			name:      "should solve multiplication",
			problem:   problem{operation: "*", operands: []int{34, 46, 1}},
			expected:  1564,
			shouldErr: false,
		},
		{
			name:      "should raise an error : no operand",
			problem:   problem{operation: "+", operands: []int{}},
			expected:  0,
			shouldErr: true,
		},
		{
			name:      "should raise an error : unique operand",
			problem:   problem{operation: "*", operands: []int{34}},
			expected:  0,
			shouldErr: true,
		},
		{
			name:      "should raise an error : invalid operation",
			problem:   problem{operation: "-", operands: []int{34, 46}},
			expected:  0,
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.problem.solve()
			if err != nil && !tc.shouldErr {
				t.Errorf("p.solve(%v) unexpected error: %v", tc.problem, err)
				return
			}
			if err == nil && tc.shouldErr {
				t.Errorf("p.solve(%v) returns %v; expected an error", tc.problem, result)
				return
			}
			if err != nil && tc.shouldErr {
				return
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("p.solve(%v) = %v; want %v", tc.problem, result, tc.expected)
			}
		})
	}
}
