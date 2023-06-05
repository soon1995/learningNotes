// Write a function tat reports whether its argument is a cyclic data structure.
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func cyclic(x reflect.Value, seen map[comparison]bool) bool {
	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		c := comparison{xptr, x.Type()}
		fmt.Println(xptr)
		fmt.Println(x.Type())
		if seen[c] {
			return true
		}
		seen[c] = true
	}

	fmt.Printf("x=%s", x.Kind())

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return cyclic(x.Elem(), seen)
	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if cyclic(x.Index(i), seen) {
				return true
			}
		}
		return false
	case reflect.Struct:
		for i := 0; i < x.NumField(); i++ {
			if cyclic(x.Field(i), seen) {
				return true
			}
		}
		return false
	case reflect.Map:
		for _, k := range x.MapKeys() {
			if cyclic(x.MapIndex(k), seen) || cyclic(k, seen) {
				return true
			}
		}
		return false
	}
	return false
}

func Cyclic(x interface{}) bool {
	seen := make(map[comparison]bool)
	return cyclic(reflect.ValueOf(x), seen)
}

type comparison struct {
	x unsafe.Pointer
	t reflect.Type
}
