package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEnc(t *testing.T) {
	buf := &bytes.Buffer{}
	type Test struct {
		Name string `sexpr:"nama"`
		Age  int    `sexpr:"umur"`
		Work bool   `sexpr:"kerja"`
	}
	mystruct := &Test{"ABC", 16, false}
	want := `((nama "ABC") (umur 16) (kerja false))`
	if err := encode(buf, reflect.ValueOf(mystruct)); err != nil {
		t.Errorf("unexpected error encode(%#v), err: %s", mystruct, err)
	}
	if want != buf.String() {
		t.Errorf("encode(%#v) want: %s, got %s", mystruct, want, buf.String())
	}
}

func TestDec(t *testing.T) {
	type Test struct {
		Name string `sexpr:"nama"`
		Age  int    `sexpr:"umur"`
		Work bool   `sexpr:"kerja"`
	}
	test := &Test{}
	from := `((nama "ABC") (umur 16) (kerja true))`
	r := bytes.NewReader([]byte(from))
	dec := NewDecoder(r)
	dec.Decode(test)

	mystruct := &Test{"ABC", 16, true}
	if !reflect.DeepEqual(mystruct, test) {
		t.Errorf("decode(%s) want: %#v, got %#v", from, mystruct, test)
	}
}
