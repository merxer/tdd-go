package iteration

import "testing"

func TestRepeatWhenInputIsA(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func TestRepeatWhenInputIsB(t *testing.T) {
	repeated := Repeat("b")
	expected := "bbbbb"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}
