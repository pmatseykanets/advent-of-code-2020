package main

import (
	"strings"
	"testing"
)

func TestGroup(t *testing.T) {
	g := newGroup()

	if want, got := true, g.isZero(); want != got {
		t.Fatalf("Expected isZero %v got %v", want, got)
	}
}

func TestSumUnanimous(t *testing.T) {
	input := `
abc

a
b
c

ab
ac

a
a
a
a

b`

	got, err := sumUnanimous(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	if want := 6; want != got {
		t.Errorf("Expected sum of counts %d got %d", want, got)
	}
}
