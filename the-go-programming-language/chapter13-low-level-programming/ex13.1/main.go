// Define a deep comparison function that considers numbers (of any type) equal
// if they differ by less than one part in a billion
package main

import (
	"reflect"
	"unsafe"
)

const multiplier = 1000000000

// copied torbiak/gopl/ex13.1/equal.go
func numberEqual(x, y float64) bool {
	if x == y {
		return true
	}
	var diff float64
	if x > y {
		diff = x - y
	} else {
		diff = y - x
	}
	d := diff * multiplier
	if d < x && d < y {
		return true
	}
	return false
}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true
		}
		seen[c] = true
	}

	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}
	// ...cycle check omitted (shown later)...

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()
	case reflect.String:
		return x.String() == y.String()
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return numberEqual(float64(x.Int()), float64(y.Int()))
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return numberEqual(float64(x.Uint()), float64(y.Uint()))
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()
	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)
	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true
		// ...struct and map cases omitted for brevity ...
	}
	panic("unreachable")
}

func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}
