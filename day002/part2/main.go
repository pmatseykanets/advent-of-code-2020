package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type policy struct {
	min  int
	max  int
	char string
}

func (p *policy) validate(password string) bool {
	length := len(password)
	if length < p.min || length < p.max {
		return false
	}

	occurencies := 0
	if string(password[p.min-1]) == p.char {
		occurencies++
	}
	if string(password[p.max-1]) == p.char {
		occurencies++
	}

	return occurencies == 1
}

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

	var ln, valid int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ln++
		if scanner.Text() == "" {
			continue
		}

		policy, password, err := parseLine(scanner.Text())
		if err != nil {
			log.Fatalf("line %d: %s", ln, err)
		}
		if policy.validate(password) {
			valid++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d valid passwords\n", valid)
}

func parseLine(line string) (policy, string, error) {
	var (
		policy   policy
		password string
	)

	parts := strings.Split(line, ": ")
	if len(parts) != 2 {
		return policy, "", fmt.Errorf("invalid format")
	}
	password = parts[1]

	parts = strings.Split(parts[0], " ")
	if len(parts) != 2 {
		return policy, "", fmt.Errorf("invalid format")
	}
	policy.char = parts[1]

	parts = strings.Split(parts[0], "-")
	if len(parts) != 2 {
		return policy, "", fmt.Errorf("invalid format")
	}
	var err error
	policy.min, err = strconv.Atoi(parts[0])
	if err != nil {
		return policy, "", err
	}
	policy.max, err = strconv.Atoi(parts[1])
	if err != nil {
		return policy, "", err
	}

	return policy, password, nil
}
