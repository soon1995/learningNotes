package main

import (
	"testing"
)

func TestCyclic(t *testing.T) {
	type A struct {
		Cycle *A
	}
	tc1a := A{}
	tc1a.Cycle = &tc1a
	if !Cyclic(tc1a) {
		t.Errorf("Cyclic(%#v) got %t", tc1a, false)
	}

	tc2a := A{}
	tc2b := A{}
	tc2a.Cycle = &tc2b
	if Cyclic(tc2a) {
		t.Errorf("Cyclic(%#v) got %t", tc1a, true)
	}
}
