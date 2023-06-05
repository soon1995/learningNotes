package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		a, want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 1}},
	}
	for _, test := range tests {
		rotate(test.a)
		if !reflect.DeepEqual(test.a, test.want) {
			t.Errorf("got %v, want %v", test.a, test.want)
		}
	}
}

func TestRotateEx(t *testing.T) {
	rotateEx()
}
