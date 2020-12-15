package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPassportIsZero(t *testing.T) {
	tests := []struct {
		p    passport
		want bool
	}{
		{passport{}, true},
		{passport{byr: "1937"}, false},
		{passport{iyr: "2017"}, false},
		{passport{eyr: "2020"}, false},
		{passport{hgt: "183cm"}, false},
		{passport{hcl: "#fffffd"}, false},
		{passport{ecl: "gry"}, false},
		{passport{pid: "860033327"}, false},
		{passport{cid: "147"}, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.p), func(t *testing.T) {
			t.Parallel()

			if got := tt.p.isZero(); tt.want != got {
				t.Errorf("Expected %v got %v", tt.want, got)
			}
		})
	}
}

func TestPassportIsValid(t *testing.T) {
	tests := []struct {
		p    passport
		want bool
	}{
		{passport{byr: "1937", iyr: "2017", eyr: "2020", hgt: "183cm", hcl: "#fffffd", ecl: "gry", pid: "860033327", cid: "147"}, true},
		{passport{byr: "1937", iyr: "2017", eyr: "2020", hgt: "183cm", hcl: "#fffffd", ecl: "gry"}, false},
		{passport{byr: "1937", iyr: "2017", eyr: "2020", hgt: "183cm", hcl: "#fffffd", pid: "860033327"}, false},
		{passport{byr: "1937", iyr: "2017", eyr: "2020", hgt: "183cm", ecl: "gry", pid: "860033327"}, false},
		{passport{byr: "1937", iyr: "2017", eyr: "2020", hcl: "#fffffd", ecl: "gry", pid: "860033327"}, false},
		{passport{byr: "1937", iyr: "2017", hgt: "183cm", hcl: "#fffffd", ecl: "gry", pid: "860033327"}, false},
		{passport{byr: "1937", eyr: "2020", hgt: "183cm", hcl: "#fffffd", ecl: "gry", pid: "860033327"}, false},
		{passport{iyr: "2017", eyr: "2020", hgt: "183cm", hcl: "#fffffd", ecl: "gry", pid: "860033327"}, false},
		{passport{}, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.p), func(t *testing.T) {
			t.Parallel()

			if got := tt.p.isValid(); tt.want != got {
				t.Errorf("Expected %v got %v", tt.want, got)
			}
		})
	}
}

func TestValidatePassports(t *testing.T) {
	input := `
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`

	got, err := validatePassports(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	if want := 2; want != got {
		t.Errorf("Expected %d got %d", want, got)
	}
}
