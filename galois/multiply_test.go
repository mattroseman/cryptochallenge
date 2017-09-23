package galois

import (
	"testing"
)

func TestSlowMult(t *testing.T) {
	a := byte(0x89)
	b := byte(0x2a)
	got := slowMult(a, b)
	want := byte(0xc3)

	if got != want {
		t.Fatalf("slowMult(%#x, %#x) got %#x, want %#x", a, b, got, want)
	}
}

func TestSlowMultByZero(t *testing.T) {
	a := byte(0x89)
	b := byte(0x0)
	got := slowMult(a, b)
	want := byte(0x0)

	if got != want {
		t.Fatalf("slowMult(%#x, %#x) got %#x, want %#x", a, b, got, want)
	}
}

func TestMult(t *testing.T) {
	a := byte(0x89)
	b := byte(0x2a)
	got := Mult(a, b)
	want := byte(0xc3)

	if got != want {
		t.Fatalf("Mult(%#x, %#x) got %#x, want %#x", a, b, got, want)
	}
}

func TestMultByZero(t *testing.T) {
	a := byte(0x89)
	b := byte(0x0)
	got := Mult(a, b)
	want := byte(0x0)

	if got != want {
		t.Fatalf("Mult(%#x, %#x) got %#x, want %#x", a, b, got, want)
	}
}
