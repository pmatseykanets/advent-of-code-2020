package main

import (
	"reflect"
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

func TestCanBeContained(t *testing.T) {
	input := `
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	invertedRules, err := parseRules(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	want := map[string]struct{}{
		"bright white": struct{}{},
		"dark orange":  struct{}{},
		"light red":    struct{}{},
		"muted yellow": struct{}{},
	}

	if got := canBeContained("shiny gold", invertedRules); !reflect.DeepEqual(want, got) {
		t.Errorf("Expected colors\n%v\ngot\n%v", want, got)
	}
}
