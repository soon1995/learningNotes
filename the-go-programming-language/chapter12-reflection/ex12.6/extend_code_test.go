package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestFloat32(t *testing.T) {
	type Wrapper struct {
		v float32
	}
	tcs := []struct {
		v    Wrapper
		want string
	}{
		{Wrapper{0}, "()"},
		{Wrapper{1.0}, `((v 1))`},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.v))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%g) got %s, want %s", tc.v, got, tc.want)
		}
	}
}

func TestFloat64(t *testing.T) {
	type Wrapper struct {
		v float64
	}
	tcs := []struct {
		v    Wrapper
		want string
	}{
		{Wrapper{0}, "()"},
		{Wrapper{1.0}, `((v 1))`},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.v))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%g) got %s, want %s", tc.v, got, tc.want)
		}
	}
}

func TestComplex64(t *testing.T) {
	type Wrapper struct {
		v complex64
	}
	tcs := []struct {
		v    Wrapper
		want string
	}{
		{Wrapper{0 + 0i}, "()"},
		{Wrapper{3 - 2i}, `((v #C(3 -2)))`},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.v))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%g) got %s, want %s", tc.v, got, tc.want)
		}
	}
}

func TestComplex128(t *testing.T) {
	type Wrapper struct {
		v complex128
	}
	tcs := []struct {
		v    Wrapper
		want string
	}{
		{Wrapper{0 + 0i}, "()"},
		{Wrapper{3 - 2i}, `((v #C(3 -2)))`},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.v))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%g) got %s, want %s", tc.v, got, tc.want)
		}
	}
}

func TestInt(t *testing.T) {
	type Wrapper struct {
		v int
	}
	tcs := []struct {
		v    Wrapper
		want string
	}{
		{Wrapper{0}, "()"},
		{Wrapper{123}, `((v 123))`},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.v))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%d) got %s, want %s", tc.v, got, tc.want)
		}
	}
}

func TestPointer(t *testing.T) {
  type InnerStruct struct{ }
	type Wrapper struct {
		v *InnerStruct
	}
	tcs := []struct {
		v    Wrapper
		want string
	}{
		{Wrapper{nil}, "()"},
		{Wrapper{&InnerStruct{}}, `((v ()))`},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.v))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%v) got %s, want %s", tc.v, got, tc.want)
		}
	}
}

func TestSlice(t *testing.T) {
	type Wrapper struct {
		v []int
	}
	tcs := []struct {
		v    Wrapper
		want string
	}{
		{Wrapper{nil}, "()"},
		{Wrapper{[]int{}}, `((v ()))`},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.v))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%v) got %s, want %s", tc.v, got, tc.want)
		}
	}
}

func TestArray(t *testing.T) {
	type Wrapper struct {
		v [3]int
	}
	tcs := []struct {
		v    Wrapper
		want string
	}{
		{Wrapper{[3]int{}}, "()"},
		{Wrapper{[3]int{1}}, `((v (1 0 0)))`},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.v))
		got := b.String()
		if tc.want != got {
			t.Errorf("Encode(%v) got %s, want %s", tc.v, got, tc.want)
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
		{Wrapper{[3]int{1, 2, 3}}, `((v ("[3]int" (1 2 3))))`},
		{Wrapper{nil}, `()`},
	}
	for _, tc := range tcs {
		b := &bytes.Buffer{}
		encode(b, reflect.ValueOf(tc.w))
		got := b.String()
		if tc.want != got {
			t.Errorf("Marshal(%v) got %s, want %s", tc.w, got, tc.want)
		}
	}
}
