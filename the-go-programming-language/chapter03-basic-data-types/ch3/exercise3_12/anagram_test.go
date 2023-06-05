package main

import (
	"fmt"
	"testing"
	"time"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"aba", "baa", true},
		{"aaa", "baa", false}, // same characters but different frequencies
		{"aaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", true}, // same characters but different frequencies
	}
	for _, test := range tests {
		t1 := time.Now()
		got := isAnagramExample(test.a, test.b)
		fmt.Printf("time taken: %v\n", time.Since(t1))
		if got != test.want {
			t.Errorf("isAnagram(%q, %q), got %v, want %v",
				test.a, test.b, got, test.want)
		}
	}
	for _, test := range tests {
		t1 := time.Now()
		got := isAnagram(test.a, test.b)
		fmt.Printf("time taken: %v\n", time.Since(t1))
		if got != test.want {
			t.Errorf("isAnagram(%q, %q), got %v, want %v",
				test.a, test.b, got, test.want)
		}
	}
}
