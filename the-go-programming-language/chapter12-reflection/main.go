package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	// ...similar cases for int16, uint32, and so on...
	case bool:
		return strconv.FormatBool(x)
	default:
		// array, chan, func, map, pointer, slice, struct
		return "???"
	}
}

// gopl.io/ch 12/format
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func main() {
	stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
	fmt.Println(stdout.Type())                  // "os.File"
	fd := stdout.FieldByName("fd")
	fmt.Println(fd.Int()) // "int"
	fd.SetInt(2)          // panic: unexported field
}
