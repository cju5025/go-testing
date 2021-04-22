package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("testing on single character input", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})

	t.Run("testing on multiple character input", func(t *testing.T) {
		got := Repeat("she", 3)
		want := "shesheshe"
		assertCorrectMessage(t, got, want)
	})

	t.Run("testing on multiple character input with spaces", func(t *testing.T) {
		got := Repeat("she wants ", 3)
		want := "she wants she wants she wants "
		assertCorrectMessage(t, got, want)
	})
}

// go test -bench=. is the command to run a benchmark test
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
