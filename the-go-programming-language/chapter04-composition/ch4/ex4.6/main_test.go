package main

import (
	"reflect"
	"testing"
)

func TestDeleteAdjadentSpace(t *testing.T) {
	tests := []struct {
		a    string
		want string
	}{
		{"你\n\n好\t\t吗  ", "你 好 吗 "},
		{"\n\n你好\t\t\n吗  ", " 你好 吗 "},
	}
	for _, test := range tests {
		res := Squashes([]byte(test.a))
		if !reflect.DeepEqual(string(res), test.want) {
			t.Errorf("got %v, want %v", string(res), test.want)
		}
	}
}
