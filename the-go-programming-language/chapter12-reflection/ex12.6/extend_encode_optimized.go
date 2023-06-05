// Adapt encode so that as an optimization, it does not encode a field whose value
// is the zero value of its type.
package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func encode(buf *bytes.Buffer, v reflect.Value) error {
	v.IsZero()
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
		return encode(buf, v.Elem())
	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')
	case reflect.Struct:
		buf.WriteByte('(')
		first := false
		for i := 0; i < v.NumField(); i++ {
			if !v.Field(i).IsValid() || v.Field(i).IsZero() {
				continue
			}
			if first == true {
				buf.WriteByte(' ')
			} else {
        first = true
      }
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Interface:
		if v.Elem().IsValid() {
			fmt.Fprintf(buf, "(%q ", v.Elem().Type())
		} else {
			fmt.Fprint(buf, `("nil" `)
		}
		encode(buf, v.Elem())
		buf.WriteByte(')')
	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

// Marshal encodes a Go value in S-expression form
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
