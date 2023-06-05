package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestBool(t *testing.T) {
	tcs := []struct {
		i    bool
		want string
	}{
		{true, `t`},
		{false, `nil`},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.i))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%v) got %s, want %s", tc.i, got, tc.want)
		}
	}
}

func TestFloat32(t *testing.T) {
	tcs := []struct {
		i    float32
		want string
	}{
		{3.2e9, "3.2e+09"},
		{1.0, "1"},
		{0, "0"},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.i))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%g) got %s, want %s", tc.i, got, tc.want)
		}
	}
}

func TestFloat64(t *testing.T) {
	tcs := []struct {
		i    float64
		want string
	}{
		{3.2e9, "3.2e+09"},
		{1.0, "1"},
		{0, "0"},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.i))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%g) got %s, want %s", tc.i, got, tc.want)
		}
	}
}

func TestComplex64(t *testing.T) {
	tcs := []struct {
		i    complex64
		want string
	}{
		{0 + 0i, "#C(0 0)"},
		{3 - 2i, "#C(3 -2)"},
		{-1e9 + -2.2e9i, "#C(-1e+09 -2.2e+09)"},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.i))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%g) got %s, want %s", tc.i, got, tc.want)
		}
	}
}

func TestComplex128(t *testing.T) {
	tcs := []struct {
		i    complex128
		want string
	}{
		{0 + 0i, "#C(0 0)"},
		{3 - 2i, "#C(3 -2)"},
		{-1e9 + -2.2e9i, "#C(-1e+09 -2.2e+09)"},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.i))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%g) got %s, want %s", tc.i, got, tc.want)
		}
	}
}

func TestInterface(t *testing.T) {
	type Wrapper struct {
		v interface{}
	}
	tcs := []struct {
		w    Wrapper
		want string
	}{
		{Wrapper{"hello"}, `((v ("string" "hello")))`},
		{Wrapper{[]int{1, 2, 3}}, `((v ("[]int" (1 2 3))))`},
		{Wrapper{[3]int{1, 2, 3}}, `((v ("[3]int" (1 2 3))))`},
		{Wrapper{nil}, `((v ("nil" nil)))`},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.w))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%v) got %s, want %s", tc.w, got, tc.want)
		}
	}
}
