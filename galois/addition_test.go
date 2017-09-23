package galois

import "testing"

func TestAdd(t *testing.T) {
	a := byte(5)
	b := byte(6)
	got := Add(a, b)
	want := byte(3)

	if got != want {
		t.Fatalf("Add(%#x, %#x) = %#x, want %#x", a, b, got, want)
	}
}
