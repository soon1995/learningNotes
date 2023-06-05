package charcount

import (
	"reflect"
	"testing"
)

func TestCharcount(t *testing.T) {
	tests := []struct {
		input           string
		wantCharCount   map[rune]int
		wantUtfLenCount [4]int
	}{
		{
			input:           "Hi, 你好!",
			wantCharCount:   map[rune]int{'H': 1, 'i': 1, ',': 1, ' ': 1, '你': 1, '好': 1, '!': 1},
			wantUtfLenCount: [...]int{0, 5, 0, 2},
		},
	}

	for _, test := range tests {
		gotCharcount, gotUtfLenCount := Charcount(test.input)
		if !reflect.DeepEqual(test.wantCharCount, gotCharcount) {
			t.Errorf("Charcount(%q) got %v, want %v", test.input, gotCharcount, test.wantCharCount)
		}
		if !reflect.DeepEqual(test.wantUtfLenCount, gotUtfLenCount) {
			t.Errorf("Charcount(%q) got %v, want %v", test.input, gotUtfLenCount, test.wantUtfLenCount)
		}
	}
}
