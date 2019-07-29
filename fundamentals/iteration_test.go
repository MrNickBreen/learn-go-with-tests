package fundamentals

import "testing"

func assertCorrectRepeat(t *testing.T, repeated string, expected string) {
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectRepeat(t, repeated, expected)
	})

	t.Run("repeat once", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"
		assertCorrectRepeat(t, repeated, expected)
	})

	t.Run("repeat zero", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertCorrectRepeat(t, repeated, expected)
	})
}

func BenchmarkREpeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1)
	}
}
