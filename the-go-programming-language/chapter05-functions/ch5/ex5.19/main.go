// Use panic and recover to write a function that contains no return statement
// yet returns a non-zero value.
package main

import "fmt"

func returnNonZero() (res string) {
	defer func() {
		recover()
		res = "recover"
	}()
	panic("gg")
}

func main() {
	fmt.Println(returnNonZero())
}
