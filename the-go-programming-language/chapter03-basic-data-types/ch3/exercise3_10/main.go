// Write a non-recursive version of comma, using bytes.Buffer
// instead of string concatenation.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123451231231"))
	fmt.Println(comma("12345123123"))
	fmt.Println(comma("1234512312"))
	fmt.Println(comma("123451231"))
	fmt.Println(comma("12345121"))
	fmt.Println(comma("1234511"))
	fmt.Println(comma("123451"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("123"))
	fmt.Println(comma("12"))
	fmt.Println(comma("1"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	pre := n % 3
	if pre == 0 {
		pre = 3
	}
	buf.WriteString(s[:pre])
	for i := pre; i < len(s); i+=3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
