package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, expected, repeated string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	}

	t.Run("test repeat when input is 'a'", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("test repeat when input is 'b'", func(t *testing.T) {
		repeated := Repeat("b")
		expected := "bbbbb"
		assertCorrectMessage(t, expected, repeated)
	})
}
