package array_slices

func Sum(numbers [5]int) (sum int) {
	sum = 0
	for _, v := range numbers {
		sum = sum + v
		println(v)
	}
	return
}
