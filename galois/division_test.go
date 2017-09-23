package galois

import "testing"

func TestDiv(t *testing.T) {
	a := byte(0xc3)
	b := byte(0x89)
	got, err := Div(a, b)
	if err != nil {
		panic(err)
	}
	want := byte(0x2a)

	if got != want {
		t.Fatalf("Div(%#x, %#x) got %#x, want %#x", a, b, got, want)
	}
}

func TestDivByZero(t *testing.T) {
	a := byte(0xc3)
	b := byte(0x0)
	got, err := Div(a, b)
	if err == nil {
		t.Fatalf("Div(%#x, %#x) got %#x, when it should have thrown a divide by 0 error", a, b, got)
	}
}

func TestDivZero(t *testing.T) {
	a := byte(0x0)
	b := byte(0xc3)
	got, err := Div(a, b)
	if err != nil {
		panic(err)
	}
	want := byte(0x0)

	if got != want {
		t.Fatalf("Div(%#x, %#x) got %#x, want %#x", a, b, got, want)
	}
}
