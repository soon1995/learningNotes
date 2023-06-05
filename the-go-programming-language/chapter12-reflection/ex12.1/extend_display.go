// Extend Display so that it can display mas whose keys are structs or arrays
package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

// gopl.io/ch12/display
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
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
	case reflect.Struct:
		var buf bytes.Buffer
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				buf.WriteString(", ")
			}
			fmt.Fprintf(&buf, "%s: %s", v.Type().Field(i).Name, formatAtom(v.Field(i)))
		}
		buf.WriteByte('}')
		return buf.String()
	case reflect.Array:
		var buf bytes.Buffer
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(formatAtom(v.Index(i)))
		}
		buf.WriteByte(']')
		return buf.String()
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

type Movie struct {
	ActorsExp map[Exp]int
}

type Exp struct {
	Year int
	Month int
}

type Max map[[3]int]int

func main() {
	max := make(Max)
	max[[3]int{1, 2, 3}] = 3
	max[[3]int{4, 6, 5}] = 6
	Display("max", max)

	movie := &Movie{
		ActorsExp: make(map[Exp]int),
	}
	movie.ActorsExp[Exp{Year: 2}] = 5
	movie.ActorsExp[Exp{Year: 3}] = 5
	Display("movie", movie)
}
