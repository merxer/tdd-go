package array_slices

func Sum(numbers []int) (sum int) {
	sum = 0
	for _, v := range numbers {
		sum = sum + v
	}
	return
}

func SumAll(numbers ...[]int) []int {

	var sums []int

	for _, number := range numbers {
		sums = append(sums, Sum(number))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, nums := range numbers {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
