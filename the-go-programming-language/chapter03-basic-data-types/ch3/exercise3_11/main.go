// Enhance comma so that it deals correctly with floating-point
// numbers and an optional sign.
package main

import (
	"bytes"
	"fmt"
	"strings"
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
	fmt.Println(comma("123451231231.123"))
	fmt.Println(comma("12345123123.123"))
	fmt.Println(comma("1234512312.123"))
	fmt.Println(comma("123451231.123"))
	fmt.Println(comma("12345121.123"))
	fmt.Println(comma("+1234511.123"))
	fmt.Println(comma("+123451.123"))
	fmt.Println(comma("+12345.123"))
	fmt.Println(comma("+1234.123"))
	fmt.Println(comma("+123.123"))
	fmt.Println(comma("+12.123"))
	fmt.Println(comma("+1.123"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	buf := bytes.Buffer{}
	mantissaStart := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		mantissaStart = 1
	}
	mantissaEnd := strings.LastIndex(s, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(s)
	}
	mantissa := s[mantissaStart:mantissaEnd]
	pre := len(mantissa) % 3
	if pre == 0 {
		pre = 3
	}
	buf.WriteString(mantissa[:pre])
	for i := pre; i < len(mantissa); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(mantissa[i : i+3])
	}
	buf.WriteString(s[mantissaEnd:])
	return buf.String()
}
