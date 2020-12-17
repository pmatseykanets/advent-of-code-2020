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

	numbers, err := loadNumbers(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d is the weakness\n", findWeakness(numbers, findInvalid(numbers, 25)))
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

func loadNumbers(input io.Reader) ([]int, error) {
	var (
		text    string
		number  int
		numbers []int
		err     error
		scanner = bufio.NewScanner(input)
	)
	for scanner.Scan() {
		text = scanner.Text()

		if text == "" {
			continue
		}

		number, err = strconv.Atoi(text)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func findInvalid(numbers []int, size int) int {
	slice := newSlice(size)

	for i := 0; i < len(numbers); i++ {
		if i+1 > size {
			if !slice.hasSum(numbers[i]) {
				return numbers[i]
			}
		}

		slice.push(numbers[i])
	}

	return 0
}

func findWeakness(numbers []int, invalid int) int {
	var i, j, sum int

O:
	for i = 0; i < len(numbers); i++ {
		sum = 0
	I:
		for j = i; j < len(numbers); j++ {
			sum += numbers[j]
			switch {
			case sum < invalid:
				continue
			case sum == invalid:
				break O
			default:
				break I
			}
		}
	}

	var min, max = numbers[i], numbers[i]
	for ; i <= j; i++ {
		if numbers[i] < min {
			min = numbers[i]
		} else if numbers[i] > max {
			max = numbers[i]
		}
	}

	return min + max
}
