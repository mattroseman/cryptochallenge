package galois

import (
	"testing"
)

func TestOneGetMSBPosition(t *testing.T) {
	got := getMSBPosition(uint16(1))
	want := uint16(1)

	if got != want {
		t.Fatalf("getMSBPosition(%#x) got %#x, want %#x", uint16(1), got, want)
	}
}

func TestZeroGetMSBPositions(t *testing.T) {
	got := getMSBPosition(uint16(0))
	want := uint16(0)

	if got != want {
		t.Fatalf("getMSBPosition(%#x) got %#x, want %#x", uint16(0), got, want)
	}
}

func TestPrimitiveGetMSBPosition(t *testing.T) {
	got := getMSBPosition(uint16(0x11d))
	want := uint16(9)

	if got != want {
		t.Fatalf("getMSBPosition(%#x) got %#x, want %#x", uint16(0x11d), got, want)
	}
}

func TestReduce(t *testing.T) {
	a := uint16(0x147a)
	got := reduce(a)
	want := byte(0xc3)

	if got != want {
		t.Fatalf("recude(%#x) got %#x, want %#x", a, got, want)
	}
}

func TestReduceZero(t *testing.T) {
	a := uint16(0x0)
	got := reduce(a)
	want := byte(0x0)

	if got != want {
		t.Fatalf("recude(%#x) got %#x, want %#x", a, got, want)
	}
}
