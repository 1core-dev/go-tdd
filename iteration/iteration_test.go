package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}

}

func BenchmarkRepeatV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatV2("a")
	}

}
