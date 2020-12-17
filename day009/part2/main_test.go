package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestLoadNumbers(t *testing.T) {
	input := `
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

	got, err := loadNumbers(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	want := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected numbers\n%+v\ngot\n%#v", want, got)
	}

}

func TestFindInvalidNumber(t *testing.T) {
	numbers := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

	if want, got := 127, findInvalid(numbers, 5); want != got {
		t.Errorf("Expected invalid number %d got %d", want, got)
	}
}

func TestFindWeakness(t *testing.T) {
	numbers := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

	if want, got := 62, findWeakness(numbers, 127); want != got {
		t.Errorf("Expected weakness %d got %d", want, got)
	}
}
