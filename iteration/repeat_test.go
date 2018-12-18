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
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("test repeat when input is 'b'", func(t *testing.T) {
		repeated := Repeat("b", 5)
		expected := "bbbbb"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("test repeat when input is 'A'", func(t *testing.T) {
		repeated := Repeat("A", 5)
		expected := "AAAAA"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("test repeat string when input 'C' 3 times", func(t *testing.T) {
		repeated := Repeat("C", 3)
		expected := "CCC"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("test repeat string when input 'D' 10 times", func(t *testing.T) {
		repeated := Repeat("D", 10)
		expected := "DDDDDDDDDD"
		assertCorrectMessage(t, expected, repeated)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("8", i)
	}
}
