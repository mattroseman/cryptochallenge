package galois

import (
	"testing"
)

func TestGenerateExponentTable(t *testing.T) {
	table := GenerateExponentTable()
	for i, exp := range table {
		if expTable[i] != exp {
			t.Fatalf("ExponentTable at index %d got %#x, want %#x", i, expTable[i], exp)
		}
	}
}

func TestGenerateLogarithmTable(t *testing.T) {
	table := GenerateLogarithmTable()
	for i, log := range table {
		if logTable[i] != log {
			t.Fatalf("LogarithmTable at index %d got %#x, want %#x", i, logTable[i], log)
		}
	}
}
