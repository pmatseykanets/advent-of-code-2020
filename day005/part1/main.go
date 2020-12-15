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

	maxID, err := findMaxSeatID(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d is the highest seat ID\n", maxID)
}

func seatID(pass string) int {
	if len(pass) != 10 {
		return 0
	}

	row, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(pass[:7], "F", "0"), "B", "1"), 2, 64)
	col, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(pass[7:], "L", "0"), "R", "1"), 2, 64)

	return int(row*8 + col)
}

func findMaxSeatID(input io.Reader) (int, error) {
	var (
		id, max int
		line    string // avoid allocations.
	)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line = scanner.Text()

		if line == "" {
			continue
		}

		if id = seatID(line); id > max {
			max = id
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return max, nil
}
