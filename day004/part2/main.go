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

var (
	pidRegexp       = regexp.MustCompile("^[0-9]{9}$")
	eyeColorRegexp  = regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	hairColorRegexp = regexp.MustCompile("^#[0-9a-f]{6}$")
	heightRegexp    = regexp.MustCompile("^[0-9]{2,3}(cm|in)$")
)

func isValidYear(year string, min, max int) bool {
	if year == "" {
		return false
	}
	v, _ := strconv.Atoi(year)
	return v >= min && v <= max
}

func isValidHeight(height string) bool {
	if height == "" || len(height) < 4 {
		return false
	}

	switch height[len(height)-2:] {
	case "cm":
		v, _ := strconv.Atoi(height[:len(height)-2])
		return v >= 150 && v <= 193
	case "in":
		v, _ := strconv.Atoi(height[:len(height)-2])
		return v >= 59 && v <= 76
	default:
		return false
	}
}

// Validation rules:
// - byr (Birth Year) - four digits; at least 1920 and at most 2002.
// - iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// - eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// - hgt (Height) - a number followed by either cm or in:
//   - If cm, the number must be at least 150 and at most 193.
//   - If in, the number must be at least 59 and at most 76.
// - hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// - ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// - pid (Passport ID) - a nine-digit number, including leading zeroes.
// - cid (Country ID) - ignored, missing or not.
func (p *passport) isValid() bool {
	return isValidYear(p.byr, 1920, 2002) &&
		isValidYear(p.iyr, 2010, 2020) &&
		isValidYear(p.eyr, 2020, 2030) &&
		isValidHeight(p.hgt) &&
		hairColorRegexp.MatchString(p.hcl) &&
		eyeColorRegexp.MatchString(p.ecl) &&
		pidRegexp.MatchString(p.pid)
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
		p           passport // avoid allocations
		kv          []string // avoid allocations
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
	if !p.isZero() {
		if p.isValid() {
			valid++
		}
	}

	return valid, nil
}
