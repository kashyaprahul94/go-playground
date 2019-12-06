package util

// CompareSlices tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice
func CompareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// CloneArray can be used to create clone numbers
func CloneArray(input []int) []int {

	length := len(input)
	numbers := make([]int, length, length)

	for i := 0; i < length; i++ {
		numbers[i] = input[i]
	}

	return numbers
}
