package main

import (
	"fmt"
	"testing"
)

func TestFindThree(t *testing.T) {
	tests := []struct {
		input      []int
		a, b, c, d int
	}{
		{[]int{0, 1010, 1010}, 1010, 0, 1010, 0},
		{[]int{0, 1, 1009, 1010, 2, 3}, 1, 1009, 1010, 1009 * 1010},
		{[]int{1, 2, 0, 2019, 1010}, 1, 0, 2019, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v:%d", tt.input, tt.c), func(t *testing.T) {
			t.Parallel()

			a, b, c, d := findThree(tt.input)
			if want, got := tt.a, a; want != got {
				t.Fatalf("Expected a %d got %d", want, got)
			}
			if want, got := tt.b, b; want != got {
				t.Fatalf("Expected b %d got %d", want, got)
			}
			if want, got := tt.c, c; want != got {
				t.Fatalf("Expected c %d got %d", want, got)
			}
			if want, got := tt.d, d; want != got {
				t.Fatalf("Expected d %d got %d", want, got)
			}
		})
	}
}
