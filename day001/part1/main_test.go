package main

import (
	"fmt"
	"testing"
)

func TestFindTwo(t *testing.T) {
	tests := []struct {
		input   []int
		a, b, c int
	}{
		{[]int{1010, 1010}, 1010, 1010, 1010 * 1010},
		{[]int{1, 1010, 1010, 2, 3}, 1010, 1010, 1010 * 1010},
		{[]int{1, 2, 2019, 1010}, 1, 2019, 2019},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v:%d", tt.input, tt.c), func(t *testing.T) {
			t.Parallel()

			a, b, c := findTwo(tt.input)
			if want, got := tt.a, a; want != got {
				t.Fatalf("Expected a %d got %d", want, got)
			}
			if want, got := tt.b, b; want != got {
				t.Fatalf("Expected b %d got %d", want, got)
			}
			if want, got := tt.c, c; want != got {
				t.Fatalf("Expected c %d got %d", want, got)
			}
		})
	}
}
