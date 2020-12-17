package main

import (
	"strings"
	"testing"
)

func TestProgramExecute(t *testing.T) {
	input := `
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

	prg, err := parseProgram(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	err = prg.execute()
	if want, got := errInfiniteLoop, err; want != got {
		t.Fatalf("Expected error %v got %v", want, got)
	}

	if want, got := 5, prg.accumulator; want != got {
		t.Errorf("Expected accumulator value %d got %d", want, got)
	}
}
