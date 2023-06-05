// Extend TestSplit to use a table of inputs and expected outputs.
package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tcs := []struct {
		in   string
		sep  string
		want []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"a!b!c", "!", []string{"a", "b", "c"}},
		{"a", "!", []string{"a"}},
		{"ab", "!", []string{"ab"}},
		{":ab", ":", []string{"", "ab"}},
		{"ab:", ":", []string{"ab", ""}},
		{"", ":", []string{""}},
	}
	for _, v := range tcs {

		words := strings.Split(v.in, v.sep)
		if !reflect.DeepEqual(words, v.want) {
			// if len(words) != len(v.want) {
			t.Errorf("Split(%q, %q) returned %v, want %v", v.in, v.sep, words, v.want)
		}
	}
}
