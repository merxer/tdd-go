package array_slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumWhenInput1_to_5(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	sum := Sum(numbers)
	expected := 15

	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}
}

func TestSumWhenInput1_to_2(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4}

	sum := Sum(numbers)
	expected := 10

	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}
}

func TestSumWhenInput3to2(t *testing.T) {
	numbers := []int{3, 2}

	sum := Sum(numbers)
	expected := 5
	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}
}

func ExampleSum() {
	s := []int{5, 5, 5, 5}
	sum := Sum(s)
	fmt.Printf("%d", sum)
	// Output: 20
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{i, b.N})
	}
}

func TestSumAll(t *testing.T) {
	sum := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(sum, expected) {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}
}
func TestSumAll_2(t *testing.T) {
	sum := SumAll([]int{1, 2}, []int{2, 3})
	expected := []int{3, 5}

	if !reflect.DeepEqual(sum, expected) {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}
}

func TestSumAllTails(t *testing.T) {
	sumTails := SumAllTails([]int{1, 2}, []int{0, 9})
	expected := []int{2, 9}

	if !reflect.DeepEqual(sumTails, expected) {
		t.Errorf("expected '%d', but got '%d'", expected, sumTails)
	}
}
