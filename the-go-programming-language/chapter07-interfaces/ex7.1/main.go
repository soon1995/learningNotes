// Using the ideas from ByteCounter, implement counters for words
// and for lines. You will find bufio.ScanWords useful.
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int
type LineCounter int
type WordCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)
	count := 0
	for sc.Scan() {
		count++
	}
  *c += LineCounter(count)
	return len(p), sc.Err()
}

func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	count := 0
	for sc.Scan() {
		count++
	}
  *c += WordCounter(count)
	return len(p), sc.Err()
}

func main() {
  txt := `  Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.
  Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.
  Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.
  Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.`
  var lc LineCounter
  fmt.Fprintf(&lc, txt)
  fmt.Println(lc)

  var wc WordCounter
  fmt.Fprintf(&wc, txt)
  fmt.Println(wc)
}
