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
		// Valid
		{passport{byr: "1980", iyr: "2012", eyr: "2030", hgt: "74in", hcl: "#623a2f", ecl: "grn", pid: "087499704", cid: ""}, true},
		{passport{byr: "1989", iyr: "2014", eyr: "2029", hgt: "165cm", hcl: "#a97842", ecl: "blu", pid: "896056539", cid: "129"}, true},
		{passport{byr: "2001", iyr: "2015", eyr: "2022", hgt: "164cm", hcl: "#888785", ecl: "hzl", pid: "545766238", cid: "88"}, true},
		{passport{byr: "1944", iyr: "2010", eyr: "2021", hgt: "158cm", hcl: "#b6652a", ecl: "blu", pid: "093154719", cid: ""}, true},
		// Invalid
		{passport{byr: "1926", iyr: "2018", eyr: "1972", hgt: "170", hcl: "#18171d", ecl: "amb", pid: "186cm", cid: "100"}, false},
		{passport{byr: "1946", iyr: "2019", eyr: "1967", hgt: "170cm", hcl: "#602927", ecl: "grn", pid: "012533040", cid: ""}, false},
		{passport{byr: "1992", iyr: "2012", eyr: "2020", hgt: "182cm", hcl: "dab227", ecl: "brn", pid: "021572410", cid: "277"}, false},
		{passport{byr: "2007", iyr: "2023", eyr: "2038", hgt: "59cm", hcl: "74454a", ecl: "zzz", pid: "3556412378", cid: ""}, false},
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
pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719

eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007
`

	got, err := validatePassports(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	if want := 4; want != got {
		t.Errorf("Expected %d got %d", want, got)
	}
}
