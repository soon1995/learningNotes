package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		a, want [5]int
	}{
		{[5]int{1, 2, 3, 4, 5}, [5]int{5, 4, 3, 2, 1}},
	}
	for _, test := range tests {
		reverse(&test.a)
		if test.a != test.want {
			t.Errorf("got %v, want %v\n", test.a, test.want)
		}
	}
}
