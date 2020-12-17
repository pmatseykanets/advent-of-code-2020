package main

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	input := `
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

	got, err := findInvalid(strings.NewReader(input), 5)
	if err != nil {
		t.Fatal(err)
	}

	if want := 127; want != got {
		t.Errorf("Expected number %d got %d", want, got)
	}

}
