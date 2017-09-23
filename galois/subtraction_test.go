package galois

import "testing"

func TestSub(t *testing.T) {
	a := byte(5)
	b := byte(6)
	got := Sub(a, b)
	want := byte(3)

	if got != want {
		t.Fatalf("Sub(%#x, %#x) = %#x, want %#x", a, b, got, want)
	}
}
