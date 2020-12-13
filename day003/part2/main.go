package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type slope struct{ x, y int }

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

	trees, err := countTrees(f, []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d trees\n", trees)
}

func countTrees(input io.Reader, slopes []slope) (int, error) {
	var (
		x, y  int    // avoid allocations.
		line  string // avoid allocations.
		paths = make([]struct{ y, trees int }, len(slopes))
	)

	// Do a single pass over input and track each slope's path independently.
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		// Skip empty lines.
		if scanner.Text() == "" {
			continue
		}

		line = scanner.Text()
		for i := range slopes {
			// Skip lines depending on the slope's vertical offset y.
			if y%slopes[i].y != 0 {
				continue
			}

			// Calculate the horizontal offset.
			x = paths[i].y * slopes[i].x % len(line)
			if line[x:x+1] == "#" {
				paths[i].trees++
			}

			// Increment path's vertical coordinate.
			paths[i].y++
		}

		// Increment overall vertical coordinate (line numer).
		y++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	trees := 0
	for i := range paths {
		if i == 0 {
			trees = paths[i].trees
			continue
		}

		trees *= paths[i].trees
	}

	return trees, nil
}
