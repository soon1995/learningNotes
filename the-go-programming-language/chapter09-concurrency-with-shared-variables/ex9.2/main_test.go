package main

import "testing"

func TestPopCount(t *testing.T) {
	testcases := []struct {
		data uint64
		want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
	}
	for _, tc := range testcases {
		got := PopCount(tc.data)
		if got != tc.want {
			t.Errorf("PopCount %d - want %d, got %d", tc.data, tc.want, got)
		}
	}
}
