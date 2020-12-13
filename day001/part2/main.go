package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fname := "input.txt"
	if len(os.Args) > 1 && os.Args[1] != "" {
		fname = os.Args[1]
	}

	numbers, err := readFromFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	a, b, c, d := findThree(numbers)
	fmt.Printf("%d * %d * %d = %d\n", a, b, c, d)
}

func findThree(numbers []int) (int, int, int, int) {
	for _, i := range numbers {
		for _, j := range numbers {
			for _, k := range numbers {
				if i+j+k == 2020 {
					return i, j, k, i * j * k
				}
			}
		}
	}
	return 0, 0, 0, 0
}

func readFromFile(name string) ([]int, error) {
	var numbers []int

	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}
