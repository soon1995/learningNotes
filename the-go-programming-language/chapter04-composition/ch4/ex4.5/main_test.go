package main

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicate(t *testing.T) {
	tests := []struct {
		a, want []string
	}{
		{[]string{"a", "b", "b", "c", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "b", "d", "c", "c"}, []string{"a", "b", "d", "c"}},
		{[]string{"a", "a", "b", "c", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "b", "c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{"a"}, []string{"a"}},
		{[]string{}, []string{}},
	}
	for _, test := range tests {
		test.a = deleteDuplicate(test.a)
		if !reflect.DeepEqual(test.a, test.want) {
			t.Errorf("got %v, want %v", test.a, test.want)
		}
	}
}
