package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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

	invertedRules, err := parseRules(f)
	if err != nil {
		log.Fatal(err)
	}

	colors := canBeContained("shiny gold", invertedRules)

	fmt.Printf("%d colors\n", len(colors))
}

type rule struct {
	color    string
	contains map[string]int
}

var bagRegexp = regexp.MustCompile(` bag(s?)([.,]?)( ?)`)

func parseRule(s string) (rule, error) {
	if strings.HasSuffix(s, " bags contain no other bags.") {
		return rule{color: strings.TrimSuffix(s, " bags contain no other bags.")}, nil
	}

	var (
		parts []string
		n     int
		err   error
	)

	parts = strings.Split(s, " bags contain ")
	if len(parts) != 2 {
		return rule{}, fmt.Errorf("malformed rule")
	}

	r := rule{
		color:    parts[0],
		contains: make(map[string]int),
	}

	for _, part := range bagRegexp.Split(parts[1], -1) {
		if part == "" || part == "." || part == "," {
			continue
		}

		parts = strings.Split(part, " ")
		if len(parts) < 2 {
			return rule{}, fmt.Errorf("malformed rule")
		}

		n, err = strconv.Atoi(parts[0])
		if err != nil {
			return rule{}, fmt.Errorf("malformed rule")
		}

		r.contains[strings.Join(parts[1:], " ")] = n
	}

	return r, nil
}

func parseRules(input io.Reader) (map[string]map[string]struct{}, error) {
	var (
		line, color string // avoid allocations.
		r           rule   // avoid allocations.
		err         error  // avoid allocations.
		ok          bool   // avoid allocations.
		rules       = make(map[string]map[string]struct{})
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line = scanner.Text()

		if line == "" {
			continue
		}

		r, err = parseRule(line)
		if err != nil {
			return nil, err
		}

		for color = range r.contains {
			if _, ok = rules[color]; !ok {
				rules[color] = make(map[string]struct{})
			}
			rules[color][r.color] = struct{}{}
		}
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return rules, nil
}

func canBeContained(needle string, hay map[string]map[string]struct{}) map[string]struct{} {
	var (
		result = make(map[string]struct{})
		color  string
	)

	for color = range hay[needle] {
		result[color] = struct{}{}
		for color = range canBeContained(color, hay) {
			result[color] = struct{}{}
		}
	}

	return result
}
