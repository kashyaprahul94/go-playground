package algorithm

import "github.com/kashyaprahul94/go-playground/pkg/util"

// Sorter is the common interface that each sorting technique should adhere to
type Sorter interface {
	Sort() []int
}

//
//

/*
	Bubble Sort
*/

// BubbleSort should operate upon list of numbers
type BubbleSort struct {
	Numbers []int
}

// Sort helps to sort
func (b *BubbleSort) Sort() []int {

	length := len(b.Numbers)
	numbers := util.CloneSlice(b.Numbers)

	for i := length; i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			}
		}
	}

	return numbers
}

/*
	Selection Sort
*/

// SelectionSort should operate upon list of numbers
type SelectionSort struct {
	Numbers []int
}

// Sort helps to sort
func (s *SelectionSort) Sort() []int {

	length := len(s.Numbers)
	numbers := util.CloneSlice(s.Numbers)

	for i := 0; i < length; i++ {
		minimumIndex := i

		for j := i; j < length; j++ {
			if numbers[j] < numbers[minimumIndex] {
				minimumIndex = j
			}
		}

		if minimumIndex != i {
			numbers[i], numbers[minimumIndex] = numbers[minimumIndex], numbers[i]
		}
	}

	return numbers
}

/*
	Insertion Sort
*/

// InsertionSort should operate upon list of numbers
type InsertionSort struct {
	Numbers []int
}

// Sort helps to sort
func (i *InsertionSort) Sort() []int {

	length := len(i.Numbers)
	numbers := util.CloneSlice(i.Numbers)

	for i := 1; i < length; i++ {

		element := numbers[i]
		j := i - 1

		for j >= 0 && numbers[j] > element {

			numbers[j+1] = numbers[j]

			j--
		}

		numbers[j+1] = element
	}

	return numbers
}
