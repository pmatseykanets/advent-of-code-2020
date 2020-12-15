package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

	ids, err := getSeatIDs(f)
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(ids)

	var id int
	for i := 1; i < len(ids); i++ {
		if ids[i]-ids[i-1] == 2 {
			id = ids[i] - 1
			break
		}
	}

	fmt.Printf("%d is the missing seat ID\n", id)
}

func seatID(pass string) int {
	if len(pass) != 10 {
		return 0
	}

	row, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(pass[:7], "F", "0"), "B", "1"), 2, 64)
	col, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(pass[7:], "L", "0"), "R", "1"), 2, 64)

	return int(row*8 + col)
}

func getSeatIDs(input io.Reader) ([]int, error) {
	var (
		line string // avoid allocations.
		ids  []int
	)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line = scanner.Text()

		if line == "" {
			continue
		}

		ids = append(ids, seatID(line))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}
