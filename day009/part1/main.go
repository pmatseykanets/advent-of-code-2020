package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

	number, err := findInvalid(f, 25)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d is the invalid number\n", number)
}

type slice struct {
	values []int
	size   int
}

func newSlice(size int) *slice {
	return &slice{size: size, values: make([]int, size)}
}

func (s *slice) push(value int) {
	s.values = append([]int{value}, s.values[:s.size-1]...)
}

func (s *slice) hasSum(sum int) bool {
	for i := 0; i < s.size; i++ {
		for j := 0; j < s.size; j++ {
			if s.values[i]+s.values[j] == sum {
				return true
			}
		}
	}

	return false
}

func findInvalid(input io.Reader, size int) (int, error) {
	var (
		line, values, number int    // avoid allocations.
		text                 string // avoid allocations.
		err                  error  // avoid allocations.
		slice                = newSlice(size)
		scanner              = bufio.NewScanner(input)
	)
	for scanner.Scan() {
		line++
		text = scanner.Text()

		if text == "" {
			continue
		}
		values++

		number, err = strconv.Atoi(text)
		if err != nil {
			return 0, err
		}

		if values > size {
			if !slice.hasSum(number) {
				break
			}
		}

		slice.push(number)
	}
	if err = scanner.Err(); err != nil {
		return 0, err
	}

	return number, nil
}
