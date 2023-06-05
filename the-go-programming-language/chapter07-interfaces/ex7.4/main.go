// The strings.NewReader function returns a value that satisfies the io.Reader
// interface (and others) by reading from its argument, a string. Implement a simple version of
// NewReader yourself, and use it to make the HTML parser takes input from string.
package main

import (
	"io"
)

type MyReader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func NewReader(s string) *MyReader {
	return &MyReader{s, 0, -1}
}

func (r *MyReader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}
