package arrays_and_slices

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	sum := Sum(numbers)
	expectedSum := 15

	if sum != expectedSum {
		t.Errorf("expected '%d' but got '%d', given arraz %v", expectedSum, sum, numbers)
	}
}
