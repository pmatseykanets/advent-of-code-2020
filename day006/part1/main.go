package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

	sum, err := sumCounts(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d is the sum of counts\n", sum)
}

type group struct {
	q map[byte]int // Positive answers.
}

func newGroup() group {
	return group{q: map[byte]int{}}
}

func (g *group) add(answers string) {
	for _, b := range []byte(answers) {
		g.q[b]++
	}
}

func (g *group) total() int {
	return len(g.q)
}

func (g *group) isZero() bool {
	return len(g.q) == 0
}

func sumCounts(input io.Reader) (int, error) {
	var (
		line  string // avoid allocations.
		total int
		group = newGroup()
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line = scanner.Text()

		if line == "" {
			if !group.isZero() {
				total += group.total()
				group = newGroup()
			}
			continue
		}
		group.add(line)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	if line != "" {
		group.add(line)
		total += group.total()
	}

	return total, nil
}
