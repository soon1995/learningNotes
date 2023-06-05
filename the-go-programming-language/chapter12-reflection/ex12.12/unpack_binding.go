package main

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var bindingMap map[string]func(string) error
var (
	emailRegex      = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	creditcardRegex = regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	zipRegex        = regexp.MustCompile(`^((\d{5}-\d{4})|(\d{5})|([A-Z]\d[A-Z]\s\d[A-Z]\d))$`)
)

func init() {
	bindingMap = make(map[string]func(string) error)
	bindingMap["email"] = func(v string) error {
		if !emailRegex.MatchString(v) {
			return fmt.Errorf("%v is not a valid email", v)
		}
		return nil
	}
	bindingMap["creditCard"] = func(v string) error {
		if !creditcardRegex.MatchString(v) {
			return fmt.Errorf("%v is not a valid credit card format", v)
		}
		return nil
	}
	bindingMap["usZip"] = func(v string) error {
		if !zipRegex.MatchString(v) {
			return fmt.Errorf("%v is not a valid zip portal format", v)
		}
		return nil
	}
}

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func UnpackBinding(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	type Value struct {
		v       reflect.Value
		binding []string
	}
	// Build map of fields keyed by effective name.
	fields := make(map[string]*Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		binding := tag.Get("binding")
		bindings := strings.Split(binding, ",")
		fields[name] = &Value{v.Field(i), bindings}
	}
	// Update struct field for each parameter in the request
	for name, values := range req.Form {
		f := fields[name]
		if f == nil || !f.v.IsValid() {
			continue // ignore unrecognized HTTP parameteres
		}
		for _, value := range values {
			for _, b := range f.binding {
				if fn, ok := bindingMap[b]; ok {
					if err := fn(value); err != nil {
						return err
					}
				}
			}

			if f.v.Kind() == reflect.Slice {
				elem := reflect.New(f.v.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.v.Set(reflect.Append(f.v, elem))
			} else {
				if err := populate(f.v, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

// Take care of setting a single field v from a parameter value
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
