package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fname := "input.txt"
	if len(os.Args) > 1 && os.Args[1] != "" {
		fname = os.Args[1]
	}

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	prg, err := parseProgram(f)
	if err != nil {
		log.Fatal(err)
	}

	err = prg.execute()
	if err != nil {
		fmt.Printf("runtime error: %s\n", err)
	}

	fmt.Printf("%d is in the accumulator \n", prg.accumulator)
}

type program struct {
	accumulator  int
	instructions []instruction
	curr         int
	executed     map[int]int
}

var (
	errInvalidAddress = fmt.Errorf("invalid address")
	errInfiniteLoop   = fmt.Errorf("infinite loop")
)

func (p *program) execute() error {
	last := len(p.instructions) - 1

	for {
		if p.curr < 0 || p.curr > last {
			return errInvalidAddress
		}

		p.executed[p.curr]++
		if p.executed[p.curr] > 1 {
			return errInfiniteLoop
		}

		switch p.instructions[p.curr].op {
		case "acc":
			p.accumulator += p.instructions[p.curr].arg
		case "jmp":
			p.curr += p.instructions[p.curr].arg
			continue
		}
		p.curr++

		if p.curr == last {
			break
		}
	}

	return nil
}

type instruction struct {
	op  string
	arg int
}

func parseProgram(input io.Reader) (*program, error) {
	var (
		line    int // avoid allocations.
		text    string
		parts   []string
		ins     instruction
		err     error
		program = &program{executed: make(map[int]int)}
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line++
		text = scanner.Text()

		if text == "" {
			continue
		}

		parts = strings.Fields(text)
		if len(parts) != 2 {
			return nil, fmt.Errorf("line %d: malformed instruction", line)
		}

		switch parts[0] {
		case "nop", "acc", "jmp":
			ins.op = parts[0]
		default:
			return nil, fmt.Errorf("line %d: unknown operation %s", line, parts[0])
		}

		ins.arg, err = strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("line %d: invalid arg %s", line, parts[1])
		}

		program.instructions = append(program.instructions, ins)
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return program, nil
}
