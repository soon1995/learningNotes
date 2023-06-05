// Write a variadic version of strings.Join
package main

import "bytes"

func join(sep string, s ...string) string {
	buf := &bytes.Buffer{}
	for i, w := range s {
		buf.WriteString(w)
		if i != len(s) - 1 {
			buf.WriteString(sep)
		}
	}
	return buf.String()
}
