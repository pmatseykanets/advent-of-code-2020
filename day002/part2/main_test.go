package main

import (
	"fmt"
	"testing"
)

func TestPolicyValidate(t *testing.T) {
	tests := []struct {
		policy   policy
		password string
		want     bool
	}{
		{policy{min: 1, max: 2, char: "o"}, "foo", true},
		{policy{min: 1, max: 2, char: "f"}, "foo", true},
		{policy{min: 2, max: 3, char: "o"}, "foo", false},
		{policy{min: 1, max: 2, char: "r"}, "bar", false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-%d %s: %s", tt.policy.min, tt.policy.max, tt.policy.char, tt.password), func(t *testing.T) {
			t.Parallel()

			if got := tt.policy.validate(tt.password); tt.want != got {
				t.Errorf("Expected %v got %v", tt.want, got)
			}
		})
	}
}
