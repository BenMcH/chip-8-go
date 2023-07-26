package commands

import "testing"

func TestCLS(t *testing.T) {
	opcode := CLS()

	if opcode != "00E0" {
		t.Errorf("Expected CLS to return opcode 00E0, got %s", opcode)
	}
}
