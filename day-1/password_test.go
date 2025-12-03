package day1

import "testing"

func TestHandleOneInput(t *testing.T) {
	tests := []struct {
		name            string
		step            string
		currentPosition int
		expectedPos     int
	}{
		{"move left", "L10", 50, 40},
		{"move right", "R20", 50, 70},
		{"move left bellow 0", "L110", 50, 40},
		{"move right above 99", "R120", 50, 70},
		{"empty step", "", 50, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := handleOneInput(tt.step, tt.currentPosition)
			if pos != tt.expectedPos {
				t.Errorf("expected position %d, got %d", tt.expectedPos, pos)
			}
		})
	}
}

func TestParseStep(t *testing.T) {
	tests := []struct {
		name          string
		step          string
		expectedShift int
	}{
		{"parse left", "L10", -10},
		{"parse right", "R20", 20},
		{"parse left with CRLF", "L30\r", -30},
		{"parse right with CRLF", "R40\r", 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shift := parseStep(tt.step)
			if shift != tt.expectedShift {
				t.Errorf("expected shift %d, got %d", tt.expectedShift, shift)
			}
		})
	}
}
