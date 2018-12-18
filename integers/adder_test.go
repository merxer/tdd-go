package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestAdderWhenInputIs3_3(t *testing.T) {
	sum := Add(3, 3)
	expected := 6

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
