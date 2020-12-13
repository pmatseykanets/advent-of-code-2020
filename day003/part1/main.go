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

	trees, err := countTrees(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d trees\n", trees)
}

func countTrees(input io.Reader) (int, error) {
	var (
		x, y, trees int    // avoid allocations.
		line        string // avoid allocations.
	)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		line = scanner.Text()
		x = y * 3 % len(line)
		if line[x:x+1] == "#" {
			trees++
		}

		y++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return trees, nil
}
