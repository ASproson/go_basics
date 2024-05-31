package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("passes a and repeats 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"
		assertCorrectRepeats(t, repeated, expected)
	})
	t.Run("passes a and repeats 2 times", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"
		assertCorrectRepeats(t, repeated, expected)
	})
}

func assertCorrectRepeats(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("repeated %q expected %q", repeated, expected)
	}
}

// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}
