// Modify encode to pretty-print the S-expression in the style shown above
package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func encode(buf *bytes.Buffer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		fmt.Fprintf(buf, "#C(%g %g)", real(c), imag(c))
	case reflect.Bool:
		if v.Bool() {
			buf.WriteByte('t')
		} else {
			buf.WriteString("nil")
		}
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Pointer:
		return encode(buf, v.Elem(), indent)
	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		indent += 1
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				fmt.Fprintf(buf, "\n%*s", indent, "")
			}
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}
		}
		buf.WriteByte(')')
	case reflect.Struct:
		buf.WriteByte('(')
		indent += 1
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				fmt.Fprintf(buf, "\n%*s", indent, "")
			}
			start := buf.Len()
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indent+buf.Len()-start); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Map:
		buf.WriteByte('(')
		indent += 1
		for i, key := range v.MapKeys() {
			if i > 0 {
				fmt.Fprintf(buf, "\n%*s", indent, "")
			}
			start := buf.Len()
			buf.WriteByte('(')
			if err := encode(buf, key, indent+1); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), indent+buf.Len()-start); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Interface:
		if v.Elem().Kind() != reflect.Invalid {
			fmt.Fprintf(buf, "(%q ", v.Elem().Type())
		} else {
			fmt.Fprint(buf, `("nil" `)
		}
		encode(buf, v.Elem(), indent)
		buf.WriteByte(')')
	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

// Marshal encodes a Go value in S-expression form
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
