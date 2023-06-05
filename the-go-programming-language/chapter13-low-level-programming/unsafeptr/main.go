// Takes the address of variable x, adds the offset of its b field,
// converts the resulting to *int16, and through that pointer
// updates x.b
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
}

func main() {
	// equivalent to pb := &x.b
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42

	fmt.Println(x.b)
  reflect.ValueOf(1).Pointer()
}
