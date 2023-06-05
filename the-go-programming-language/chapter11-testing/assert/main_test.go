package main

import (
	"fmt"
	"strings"
	"testing"
)

// a poor assertion function.
// suffer from premature abstraction: by treating
// the failure of this particular test as a mere difference of two integers.
// we forfeit the opportunity to provide meaningful context.
// Only once repetitive pattern emerge in a given test suite is it time to introduce
// abstractions
func assertEqual(x, y int) {
	if x != y {
		panic(fmt.Sprintf("%d != %d", x, y))
	}
}

func TestSplit(t *testing.T) {
	words := strings.Split("a:b:c", ":")
	assertEqual(len(words), 3)
}

func TestSplit1(t *testing.T) {
	s, sep := "a:b:c", ":"
	words := strings.Split(s, sep)
	if got, want := len(words), 3; got != want {
		t.Errorf("Split(%q, %q) returned %d words, want %d", s, sep, got, want)
	}
}
