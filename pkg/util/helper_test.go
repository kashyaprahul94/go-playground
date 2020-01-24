package util

import (
	"testing"
)

func TestCompareSlice(t *testing.T) {

	testCases := []struct {
		name   string
		input1 []int
		input2 []int
		want   bool
	}{
		{name: "basic case", input1: []int{1, 2, 3, 4}, input2: []int{1, 2, 3, 4}, want: true},
		{name: "same size but different elements", input1: []int{1, 2, 3}, input2: []int{3, 3, 3}, want: false},
		{name: "different elements", input1: []int{1, 2, 3, 4}, input2: []int{1}, want: false},
		{name: "empty case", input1: []int{}, input2: []int{}, want: true},
	}

	for _, tt := range testCases {

		t.Run(tt.name, func(t *testing.T) {

			input1 := tt.input1
			input2 := tt.input2

			got := CompareSlices(input1, input2)

			if got != tt.want {
				t.Errorf("expected comparision of %v & %v as %v but got %v", input1, input2, true, got)
			}
		})
	}

}

func TestCloneSlice(t *testing.T) {

	testCases := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "empty case", input: []int{}, want: []int{}},
		{name: "basic case", input: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}},
	}

	for _, tt := range testCases {

		t.Run(tt.name, func(t *testing.T) {

			input := tt.input

			result := CloneSlice(input)

			if !CompareSlices(input, result) {
				t.Errorf("expected %v & %v to be same, but they are not", input, result)
			}

			originalAddr := &input
			resultAddr := &result

			if originalAddr == resultAddr {
				t.Errorf("Cloned slice should not point to same memory location as input")
			}

			if len(input) > 0 {
				input[0] = 555555

				if CompareSlices(input, result) {
					t.Errorf("changes made to original slice shouldn't affect the cloned one")
				}

				result[0] = 9999

				if CompareSlices(input, result) {
					t.Errorf("changes made to cloned slice shouldn't affect the original one")
				}
			}

		})
	}

}
