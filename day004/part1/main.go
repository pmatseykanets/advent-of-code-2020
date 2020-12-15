package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type passport struct {
	byr string // Birth Year e.g. 1931
	iyr string // Issue Year e.g. 2015
	eyr string // Expiration Year e.g. 2029
	hgt string // Height e.g.150cm
	hcl string // Hair Color e.g. z
	ecl string // Eye Color e.g. amb
	pid string // Passport ID e.g. 148714704
	cid string // Country ID e.g. 128
}

func (p *passport) isZero() bool {
	return p.byr == "" && p.iyr == "" && p.eyr == "" && p.hgt == "" && p.hcl == "" && p.ecl == "" && p.pid == "" && p.cid == ""
}

func (p *passport) isValid() bool {
	return p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != ""
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

	valid, err := validatePassports(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d valid passports\n", valid)
}

func validatePassports(input io.Reader) (int, error) {
	var (
		line, field string   // avoid allocations.
		p           passport // avoid allocations.
		kv          []string // avoid allocations.
		valid       int
	)

	// Do a single pass over input and track each slope's path independently.
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line = scanner.Text()

		if line == "" && !p.isZero() {
			if p.isValid() {
				valid++
			}
			p = passport{}
			continue
		}

		for _, field = range strings.Fields(line) {
			kv = strings.Split(field, ":")
			if len(kv) != 2 {
				return 0, fmt.Errorf("invalid field: %s", field)
			}

			switch kv[0] {
			case "byr":
				p.byr = kv[1]
			case "iyr":
				p.iyr = kv[1]
			case "eyr":
				p.eyr = kv[1]
			case "hgt":
				p.hgt = kv[1]
			case "hcl":
				p.hcl = kv[1]
			case "ecl":
				p.ecl = kv[1]
			case "pid":
				p.pid = kv[1]
			case "cid":
				p.cid = kv[1]
			default:
				return 0, fmt.Errorf("unknown field: %s", kv[0])
			}
		}

	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	// The trailing record.
	if !p.isZero() && p.isValid() {
		valid++
	}

	return valid, nil
}
