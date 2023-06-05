// The LimitReader function in the io package accepts an io.Reader r and a
// number of bytes n, and returns another Reader that reads from r but reports an
// end-of-file condition after n bytes. Implement it.
package main

import "io"

type MyLimitReader struct {
	r io.Reader
	n int64 // byte left
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &MyLimitReader{r, n}
}

func (r *MyLimitReader) Read(b []byte) (n int, err error) {
	if r.n <= 0 {
		return 0, io.EOF
	}
  var errEOF error
	if int64(len(b)) > r.n {
		b = b[:r.n]
    errEOF = io.EOF
	}
  n, err = r.r.Read(b)
  if err == nil {
    err = errEOF
  }
	r.n += int64(n)
	return n, err
}
