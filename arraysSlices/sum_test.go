package arraysslices

import "testing"

func TestSum(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertCorrectArraySum(t, got, want, numbers)
	})
}

func assertCorrectArraySum(t testing.TB, got int, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// go test -cover gives how well our code is covered
