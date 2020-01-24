package algorithm

import (
	"testing"

	"github.com/kashyaprahul94/go-playground/pkg/util"
)

func assertResult(t *testing.T, expected, received []int) {
	t.Helper()

	if !util.CompareSlices(expected, received) {
		t.Errorf("expected %v received %v", expected, received)
	}
}

func TestSortings(t *testing.T) {

	basicCaseInput := []int{5, 2, 3, 1, 4}
	basicCaseExpected := []int{1, 2, 3, 4, 5}

	sortingTests := []struct {
		name     string
		method   Sorter
		expected []int
	}{
		{name: "BubbleSort", method: &BubbleSort{Numbers: basicCaseInput}, expected: basicCaseExpected},
		{name: "SelectionSort", method: &SelectionSort{Numbers: basicCaseInput}, expected: basicCaseExpected},
		{name: "InsertionSort", method: &InsertionSort{Numbers: basicCaseInput}, expected: basicCaseExpected},
	}

	for _, tt := range sortingTests {
		t.Run(tt.name, func(t *testing.T) {
			assertResult(t, tt.expected, tt.method.Sort())
		})
	}
}
