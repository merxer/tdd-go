package array_slices

import "testing"

func TestSumWhenInput1_to_5(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	sum := Sum(numbers)
	expected := 15

	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}
}

func TestSumWhenInputWhenInput1_to_2(t *testing.T) {
	numbers := [5]int{0, 1, 2, 3, 4}

	sum := Sum(numbers)
	expected := 10

	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}
}
