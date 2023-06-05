// Write the corresponding Pack function. Given a struct value, Pack should return a
// URL incorporating the parameter values from the struct.
package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Pack(ptr interface{}) (url.URL, error) {
	v := reflect.ValueOf(ptr).Elem()
	if v.Type().Kind() != reflect.Struct {
		return url.URL{}, fmt.Errorf("pack: %v is not a struct", ptr)
	}
	vals := &url.Values{}
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		vals.Add(name, fmt.Sprintf("%v", v.Field(i)))
	}
	return url.URL{RawQuery: vals.Encode()}, nil
}

// func Pack(ptr interface{}) (string, error) {
// 	fields := make(map[string]reflect.Value)
// 	v := reflect.ValueOf(ptr).Elem()
// 	for i := 0; i < v.NumField(); i++ {
// 		fieldInfo := v.Type().Field(i)
// 		tag := fieldInfo.Tag
// 		name := tag.Get("http")
// 		if name == "" {
// 			name = strings.ToLower(fieldInfo.Name)
// 		}
// 		fields[name] = v.Field(i)
// 	}
// 	buf := &bytes.Buffer{}
// 	for k, v := range fields {
// 		if v.Kind() == reflect.Slice {
// 			for i := 0; i < v.Len(); i++ {
// 				if buf.Len() > 0 {
// 					buf.WriteByte('&')
// 				}
// 				if err := populateurl(buf, k, v.Index(i)); err != nil {
// 					return "", fmt.Errorf("%s: %v", k, err)
// 				}
// 			}
// 		} else {
// 			if buf.Len() > 0 {
// 				buf.WriteByte('&')
// 			}
// 			if err := populateurl(buf, k, v); err != nil {
// 				return "", fmt.Errorf("%s: %v", k, err)
// 			}
// 		}
// 	}
// 	return buf.String(), nil
// }

// func populateurl(buf io.Writer, name string, v reflect.Value) error {
// 	switch v.Kind() {
// 	case reflect.String:
// 		fmt.Fprintf(buf, "%s=%s", name, v.String())
// 	case reflect.Int:
// 		fmt.Fprintf(buf, "%s=%d", name, v.Int())
// 	case reflect.Bool:
// 		fmt.Fprintf(buf, "%s=%t", name, v.Bool())
// 	default:
// 		return fmt.Errorf("unsupported kind %s", v.Type())
// 	}
// 	return nil
// }
