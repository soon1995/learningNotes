// Make display safe to use on cyclic data structures by bounding the number
// of steps it takes before abandoning the recursion. (In Section 13.3), we'll
// see another way to detect cycles.
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
	display(name, reflect.ValueOf(x), 0)
}

func display(path string, v reflect.Value, level int) {
	if level > 5 {
		fmt.Printf("%s = %s\n", path, formatAtom(v))
		return
	}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), level+1)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), level+1)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key), level+1)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), level+1)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), level+1)
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

func main() {
	type Cycle struct {
		Value int
		C     *Cycle
	}
	c := &Cycle{}
	c.Value = 1
	c.C = c
	Display("Cycle", c)
}
