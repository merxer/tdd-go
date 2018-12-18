package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, sum, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}

	t.Run("test adder when input is 2, 2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, sum, expected)
	})

	t.Run("test adder when input is 3, 3", func(t *testing.T) {
		sum := Add(3, 3)
		expected := 6
		assertCorrectMessage(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
