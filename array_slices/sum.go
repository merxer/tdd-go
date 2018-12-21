package array_slices

func Sum(numbers []int) (sum int) {
	sum = 0
	for _, v := range numbers {
		sum = sum + v
	}
	return
}

func SumAll(numbers ...[]int) []int {

	lengthOfNumbers := len(numbers)
	sums := make([]int, lengthOfNumbers)

	for i, number := range numbers {
		sums[i] = Sum(number)
	}
	return sums
}
