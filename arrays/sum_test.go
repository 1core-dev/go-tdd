package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{2, 3, 4, 6, 10}

	got := Sum(numbers)
	want := 25
	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}
