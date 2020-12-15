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

	sum, err := sumUnanimous(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d is the sum of counts\n", sum)
}

type group struct {
	people int          // Number of people in the group.
	q      map[byte]int // Number of people answered positive per question.
}

func newGroup() group {
	return group{q: map[byte]int{}}
}

func (g *group) add(answers string) {
	for _, b := range []byte(answers) {
		g.q[b]++
	}
	g.people++
}

func (g *group) unanimous() int {
	var total int
	for _, n := range g.q {
		if n == g.people {
			total++
		}
	}
	return total
}

func (g *group) isZero() bool {
	return len(g.q) == 0 && g.people == 0
}

func sumUnanimous(input io.Reader) (int, error) {
	var (
		line  string // avoid allocations.
		total int
		group group = newGroup()
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line = scanner.Text()

		if line == "" {
			if !group.isZero() {
				total += group.unanimous()
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
	}
	total += group.unanimous()

	return total, nil
}
