package arrays_and_slices

func Sum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(inputNumbers ...[]int) (result []int) {

	result = make([]int, len(inputNumbers))

	for i, numbers := range inputNumbers {
		result[i] = Sum(numbers)
	}

	return result
}
