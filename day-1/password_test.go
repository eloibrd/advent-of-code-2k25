package day1

import "testing"

func TestHandleOneInput(t *testing.T) {
	tests := []struct {
		name            string
		step            string
		currentPosition int
		expectedPos     int
		expectedIncr    int
	}{
		{"move left without increment", "L10", 50, 40, 0},
		{"move right without increment", "R20", 50, 70, 0},
		{"move left with 1 increment", "L60", 50, 90, 1},
		{"move left with 1 increment ending at 0", "L50", 50, 0, 1},
		{"move right with 1 increment", "R60", 50, 10, 1},
		{"move right with 1 increment ending at 0", "R50", 50, 0, 1},
		{"move left with 2 increment", "L160", 50, 90, 2},
		{"move left with 2 increment ending at 0", "150", 50, 0, 2},
		{"move right with 3 increment", "R260", 50, 10, 3},
		{"move right with 3 increment ending at 0", "R250", 50, 0, 3},
		{"empty step", "", 50, 50, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos, incr := handleOneInput(tt.step, tt.currentPosition)
			if pos != tt.expectedPos {
				t.Errorf("expected position %d, got %d", tt.expectedPos, pos)
			}
			if incr != tt.expectedIncr {
				t.Errorf("expected increment %d, got %d", tt.expectedIncr, incr)
			}
		})
	}
}

func TestParseStep(t *testing.T) {
	tests := []struct {
		name              string
		step              string
		expectedShift     int
		expectedDirection Direction
	}{
		{"parse left", "L10", -10, Left},
		{"parse right", "R20", 20, Right},
		{"parse left with CRLF", "L30\r", -30, Left},
		{"parse right with CRLF", "R40\r", 40, Right},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			direction, shift := parseStep(tt.step)
			if shift != tt.expectedShift {
				t.Errorf("expected shift %d, got %d", tt.expectedShift, shift)
			}
			if direction != tt.expectedDirection {
				t.Errorf("expected direction %d, got %d", tt.expectedDirection, direction)
			}
		})
	}
}
