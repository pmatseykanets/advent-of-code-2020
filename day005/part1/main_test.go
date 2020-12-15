package main

import "testing"

func TestSeatID(t *testing.T) {
	tests := []struct {
		pass string
		want int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, tt := range tests {
		t.Run(tt.pass, func(t *testing.T) {
			t.Parallel()

			if got := seatID(tt.pass); tt.want != got {
				t.Errorf("Expected seat ID %d got %d", tt.want, got)
			}
		})
	}
}
