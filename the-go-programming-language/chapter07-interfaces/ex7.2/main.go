// Write a function CountingWriter with the signature below that, given an io.Writer,
// returns a new Writer that wraps the original, and a pointer to an int64 variable
// that at any moment contains the number of bytes written to the new Writer
package main

import "io"

type MyWriter struct {
	written int64
	w       io.Writer
}

func (c *MyWriter) Write(b []byte) (int, error) {
	n, err := c.w.Write(b)
	c.written += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &MyWriter{0, w}
	return c, &c.written
}
