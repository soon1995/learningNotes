package main

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
)

func main() {
	buf := &bytes.Buffer{}
	type Test struct {
		Name string
		Age  int
		Work bool
	}
	mystruct := &Test{"ABC", 16, false}
	if err := encode(buf, reflect.ValueOf(mystruct)); err != nil {
		log.Fatal(err)
	}
  fmt.Println(buf.String())

}
