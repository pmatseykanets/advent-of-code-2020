package main

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestParseRule(t *testing.T) {
	tests := []struct {
		input string
		want  rule
	}{
		{
			input: "vibrant salmon bags contain 1 vibrant gold bag, 2 wavy aqua bags, 1 dotted crimson bag.",
			want: rule{
				color: "vibrant salmon",
				contains: map[string]int{
					"vibrant gold":   1,
					"wavy aqua":      2,
					"dotted crimson": 1,
				},
			},
		},
		{
			input: "dotted plum bags contain 3 wavy cyan bags.",
			want: rule{
				color: "dotted plum",
				contains: map[string]int{
					"wavy cyan": 3,
				},
			},
		},
		{
			input: "dark tomato bags contain no other bags.",
			want: rule{
				color: "dark tomato",
				// contains: map[string]int{},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got, err := parseRule(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Expected rule\n%+v\ngot\n%+v", tt.want, got)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`,
			want: 32,
		},
		{
			input: `
shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`,
			want: 126,
		},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			rules, err := parseRules(strings.NewReader(tt.input))
			if err != nil {
				t.Fatal(err)
			}
			// fmt.Printf("%#v\n", rules)
			if got := contains("shiny gold", rules) - 1; tt.want != got {
				t.Errorf("Expected bags %d got %d", tt.want, got)
			}
		})
	}
}
