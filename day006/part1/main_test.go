package main

import (
	"strings"
	"testing"
)

func TestSumCounts(t *testing.T) {
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

	got, err := sumCounts(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	if want := 11; want != got {
		t.Errorf("Expected sum of counts %d got %d", want, got)
	}
}
